package api

import (
	"io"
	"net/http"
)

func CreateRoute() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "API route.")
	})

	return r
}
