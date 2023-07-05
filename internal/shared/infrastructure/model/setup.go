package model

import (
	"database/sql"
	"fmt"

	"github.com/andreluizmicro/go-soccer/configs"
)

var DB *sql.DB

func ConnectDatabase() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	dbDriver := config.DBDriver
	dbName := config.DBName
	dbHost := config.DBHost
	dbPort := config.DBPort
	dbuser := config.DBUser
	dbPassword := config.DBPassword

	db, err := sql.Open(fmt.Sprintf("%s", dbDriver), fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbuser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()
}
