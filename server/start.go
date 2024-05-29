package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nameless9000/cm2workshop/api"
	"github.com/nameless9000/cm2workshop/web"
)

func StartServer(port uint16) {
	webRoute := web.CreateRoute()
	apiRoute := api.CreateRoute()

	webAddr := ":" + strconv.FormatUint(uint64(port), 10)
	apiAddr := ":" + strconv.FormatUint(uint64(port+1), 10)

	go func() {
		log.Println("Started API server on", apiAddr)
		log.Fatalln(http.ListenAndServe(apiAddr, apiRoute))
	}()

	log.Println("Started web server on", webAddr)
	log.Fatalln(http.ListenAndServe(webAddr, webRoute))
}
