package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func GetParam(r *http.Request, key string) (*string, error) {
	param := r.PathValue(key)

	if len(param) == 0 {
		return nil, fmt.Errorf("param [%s] is required", key)
	}

	return &param, nil
}

func GetIntParam(r *http.Request, key string) (*int, error) {
	param, err := GetParam(r, key)

	if err != nil {
		return nil, err
	}

	number, err := strconv.Atoi(*param)

	if err != nil {
		return nil, fmt.Errorf("param [%s] should be a number", key)
	}

	return &number, nil
}
