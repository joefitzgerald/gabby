package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	msgraph "github.com/yaegashi/msgraph.go/beta"
	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	"golang.org/x/oauth2"
)

const (
	defaultTenantID       = "common"
	defaultClientID       = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultTokenCachePath = "token_cache.json"
)

var defaultScopes = []string{"offline_access", "User.Read", "Calendars.Read", "Files.Read"}

func dump(o interface{}) {
	enc := jsonx.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID, tokenCachePath string
	flag.StringVar(&tenantID, "tenant-id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client-id", defaultClientID, "Client ID")
	flag.StringVar(&tokenCachePath, "token-cache-path", defaultTokenCachePath, "Token cache path")
	flag.Parse()

	ctx := context.Background()
	m := msauth.NewManager()
	m.LoadFile(tokenCachePath)
	ts, err := m.DeviceAuthorizationGrant(ctx, tenantID, clientID, defaultScopes, nil)
	if err != nil {
		log.Fatal(err)
	}
	m.SaveFile(tokenCachePath)

	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)

	req := graphClient.Me().Request()
	user, err := req.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s (%s)\n", *user.DisplayName, *user.Mail)

	calendars, err := graphClient.Me().Calendars().Request().Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var calendar msgraph.Calendar
	for i := range calendars {
		if *calendars[i].Name == "Calendar" {
			calendar = calendars[i]
		}
	}
	r := graphClient.Me().Calendars().ID(*calendar.ID).Events().Request()
	// r.Filter("any(seriesMasterId)")
	events, err := r.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var recurring []msgraph.Event
	for i := range events {
		s := ""
		if events[i].Subject != nil {
			s = *events[i].Subject
		}
		if s == "🚨 Do Not Book 🚨" || s == "🚨 Do Not Book - Work Happening 🚨" {
			dump(events[i])
		}
		if events[i].Recurrence == nil {
			continue
		}
		if events[i].Recurrence.Range == nil {
			log.Printf("skipping event with nil range %s\n", s)
			continue
		}
		if events[i].Recurrence.Range.EndDate == nil || *events[i].Recurrence.Range.Type == "noEnd" {
			recurring = append(recurring, events[i])
			continue
		}
		t, err := events[i].Recurrence.Range.EndDate.Time()
		if err != nil {
			// TODO: is this a valid event to include?
			log.Printf("skipping event %s\n", s)
			continue
		}
		if t.After(time.Now()) {
			recurring = append(recurring, events[i])
		}
	}
	log.Printf("%v recurring events, of %v total events.", len(recurring), len(events))
	for i := range recurring {
		log.Printf("%s\n", *recurring[i].Subject)
	}
}
