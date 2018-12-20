package bar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/soljarka/go-routers/internal/app/bar/beerservice"
)

// BeerHandler serves /bar/beer/ routes
type BeerHandler struct {
}

func (b *BeerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		b.serveGET(res, req)
	case "PUT":
		b.servePUT(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (b *BeerHandler) serveGET(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)

	if head == "" {
		b.getAll(res)
		return
	}

	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}

	b.getByID(res, id)
}

func (b *BeerHandler) servePUT(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)

	if head != "" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}

	keys, ok := req.URL.Query()["name"]
	if !ok || len(keys[0]) < 1 {
		http.Error(res, "Parameter 'name' is missing", http.StatusBadRequest)
		return
	}
	b.add(res, keys[0])
}

func (b *BeerHandler) getAll(res http.ResponseWriter) {
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

func (b *BeerHandler) getByID(res http.ResponseWriter, id int) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(beerservice.GetById(id)))
}

func (b *BeerHandler) add(res http.ResponseWriter, beer string) {
	beerservice.Add(beer)
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("done"))
}
