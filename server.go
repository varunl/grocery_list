package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", "8080", "Port to start the webserver on")

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)

	gh := &GroceryHandler{NewInMemoryStorage()}
	http.Handle("/grocery", gh)

	fmt.Printf("Starting webserver, listening on port %s...\n", *port)
	http.ListenAndServe(":"+*port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

type GroceryHandler struct {
	store Storage
}

func (g *GroceryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Grocery Handler")
}
