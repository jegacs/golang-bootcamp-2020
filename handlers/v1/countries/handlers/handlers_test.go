package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFetchFromAPIRequest(t *testing.T) {
	Convey("Given a GET request to /countries/{country_code}/fetch", t, func() {
		response := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(response)

		r.GET("/v1/countries/:country_code/fetch", FetchFromAPI)
		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/v1/countries/mx/fetch",
			nil,
		)
		c.Request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(response, c.Request)
		So(response.Code, ShouldEqual, http.StatusOK)
	})
}

func TestGetFromCSVFile(t *testing.T) {
	Convey("Given a GET request to /countries/{country_code}", t, func() {
		response := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(response)

		r.GET("/v1/countries/:country_code", FetchFromAPI)
		c.Request, _ = http.NewRequest(
			http.MethodGet,
			"/v1/countries/mx",
			nil,
		)
		c.Request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(response, c.Request)
		So(response.Code, ShouldEqual, http.StatusOK)
	})
}
