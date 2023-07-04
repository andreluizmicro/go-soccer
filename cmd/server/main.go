package main

import "github.com/andreluizmicro/go-soccer/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBName)
}
