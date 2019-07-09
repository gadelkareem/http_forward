package http_forward

import (
	"flag"
	h "github.com/gadelkareem/go-helpers"
	"net/http"
)

var (
	infile = flag.String("infile", "https://example.com/", "URL to forward request to")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *infile, 301)
	})
	h.PanicOnError(http.ListenAndServe(":80", nil))
}
