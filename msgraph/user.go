package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"

	"github.com/joefitzgerald/gabby"
)

func GetFileContentType(b []byte) string {
	buffer := b[:512]
	contentType := http.DetectContentType(buffer)
	return contentType
}
func (a *API) GetName(ctx context.Context, id string) (string, error) {
	u, err := a.Client.Users().ID(id).Request().Get(ctx)

	if err != nil {
		log.Println("unable to get user details")
		return "", err
	}

	return *u.DisplayName, nil
}

func (a *API) GetPhoto(ctx context.Context, id string) (*gabby.Photo, error) {

	pr, err := a.Client.Users().ID(id).Request().NewRequest("GET", "/photo/$value", nil)
	if err != nil {
		log.Println("Unable to create photo request")
		return nil, err
	}

	r, err := a.HttpClient.Do(pr)

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("no photo found for ID: '%s' (%d)", id, r.StatusCode)
	}

	if err != nil {
		log.Println("Unable to retrieve photo")
		return nil, err
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("Unable to read photo binary data")
		return nil, err
	}

	exts, err := mime.ExtensionsByType(GetFileContentType(b))
	if err != nil {
		log.Println("Unable to determine extension from content type")
		return nil, err
	}

	return &gabby.Photo{Data: b, ID: id, SuggestedExtension: exts[0]}, nil
}
