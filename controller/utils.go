package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data any `json:"data"`
}

func WriteJson(w http.ResponseWriter, status int, body any) {
	jsonData, err := json.Marshal(Response{
		Data: body,
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	fmt.Fprint(w, string(jsonData))
}
