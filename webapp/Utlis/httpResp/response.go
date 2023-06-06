package httpResp

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})

}

func RespondWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	response, _ := json.Marshal(playload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
