package bar

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

// Bar is a bar REST API
type Bar struct {
	BarHandler *BarHandler
}

func (b *Bar) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)
	if head == "bar" {
		b.BarHandler.ServeHTTP(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Go to /bar/ for Bar API"))
}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
