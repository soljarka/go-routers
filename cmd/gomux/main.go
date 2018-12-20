package main

import (
	"net/http"

	"github.com/soljarka/go-routers/internal/app/gomux"
)

func main() {
	router := gomux.New()
	http.ListenAndServe(":8080", router)
}
