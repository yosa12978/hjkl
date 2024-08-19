package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yosa12978/hjkl/types"
	"github.com/yosa12978/hjkl/validation"
)

func WriteJson(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func ReadJson[T any](r *http.Request) (T, error) {
	var v T
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}

func ReadJsonAndValidate[T validation.Validatable](r *http.Request) (T, map[string]string, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, err
	}
	if problems := v.Validate(r.Context()); problems != nil {
		return v, problems, fmt.Errorf("validation error")
	}
	return v, nil, nil
}

// Writes json (types.Message) to w
func WriteMsg(w http.ResponseWriter, statusCode int, msg string) error {
	return WriteJson(w, statusCode, types.Message{
		StatusCode: statusCode,
		Message:    msg,
	})
}
