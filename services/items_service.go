package services

import (
	"github.com/nurzamanindra/golang_items-api/domain/items"
	"github.com/nurzamanindra/golang_items-api/domain/queries"
	"github.com/nurzamanindra/golang_items-api/utils/rest_errors"
)

var (
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, *rest_errors.RestErr)
}

type itemService struct{}

func (i *itemService) Create(itemRequest items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}
	return &itemRequest, nil
}

func (i *itemService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *itemService) Search(query queries.EsQuery) ([]items.Item, *rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
