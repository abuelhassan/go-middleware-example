package handler

import (
	"log"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value("logger").(*log.Logger)
	logger.Println("GetEmployees")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("employees list..."))
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value("logger").(*log.Logger)
	logger.Println("GetClients")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("clients list..."))
}
