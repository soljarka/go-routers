package bar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/soljarka/go-routers/internal/app/bar/snackservice"
)

// SnackHandler serves /bar/snack/ routes
type SnackHandler struct {
}

func (s *SnackHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.serveGET(res, req)
	case "PUT":
		s.servePUT(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (s *SnackHandler) serveGET(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)

	if head == "" {
		s.getAll(res)
		return
	}

	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}

	s.getByID(res, id)
}

func (s *SnackHandler) servePUT(res http.ResponseWriter, req *http.Request) {
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
	s.add(res, keys[0])
}

func (s *SnackHandler) getAll(res http.ResponseWriter) {
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

func (s *SnackHandler) getByID(res http.ResponseWriter, id int) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(snackservice.GetById(id)))
}

func (s *SnackHandler) add(res http.ResponseWriter, snack string) {
	snackservice.Add(snack)
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("done"))
}
