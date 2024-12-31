package utils

import "net/http"

func NewErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func NewSuccessResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
