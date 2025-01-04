package product

func FetchProductsService() ([]ProductStruct, error) {
	return FetchProductsRepository()
}
