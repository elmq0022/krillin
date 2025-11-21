package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/elmq0022/kami/types"
)

func JsonAdapter(w http.ResponseWriter, req *http.Request, handler types.Handler) {
	status, result, err := handler(req)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		// Handler returned an error; respond with 500 unless status provided
		if status == 0 {
			status = http.StatusInternalServerError
		}
		resp := map[string]any{"error": err.Error()}
		data, jerr := json.Marshal(resp)
		if jerr != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		_, _ = w.Write(data)
		return
	}

	if status == 0 {
		status = http.StatusOK
	}

	data, jerr := json.Marshal(result)
	if jerr != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(data)
}
