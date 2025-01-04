package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsHandler(t *testing.T) {

	router := gin.Default()
	router.GET("api/products", FetchProducts)
	req, _ := http.NewRequest("GET", "api/products", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `{"message":"Product list"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}
