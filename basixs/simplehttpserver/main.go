package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpServer struct {
	routes  []route
	port    string
	handler http.Handler
}

type route struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

func (h httpServer) start() error {
	fmt.Printf("http server listening on localhost%s\n", h.port)
	return http.ListenAndServe(h.port, nil)
}

func (h httpServer) registerRoutes() {
	for _, r := range h.routes {
		http.HandleFunc(r.pattern, r.handler)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Bitch"))
}

var r []route = []route{
	route{
		"/hello",
		handleHello,
	},
}

func main() {
	newHttpServer := httpServer{
		routes: r,
		port:   ":9090",
	}
	newHttpServer.registerRoutes()
	if err := newHttpServer.start(); err != nil {
		log.Fatal(err)
	}

}
