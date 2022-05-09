package controller

import (
	"customermanage/apputils"
	"customermanage/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CustomerController struct {
	Store  domain.CustomerStore
	Logger *zap.Logger //DOUBT: why we specified concrete logger ? tightly coupled
}

// Create insert new Customer document by id
// Handler for the HTTP POST - "/api/users/create"
func (controller CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	var customerData domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customerData)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	err = controller.Store.Create(customerData)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	apputils.Logger.Info("new customer has been added")
}

// Delete delete the existing Customer document by id
// Handler for the HTTP DELETE - "/api/users/{id}"
func (controller CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var customerData domain.Customer
	id := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&customerData)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	err = controller.Store.Update(id, customerData)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	apputils.Logger.Info("customer details has been updated")
}

// Delete delete the existing Customer document by id
// Handler for the HTTP DELETE - "/api/users/{id}"
func (controller CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := controller.Store.Delete(id)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	apputils.Logger.Info("deleted the requested customer details")
}

// GetById return the existing Customer document by id
// Handler for the HTTP Get - "/api/users/{id}"
func (controller CustomerController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := controller.Store.GetById(id)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(customer)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	apputils.Logger.Info("fetched the customer customer details")
}

// Getall get all existing Customer document
// Handler for the HTTP Get - "/api/users/"
func (controller CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executed GetAll")
	customerdata, err := controller.Store.GetAll()
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	allcustomers := AllCustomers{Users: customerdata}
	j, err := json.Marshal(allcustomers)
	if err != nil {
		apputils.Logger.Error(err.Error())
		apputils.DisplayAppError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	apputils.Logger.Info("fetching all customer details")
}
