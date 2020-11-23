package countries

import (
	"github.com/gin-gonic/gin"
	"github.com/jegacs/golang-bootcamp-2020/handlers/v1/countries/handlers"
)

func Countries(group *gin.RouterGroup) {
	groupCountries := group.Group("countries/:country_code")
	groupCountries.GET("/fetch", handlers.FetchFromAPI)
	groupCountries.GET("", handlers.ReadCSVFromFile)
}
