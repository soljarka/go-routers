package main

import (
	"github.com/ctco-dev/go-routers/gomux"
	"net/http"
)

func main() {
	router := gomux.New()
	http.ListenAndServe(":8080", router)
}