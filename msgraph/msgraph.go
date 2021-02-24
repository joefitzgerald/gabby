package msgraph

import (
	msgraph "github.com/yaegashi/msgraph.go/beta"
)

type API struct {
	Client *msgraph.GraphServiceRequestBuilder
}

// func dump(o interface{}) {
// 	enc := json.NewEncoder(os.Stdout)
// 	enc.SetIndent("", "  ")
// 	err := enc.Encode(o)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
