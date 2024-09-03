package app

import (
	"github.com/piresc/product-app/internal/entity"
)

func Init(config *entity.Config) {
	e := InitializeApp(config)
	e.SetupRoute()
	e.App.Start(":" + config.App.Port)
}
