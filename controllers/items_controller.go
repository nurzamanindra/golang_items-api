package controllers

import (
	"fmt"
	"net/http"

	"github.com/nurzamanindra/golang_items-api/domain/items"
	"github.com/nurzamanindra/golang_items-api/services"
)

var (
	ItemController itemControllersInterface = &itemController{}
)

type itemControllersInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (i *itemController) Create(w http.ResponseWriter, r *http.Request) {
	//hit authentication API
	//get caller id from oauth api

	item := items.Item{}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		// return error json to user
	}
	fmt.Println(result)
	//return created item as json with http status 201 - created

}

func (i *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
