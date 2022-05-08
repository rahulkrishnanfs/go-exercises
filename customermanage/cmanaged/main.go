package main

import (
	"customermanage/router"
	"net/http"
)

func main() {

	router := router.InitRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
