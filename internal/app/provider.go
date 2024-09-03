package app

import (
	"github.com/piresc/product-app/internal"

	"github.com/piresc/product-app/internal/entity"
	"github.com/piresc/product-app/internal/handler/http"
	repository "github.com/piresc/product-app/internal/repository/product"
	usecase "github.com/piresc/product-app/internal/usecase/product"

	"github.com/google/wire"
)

var (
	repoSet = wire.NewSet(
		provideProductRepo,
		wire.Bind(new(internal.ProductRepositoryItf), new(*repository.ProductRepo)),
	)

	ucSet = wire.NewSet(
		provideProductUsecase,
		wire.Bind(new(internal.ProductUsecaseItf), new(*usecase.ProductUsecase)),
	)

	allSet = wire.NewSet(
		repoSet,
		ucSet,
		provideProductApp,
	)
)

func provideProductRepo(configs *entity.Config) *repository.ProductRepo {
	return repository.NewProductRepository(configs)
}

func provideProductUsecase(repo internal.ProductRepositoryItf) *usecase.ProductUsecase {
	return usecase.NewProductUsecase(repo)
}

func provideProductApp(uc internal.ProductUsecaseItf) *http.RestApp {
	return http.NewProductApp(uc)
}
