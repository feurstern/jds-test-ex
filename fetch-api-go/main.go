package main

import (
	aggregateproduct "jdb-test-ex/aggregateProduct"
	"jdb-test-ex/auth"
	"jdb-test-ex/claims"
	"jdb-test-ex/product"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var LOCALHOST string = os.Getenv("LOCALHOST")
var CLIENT_HOST string = os.Getenv("CLIENT_HOST")

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:9000"},
		AllowMethods:  []string{"POST", "GET"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Fetch api's backend succesfully connected!"})
	})

	secured := router.Group("api")
	secured.Use(auth.ValidateJWT())
	{
		secured.GET("/products", product.FetchProducts)
		secured.GET("/products/aggregate", aggregateproduct.FetchAggregateProducts)
		secured.GET("/claims", claims.GetClaim)
	}

	router.Run("localhost:8181")
}
