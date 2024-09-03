package http

import (
	"github.com/piresc/product-app/internal"

	"github.com/labstack/echo/v4"
)

type RestApp struct {
	App       *echo.Echo
	productUC internal.ProductUsecaseItf
}

func NewProductApp(
	productUC internal.ProductUsecaseItf,
) *RestApp {
	e := echo.New()
	return &RestApp{
		App:       e,
		productUC: productUC,
	}
}

// Versioning in REST APIs is important to maintain backward compatibility,
// allowing changes to the API without breaking older client implementations.
// It provides flexibility for evolving the API while ensuring stability for existing users.
// SetupRoute is a function to setup route
func (ra *RestApp) SetupRoute() {
	v1 := ra.App.Group("/api/v1")
	v1.POST("/products", ra.Create)
	v1.GET("/products", ra.GetAll)
	v1.GET("/products/:id", ra.GetByID)
	v1.GET("/products/variant", ra.GetByVariant)
	v1.GET("/products/title", ra.GetByTitle)
	v1.PUT("/products/:id", ra.Update)
	v1.DELETE("/products/:id", ra.Delete)

}
