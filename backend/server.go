package backend

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	fmt.Fprintf(w, "<html><body>Hello, World! %s</body></html>", appengine.AppID(ctx))
}
