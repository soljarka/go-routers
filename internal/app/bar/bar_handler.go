package bar

import (
	"fmt"
	"net/http"
)

// BarHandler serves /bar/ routes
type BarHandler struct {
	BeerHandler  *BeerHandler
	SnackHandler *SnackHandler
}

func (b *BarHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)
	switch head {
	case "beer":
		b.BeerHandler.ServeHTTP(res, req)
	case "snack":
		b.SnackHandler.ServeHTTP(res, req)
	default:
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("Welcome to the Bar API. Go go /bar/beer or /bar/snacks"))
	}
}
