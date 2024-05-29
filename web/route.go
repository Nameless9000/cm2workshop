package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/nameless9000/cm2workshop/web/components"
)

func CreateRoute() *http.ServeMux {
	mux := http.NewServeMux()

	assets := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", assets))

	index_page := components.Index()
	mux.Handle("/", templ.Handler(index_page))

	countPage := components.CounterPage()
	mux.Handle("GET /count", templ.Handler(countPage))
	mux.HandleFunc("POST /count", components.IncrementCounter)

	return mux
}
