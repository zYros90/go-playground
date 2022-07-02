package v1

import (
	"http-server/router/middlewares"
	"net/http"
)

func RegisterRoutes(muxer *http.ServeMux) {
	// simple get request without middleware: HandleFunc
	muxer.HandleFunc("/hello", hello)
	muxer.HandleFunc("/", hello)

	// simple get request with middleware: HandlerFunc
	muxer.Handle("/ping", middlewares.JWT(http.HandlerFunc(pong)))

	// simple post request
	muxer.HandleFunc("/post", post)
}
