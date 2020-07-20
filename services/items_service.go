package services

import (
	"github.com/nurzamanindra/golang_items-api/domain/items"
	"github.com/nurzamanindra/golang_items-api/utils/rest_errors"
)

var (
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemService struct{}

func (i *itemService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}

func (i *itemService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}
