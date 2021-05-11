package response

import (
	"encoding/json"
	"net/http"
)

type logger interface {
	Error(args ...interface{})
}

// BadRequest error message for api
func BadRequest(l logger, err error, w http.ResponseWriter) {
	l.Error(err)
	// TODO: put error in pretty json message, not just plain text
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// InternalError error message for api
func InternalError(l logger, err error, w http.ResponseWriter) {
	l.Error(err)
	// TODO: do not show internal errors
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Ok response with 200 code and response structure
func Ok(l logger, w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		l.Error(err)
		// TODO: do not show internal errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
