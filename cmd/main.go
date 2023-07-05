package main

import (
	"log"

	"github.com/andreluizmicro/go-soccer/cmd/container"
	"github.com/andreluizmicro/go-soccer/configs/database"
	countryRoutesV1 "github.com/andreluizmicro/go-soccer/internal/country/infrastructure/http/api/v1/routes"
	"github.com/andreluizmicro/go-soccer/internal/country/infrastructure/repository"
	"github.com/andreluizmicro/go-soccer/internal/country/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	go database.ConnectDatabase()
	defer database.DB.Close()

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	repo := repository.NewCountryRepository()
	container.Testando(usecase.NewCreateCountryUseCase(repo))
	countryRoutesV1.Routes(r)

	if err := r.Run(":" + "9000"); err != nil {
		log.Fatal(err)
	}
}
