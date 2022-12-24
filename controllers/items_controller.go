package controllers

import (
	"bookstore_items-api/domain/items"
	"bookstore_items-api/services"
	"fmt"
	"net/http"

	"github.com/cowanrc/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}

	fmt.Println(result)

	//TODO: Return created item with HTTP status 201
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
