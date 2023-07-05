package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	usecase "github.com/andreluizmicro/go-soccer/internal/shared/usercase/country"

	"github.com/gin-gonic/gin"
)

func CreateCountry(c *gin.Context) {
	var dto usecase.CountryInputDTO
	err := json.NewDecoder(c.Request.Body).Decode(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "houve um erro ao tentar realizar a request",
		})
	}

	fmt.Println(dto.Name)
}
