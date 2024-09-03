package main

import (
	"github.com/piresc/product-app/cmd/config"
	"github.com/piresc/product-app/internal/app"
)

func main() {
	appName := "product-app"
	configs := config.InitConfig(appName)

	app.Init(configs)

}
