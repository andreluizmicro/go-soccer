package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/helloworld", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// v1.GET("/books", controllers.FindBooks)
		// v1.POST("/books", controllers.CreateBook)
		// v1.GET("/books/:id", controllers.FindBook)
		// v1.PUT("/books/:id", controllers.UpdateBook)
		// v1.DELETE("/books/:id", controllers.DeleteBook)
	}
}
