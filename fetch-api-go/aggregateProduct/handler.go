package aggregateproduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesAggregateProduct(router *gin.Engine) {
	router.GET("aggregate/product")
}

func FetchAggregateProducts(c *gin.Context) {
	aggregateProduct, err := FetchAggregateProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the aggregate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": aggregateProduct})

}
