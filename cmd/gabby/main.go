package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/joefitzgerald/gabby"
	"github.com/joefitzgerald/gabby/command"
	"github.com/joefitzgerald/gabby/msgraph"
	msgraphDotGo "github.com/yaegashi/msgraph.go/beta"
	"github.com/yaegashi/msgraph.go/msauth"
	"golang.org/x/oauth2"
)

var defaultScopes = []string{"offline_access", "User.Read", "Calendars.Read"}

var cli struct {
	command.Context
	Impact command.Impact `cmd:"" help:"Perform an impact analysis of events over a given time period."`
}

func main() {
	log.SetFlags(log.Lshortfile)
	ctx := kong.Parse(&cli, kong.UsageOnError(), kong.Name("gabby"))
	api, err := API(cli.TenantID, cli.ClientID, cli.TokenCacheFilename)
	if err != nil {
		log.Fatal(err)
	}
	me, err := api.GetMe(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	cli.Context.Me = me
	var w io.Writer = os.Stdout
	ctx.Bind(cli.Context)
	ctx.BindTo(api, (*gabby.API)(nil))
	ctx.BindTo(w, (*io.Writer)(nil))
	err = ctx.Run()
	ctx.FatalIfErrorf(err)
}

func API(tenantID string, clientID string, tokenCacheFilename string) (gabby.API, error) {
	ctx := context.Background()
	m := msauth.NewManager()
	tokenCachePath, err := getTokenCachePath(tokenCacheFilename)
	if err != nil {
		return nil, err
	}
	_ = m.LoadFile(tokenCachePath)
	ts, err := m.DeviceAuthorizationGrant(ctx, tenantID, clientID, defaultScopes, nil)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve device authorization grant: %w", err)
	}
	err = m.SaveFile(tokenCachePath)
	if err != nil {
		return nil, fmt.Errorf("could not save token to token cache path: %w", err)
	}

	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraphDotGo.NewClient(httpClient)
	return &msgraph.API{
		Client: graphClient,
	}, nil
}

func getTokenCachePath(filename string) (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(userCacheDir, "gabby", filename)
	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return "", err
	}
	return path, nil
}
