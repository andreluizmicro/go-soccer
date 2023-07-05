package main

import (
	"log"
	"net/http"

	"github.com/andreluizmicro/go-soccer/configs"
	"github.com/andreluizmicro/go-soccer/database"
	v1 "github.com/andreluizmicro/go-soccer/internal/manegment/infrastructure/http/api/v1"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.NewRoutes(r)

	if err := r.Run(":" + config.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
