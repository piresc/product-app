package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/piresc/product-app/internal/entity"
)

func (ra *RestApp) Create(c echo.Context) error {
	var product entity.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := ra.productUC.Create(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	return c.JSON(http.StatusCreated, product)
}

func (ra *RestApp) GetAll(c echo.Context) error {
	products, err := ra.productUC.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
	}
	return c.JSON(http.StatusOK, products)
}

func (ra *RestApp) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	product, err := ra.productUC.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func (ra *RestApp) GetByVariant(c echo.Context) error {
	variant := c.QueryParam("variant")
	if variant == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Variant cannot be empty"})
	}
	products, err := ra.productUC.GetByVariant(variant)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Products not found"})
	}
	return c.JSON(http.StatusOK, products)
}

func (ra *RestApp) GetByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title cannot be empty"})
	}
	products, err := ra.productUC.GetByTitle(title)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Products not found"})
	}
	return c.JSON(http.StatusOK, products)
}

func (ra *RestApp) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var product entity.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	product.ID = uint(id) // Set the ID for the product to update
	if err := ra.productUC.Update(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}
	return c.JSON(http.StatusOK, product)
}

func (ra *RestApp) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	product, err := ra.productUC.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}
