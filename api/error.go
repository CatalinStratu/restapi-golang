package api

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{http.StatusNotFound, "Status not found!"})
		return
	}
}
