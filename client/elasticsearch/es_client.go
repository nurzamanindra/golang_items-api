package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/nurzamanindra/golang_items-api/logger"
	"github.com/olivere/elastic/v7"
)

var (
	Client esClientInterface = &esClient{}
)

type esClient struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stderr, "", log.LstdFlags)),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (es *esClient) setClient(client *elastic.Client) {
	es.client = client
}

func (es *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := es.client.Index().Index(index).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in elasticsearch index %s", index), err)
		return nil, err
	}
	return result, nil

}

func (es *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := es.client.Get().Index(index).Type(docType).Id(id).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to get index document in elasticsearch id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (es *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := es.client.Search(index).Query(query).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to search index document in elasticsearch index %s", index), err)
		return nil, err
	}
	return result, nil
}
