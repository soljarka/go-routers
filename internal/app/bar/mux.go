package bar

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/soljarka/go-routers/internal/app/bar/beerservice"
	"github.com/soljarka/go-routers/internal/app/bar/snackservice"
)

//ServeRoot serves /
func ServeRoot(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Go to /bar/ for Bar API"))
}

//ServeBar serves /bar
func ServeBar(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Welcome to the Bar API. Go go /bar/beer or /bar/snacks"))
}

//ServeAllBeers serves GET /bar/beer
func ServeAllBeers(res http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(beerservice.GetAll())
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Can't encode response"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(bytes)
}

//ServeBeerByID serves GET /bar/beer/{id}
func ServeBeerByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(beerservice.GetById(id)))
}

//ServeAddBeer serves PUT /bar/beer?name=[snackname]
func ServeAddBeer(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["name"]
	if !ok || len(keys[0]) < 1 {
		http.Error(res, "Parameter 'name' is missing", http.StatusBadRequest)
		return
	}
	beerservice.Add(keys[0])
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("done"))
}

//ServeAllSnacks serves GET /bar/snack
func ServeAllSnacks(res http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(snackservice.GetAll())
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Can't encode response"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(bytes)
}

//ServeSnackByID serves GET /bar/snack/{id}
func ServeSnackByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(snackservice.GetById(id)))
}

//ServeAddSnack serves PUT /bar/snack?name=[snackname]
func ServeAddSnack(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["name"]
	if !ok || len(keys[0]) < 1 {
		http.Error(res, "Parameter 'name' is missing", http.StatusBadRequest)
		return
	}
	snackservice.Add(keys[0])
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("done"))
}
