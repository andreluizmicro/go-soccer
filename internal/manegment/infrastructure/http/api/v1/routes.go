package v1

import (
	"net/http"

	"github.com/andreluizmicro/go-soccer/internal/manegment/infrastructure/http/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/countries", controllers.CreateCountry)
	}
}
