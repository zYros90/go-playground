package router

import (
	v1 "http-server/router/v1"
	"net/http"
)

func New() *http.ServeMux {
	muxer := http.NewServeMux()
	v1.RegisterRoutes(muxer)
	return muxer
}
