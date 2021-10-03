package handler

import (
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("employees list..."))
	w.WriteHeader(http.StatusOK)
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("clients list..."))
	w.WriteHeader(http.StatusOK)
}
