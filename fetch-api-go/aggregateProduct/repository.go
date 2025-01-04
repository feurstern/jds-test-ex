package aggregateproduct

import (
	"encoding/json"
	"io"
	"jdb-test-ex/product"
	"net/http"
	"os"
	"strconv"
)

var usdToIdrRate string = os.Getenv("USD_TO_IDR")

func FetchAggregateProductsRepository() ([]product.ProductStruct, error) {
	resp, err := http.Get(os.Getenv("API_URL"))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var products []product.ProductStruct
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}

	p1, _ := strconv.ParseFloat(usdToIdrRate, 32)

	for i, product := range products {
		p2, _ := strconv.ParseFloat(product.Price, 32)
		products[i].PriceIDR = p1 * p2
	}

	return products, nil
}
