package common

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Access-Control-Expose-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	log.Errorf("respondWithError: %v", message)
	RespondWithJSON(w, code, map[string]string{"error": message})
}
