package main

import (
	"net/http"

	"github.com/soljarka/go-routers/internal/app/gorillamux"
)

func main() {
	router := gorillamux.New()
	http.ListenAndServe(":8080", router)
}
