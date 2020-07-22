package items

import (
	"encoding/json"
	"fmt"

	"github.com/nurzamanindra/golang_items-api/client/elasticsearch"
	"github.com/nurzamanindra/golang_items-api/domain/queries"
	"github.com/nurzamanindra/golang_items-api/utils/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item")
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get item")
	}
	if !result.Found {
		return rest_errors.NewInternalServerError(fmt.Sprintf("no item found with id %s", i.Id))
	}

	bytes, err := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to unmarshall item from elasticsearch")

	}
	i.Id = itemId
	fmt.Println(err)
	fmt.Println(string(bytes))
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, *rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents from elasticsearch")
	}
	items := make([]Item, len(result.Hits.Hits))
	fmt.Println(result)
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to unmarshal documents from elasticsearch")
		}
		item.Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no document matching the queries")
	}
	return items, nil
}
