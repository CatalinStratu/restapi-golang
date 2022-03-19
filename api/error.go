package api

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		jsonResp, err := json.Marshal(ErrorResponse{http.StatusNotFound, "Status not found!"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.Write(jsonResp)
		return
	}
}
