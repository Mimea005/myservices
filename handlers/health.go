package handlers

import (
	"encoding/json"
	"net/http"
	"myservices/services"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	services.Health()
	Log.Println("Sendt health update")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
