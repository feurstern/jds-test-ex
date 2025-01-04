package product

import (
	"encoding/json"
	"net/http"
	"os"
)

func FetchProductsRepository() ([]ProductStruct, error) {
	resp, err := http.Get(os.Getenv("API_URL"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []ProductStruct
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}
