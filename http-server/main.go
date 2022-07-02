package main

import (
	"flag"
	"fmt"
	"http-server/router"
	"net/http"
)

func main() {
	port := flag.Int("port", 9292, "port of http server")
	flag.Parse()
	fmt.Println("starting server on port: ", *port)
	router := router.New()
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
