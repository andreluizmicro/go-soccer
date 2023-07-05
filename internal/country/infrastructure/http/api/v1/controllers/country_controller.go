package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andreluizmicro/go-soccer/cmd/container"
	usecase "github.com/andreluizmicro/go-soccer/internal/country/usecase"

	"github.com/gin-gonic/gin"
)

type CountryController struct {
}

func (controller *CountryController) NewCountryController() *CountryController {
	return &CountryController{}
}

func CreateCountry(c *gin.Context) {
	var input usecase.CountryInputDTO
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	output, err := container.CountryService.Execute(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"name":          output.Name,
		"capital":       output.Capital,
		"area":          output.Area,
		"language":      output.Language,
		"timezone":      output.Timezone,
		"continent":     output.Continent,
		"offical_color": output.OfficalColor,
		"population":    output.Population,
	})
}
