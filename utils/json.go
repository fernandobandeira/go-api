package utils

import (
	"encoding/json"
	"net/http"
)

func ReadJson(r *http.Request, i interface{}) error {
	return json.NewDecoder(r.Body).Decode(i)
}

func WriteJson(w http.ResponseWriter, i interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}
