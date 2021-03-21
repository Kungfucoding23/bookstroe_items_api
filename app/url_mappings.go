package app

import (
	"net/http"

	"github.com/Kungfucoding23/bookstore_items_api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}