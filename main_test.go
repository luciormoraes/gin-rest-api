package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luciormoraes/gin-rest-api/controllers"
	"github.com/stretchr/testify/assert"
)

func RouteSetupTest() *gin.Engine {
	routes := gin.Default()
	return routes
}

// func TestFailure(t *testing.T) {
// 	t.Fatalf("Test has failed")
// }

func TestStatusCodeOK(t *testing.T) {
	r := RouteSetupTest()
	r.GET("/:name", controllers.Salutation)
	req, _ := http.NewRequest("GET", "/gin", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	// if response.Code != http.StatusOK {
	// 	t.Fatalf("Status error: Received value is %d. Expected was %d", response.Code, http.StatusOK)
	// }
	assert.Equal(t, http.StatusOK, response.Code, "Status error: Received value is not 200")
	mockResponse := `{"API says":"Hey gin, all good?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody), "Must be equal")
}
