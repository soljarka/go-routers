package main

import (
	"go-routers/gorillamux"
	"net/http"
)

func main() {
	router := gorillamux.New()
	http.ListenAndServe(":8080", router)
}