package httpserver

import (
	"fmt"
	"net/http"
)

// This function should start a web server on the PORT 'port'
// When a request is done to it, it should return "Hello World!"
func HttpServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":"+port, nil)
}
