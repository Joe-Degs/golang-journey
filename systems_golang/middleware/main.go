// This package demonstrates a simple net/http request midleware creation by
// using the http.RoundTripper interface to intercept requests.
package main

import (
	"log"
	"net/http"
	"time"
)

type logTripper struct {
	http.RoundTripper
}

// intercept requests and log the url before executing it
func (l logTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println(r.URL)
	r.Header.Set("X-Log-Time", time.Now().String())
	return l.RoundTripper.RoundTrip(r)
}

func main() {
	return
}
