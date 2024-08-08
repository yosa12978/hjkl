package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yosa12978/hjkl/types"
)

func WriteJson(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func ReadJson(r io.Reader, foo any) error {
	return json.NewDecoder(r).Decode(foo)
}

// Writes json (types.Message) to w
func WriteMsg(w http.ResponseWriter, statusCode int, msg string) error {
	return WriteJson(w, statusCode, types.Message{
		StatusCode: statusCode,
		Message:    msg,
	})
}
