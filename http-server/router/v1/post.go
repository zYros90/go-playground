package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var v map[string]interface{}
	err := decoder.Decode(&v)
	if err != nil {
		fmt.Println(err)
		errorResponse(w, "can't decode request body", http.StatusBadRequest)
		return
	}
	fmt.Println("request body: ", v)
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(map[string]string{"message": "ok"})
	if err != nil {
		fmt.Println(err)
		errorResponse(w, "can't marshal response to json", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
		return
	}
}
