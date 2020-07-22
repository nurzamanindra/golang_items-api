package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurzamanindra/golang_items-api/domain/items"
	"github.com/nurzamanindra/golang_items-api/domain/queries"
	"github.com/nurzamanindra/golang_items-api/services"
	"github.com/nurzamanindra/golang_items-api/utils/http_utils"
	"github.com/nurzamanindra/golang_items-api/utils/rest_errors"
)

var (
	ItemController itemControllersInterface = &itemController{}
)

type itemControllersInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (i *itemController) Create(w http.ResponseWriter, r *http.Request) {
	//hit authentication API
	//get caller id from oauth api

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json request")
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json request")
		http_utils.ResponseError(w, restErr)
		return
	}
	//TODO : get seller id from oauth

	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		// return error json to user
		http_utils.ResponseError(w, createErr)
		return
	}

	//return created item as json with http status 201 - created
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (i *itemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := services.ItemsService.Get(vars["id"])
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, result)
}

func (i *itemController) Search(w http.ResponseWriter, r *http.Request) {
	var query queries.EsQuery
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json request")
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &query); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json request")
		http_utils.ResponseError(w, restErr)
		return
	}

	items, errSearch := services.ItemsService.Search(query)

	if errSearch != nil {
		http_utils.ResponseError(w, errSearch)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, items)

}
