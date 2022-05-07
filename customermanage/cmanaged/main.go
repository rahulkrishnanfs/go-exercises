package main

import (
	"customermanage/router"
	"net/http"
)

func main() {

	router := router.InitRoutes()
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	server.ListenAndServe()
	// controller := CustomerController{store: mapstore.NewMapStore()}

	// controller.Create(domain.Customer{ID: "11", Name: "Rahul", Email: "rahul@test.com"})
	// controller.Create(domain.Customer{ID: "12", Name: "Mittuz", Email: "mittuz@test.com"})
	// controller.Create(domain.Customer{ID: "13", Name: "Rasil", Email: "rasil@test.com"})
	// controller.GetAll()
	// controller.Update("13", domain.Customer{ID: "13", Name: "Raaasil", Email: "raaaaasil@test.com"})
	// controller.GetAll()
	// controller.GetById("11")
	// controller.GetAll()
	// controller.Delete("11")
	// controller.Delete("13")
	// controller.GetAll()

}
