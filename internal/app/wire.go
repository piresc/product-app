//go:build wireinject
// +build wireinject

package app

import (
	"github.com/piresc/product-app/internal/entity"
	"github.com/piresc/product-app/internal/handler/http"

	"github.com/google/wire"
)

// Wire ensures proper dependency injection and helps maintain clean, manageable code.
// It prevents cyclic dependency problem as the codebase grows more complex in the future.
// InitializeApp initializes the application with all dependencies
func InitializeApp(cfg *entity.Config) *http.RestApp {
	wire.Build(allSet)
	return &http.RestApp{}
}
