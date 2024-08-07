package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteJson(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func ReadJson(r io.Reader, foo any) error {
	return json.NewDecoder(r).Decode(foo)
}
