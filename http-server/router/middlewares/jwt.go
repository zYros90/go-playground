package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

func JWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			// TODO check signing
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}
