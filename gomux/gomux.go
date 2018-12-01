package gomux

import (
	"go-routers/middleware"
	"log"
	"net/http"
)

func New() *http.ServeMux {
	handler := http.HandlerFunc(handle)
	router := http.NewServeMux()
	router.Handle("/", middleware.Auth(middleware.Log(handler)))
	return router
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("handling user request")
	w.Write([]byte("Hello world"))
}
