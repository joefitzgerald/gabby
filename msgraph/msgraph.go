package msgraph

import (
	"net/http"

	msgraph "github.com/yaegashi/msgraph.go/beta"
)

//API represents the external calendar interface
type API struct {
	HttpClient *http.Client
	Client     *msgraph.GraphServiceRequestBuilder
}

// func dump(o interface{}) {
// 	enc := json.NewEncoder(os.Stdout)
// 	enc.SetIndent("", "  ")
// 	err := enc.Encode(o)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
