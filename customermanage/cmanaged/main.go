package main

import (
	"customermanage/apputils"
	"customermanage/bootstrapper"
	"customermanage/router"
	"net/http"
)

func main() {
	bootstrapper.StartUp()
	router := router.InitRoutes() // configure the routes
	server := &http.Server{
		Addr:    apputils.AppConfig.Server,
		Handler: router,
	}

	apputils.Logger.Info("starting http server...")
	server.ListenAndServe() //Start the http server
	//server.ListenAndServeTLS("../certs/server.crt", "../keys/server.key") //Start the http server

}
