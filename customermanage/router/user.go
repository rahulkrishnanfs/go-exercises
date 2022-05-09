package router

import (
	"customermanage/controller"
	"customermanage/domain"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// SetUserRoutes resgiter the routes for user entry
func SetUserRoutes(router *mux.Router, newstore domain.CustomerStore, logger *zap.Logger) *mux.Router { //DOUBT: domain.CustomerStore interface can recieceve pointer to mapstore?

	customerController := controller.CustomerController{Store: newstore, Logger: logger} //injecting the dependency
	userRouter := router.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/create", customerController.Create).Methods("POST")
	userRouter.HandleFunc("/{id}", customerController.Update).Methods("PUT")
	userRouter.HandleFunc("/{id}", customerController.Delete).Methods("DELETE")
	userRouter.HandleFunc("/{id}", customerController.GetById).Methods("GET")
	userRouter.HandleFunc("/", customerController.GetAll).Methods("GET")

	return router

}
