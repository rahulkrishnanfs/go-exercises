package router

import (
	"customermanage/mapstore"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// InitRoutes Registers all routes for the application
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	datastore := mapstore.NewMapStore()
	logger, _ := zap.NewProduction()
	router = SetUserRoutes(router, datastore, logger)
	return router
}
