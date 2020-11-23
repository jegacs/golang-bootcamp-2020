package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jegacs/golang-bootcamp-2020/models"
)

func FetchFromAPI(c *gin.Context) {
	countryCode := c.Param("country_code")
	country := models.NewCountry(countryCode)
	err := country.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error : " + err.Error()})
		return
	}
	err = country.Store()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, country)
}

func ReadCSVFromFile(c *gin.Context) {
	countryCode := c.Param("country_code")
	country := models.NewCountry(countryCode)
	err := country.ReadFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error : " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, country)
}
