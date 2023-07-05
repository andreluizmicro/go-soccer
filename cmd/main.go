package main

import (
	"log"

	"github.com/andreluizmicro/go-soccer/configs"
	countryRoutesV1 "github.com/andreluizmicro/go-soccer/internal/country/infrastructure/http/api/v1/routes"
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

	countryRoutesV1.Routes(r)

	if err := r.Run(":" + config.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
