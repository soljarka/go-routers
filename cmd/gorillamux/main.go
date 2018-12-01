package main

import (
	"github.com/ctco-dev/go-routers/gorillamux"
	"net/http"
)

func main() {
	router := gorillamux.New()
	http.ListenAndServe(":8080", router)
}