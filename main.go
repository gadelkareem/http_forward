package http_forward

import (
	"flag"
	h "github.com/gadelkareem/go-helpers"
	"net/http"
)

var (
	url = flag.String("url", "https://example.com/", "URL to forward request to")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *url, 301)
	})
	h.PanicOnError(http.ListenAndServe(":80", nil))
}
