package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// Health GET endpoint /health
func Health(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["status"] = "OK"
	data["timestamp"] = time.Now().Format(time.RFC3339)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
