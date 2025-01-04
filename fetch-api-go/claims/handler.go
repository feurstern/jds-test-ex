package claims

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ClaimsRoutes(router *gin.Engine) {
	router.GET("/claims", GetClaim)
}

func GetClaim(c *gin.Context) {
	claims := c.MustGet(os.Getenv("JWT_KEY_SECRET")).(jwt.MapClaims)
	c.JSON(http.StatusOK, claims)
}
