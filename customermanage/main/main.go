package main

import (
	"customermanage/domain"
	"customermanage/mapstore"
	"fmt"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (controller CustomerController) Create(customer domain.Customer) {
	err := controller.store.Create(customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New customer has been added")
}

func (controller CustomerController) Update(id string, customer domain.Customer) {
	err := controller.store.Update(id, customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Customer details has been updated")
}

func (controller CustomerController) Delete(id string) {
	err := controller.store.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted the requested customer details")
}

func (controller CustomerController) GetById(id string) {
	customer, err := controller.store.GetById(id)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Fetched the customer customer details", customer)
}

func (controller CustomerController) GetAll() {
	customerdata, err := controller.store.GetAll()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Fetching all customer details")
	for _, customer := range customerdata {
		fmt.Println(customer)
	}
}

func main() {
	controller := CustomerController{store: mapstore.NewMapStore()}

	controller.Create(domain.Customer{ID: "11", Name: "Rahul", Email: "rahul@test.com"})
	controller.Create(domain.Customer{ID: "12", Name: "Mittuz", Email: "mittuz@test.com"})
	controller.Create(domain.Customer{ID: "13", Name: "Rasil", Email: "rasil@test.com"})
	controller.GetAll()
	controller.Update("13", domain.Customer{ID: "13", Name: "Raaasil", Email: "raaaaasil@test.com"})
	controller.GetAll()
	controller.GetById("11")
	controller.GetAll()
	controller.Delete("11")
	controller.Delete("13")
	controller.GetAll()

}
