package gorillamux

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soljarka/go-routers/internal/pkg/middleware"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Auth, middleware.Log)
	router.HandleFunc("/", handle)
	return router
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("handling user request")
	w.Write([]byte("Hello world"))
}
