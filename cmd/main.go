package main

import (
	"log"

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

	v1.Routes(r)

	if err := r.Run(":" + config.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
