package claims

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ClaimsRoutes(router *gin.Engine) {
	router.GET("/claims", GetClaim)
}

func GetClaim(c *gin.Context) {
	claims, ok := c.Get(os.Getenv("JWT_KEY_SECRET"))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"claims": claims})
}
