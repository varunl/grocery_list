package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", "8080", "Port to start the webserver on")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)

	fmt.Printf("Starting webserver, listening on port %s...\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
