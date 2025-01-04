package product

type ProductStruct struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      string  `json:"price"`
	Department string  `json:"department"`
	PriceIDR   float64 `json:"price_idr, omitempty"`
}
