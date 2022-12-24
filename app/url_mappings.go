package app

import (
	"bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
