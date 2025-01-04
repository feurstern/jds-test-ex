package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesProduct(router *gin.Engine) {
	router.GET("/products", FetchProducts)
}

func FetchProducts(c *gin.Context) {
	product, err := FetchProductsService()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch the product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": product})
}
