package gomux

import (
	"log"
	"net/http"

	"github.com/soljarka/go-routers/internal/pkg/middleware"
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
