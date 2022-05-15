package algolia

import (
	"sync"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/eyewa/eyewa-go-lib/log"
	"github.com/eyewa/migrator/config"
)

const (
	ProductEventType   = "Product"
	ProductIndexPrefix = "product-"
)

var stores = []string{
	"ae-en",
	"ae-ar",
	"sa-en",
	"sa-ar",
	"us-en",
	"kw-en",
	"kw-ar",
	"qa-en",
	"qa-ar",
	"om-en",
	"om-ar",
	"bh-en",
	"bh-ar",
}

func InitIndexes() []*search.Index {
	algoliaClient := search.NewClient(config.Config.Algolia.ApplicationID, config.Config.Algolia.APIKey)

	indexes := make([]*search.Index, 0)
	for _, store := range stores {
		productIndex := algoliaClient.InitIndex(ProductIndexPrefix + store)
		indexes = append(indexes, productIndex)
	}

	return indexes
}

func ClearIndices() error {
	indexes := InitIndexes()
	var wg sync.WaitGroup

	for _, index := range indexes {
		res, err := index.ClearObjects()
		if err != nil {
			log.Error("Got error while queueing " + index.GetName() + " index cleaning")
			log.Error(err.Error())
			return err
		}

		wg.Add(1)
		go func(index *search.Index) {
			defer wg.Done()
			err = res.Wait()

			if err != nil {
				log.Error("Got error while cleaning " + index.GetName() + " index")
				log.Error(err.Error())
				return
			}

			log.Info(index.GetName() + " index cleaned successfully")
		}(index)
	}

	wg.Wait()

	return nil
}
