package web

import (
	"io"
	"net/http"
)

func CreateRoute() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Web route.")
	})

	return r
}
