package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/soljarka/go-routers/internal/app/bar"
)

func main() {
	serveMux()
}

func servePureGo() {
	beerHandler := &bar.BeerHandler{}
	snackHandler := &bar.SnackHandler{}
	barHandler := &bar.BarHandler{BeerHandler: beerHandler, SnackHandler: snackHandler}
	bar := &bar.Bar{BarHandler: barHandler}
	http.ListenAndServe(":8080", bar)
}

func serveMux() {
	r := mux.NewRouter()
	r.HandleFunc("/bar/snack", bar.ServeAllSnacks).Methods("GET")
	r.HandleFunc("/bar/snack", bar.ServeAddSnack).Methods("PUT")
	r.HandleFunc("/bar/snack/{id}", bar.ServeSnackByID).Methods("GET")
	r.HandleFunc("/bar/beer", bar.ServeAllBeers).Methods("GET")
	r.HandleFunc("/bar/beer", bar.ServeAddBeer).Methods("PUT")
	r.HandleFunc("/bar/beer/{id}", bar.ServeBeerByID).Methods("GET")
	r.HandleFunc("/bar", bar.ServeBar)
	r.HandleFunc("/", bar.ServeRoot)
	http.ListenAndServe(":8080", r)
}
