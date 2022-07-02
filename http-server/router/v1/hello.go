package v1

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Header:", r.Header)
	fmt.Println("Context:", r.Context())
}
