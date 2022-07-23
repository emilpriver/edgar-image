package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpErrorMessage(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = message

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	return
}
