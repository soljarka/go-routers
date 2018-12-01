package main

import (
	"go-routers/gomux"
	"net/http"
)

func main() {
	router := gomux.New()
	http.ListenAndServe(":8080", router)
}