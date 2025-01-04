package aggregateproduct

import "jdb-test-ex/product"

func FetchAggregateProductsService() ([]product.ProductStruct, error) {
	return FetchAggregateProductsRepository()
}
