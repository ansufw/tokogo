package main

import (
	"github.com/aysf/tokogo/config"
	"github.com/aysf/tokogo/middlewares"
	"github.com/aysf/tokogo/routes"
)

func main() {

	config.InitDB()

	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
