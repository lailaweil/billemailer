package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, entity interface{}, status int) {
	response, errMarshal := json.Marshal(entity)
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
