package error

import (
	"encoding/json"
	"net/http"
)

func JResponse(w http.ResponseWriter,statusCode int, message string){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error":message})
}