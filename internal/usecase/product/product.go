package usecase

import (
	"github.com/piresc/product-app/internal"

	"github.com/piresc/product-app/internal/entity"
)

type ProductUsecase struct {
	productRepo internal.ProductRepositoryItf
}

func NewProductUsecase(
	productRepo internal.ProductRepositoryItf,
) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

func (uc *ProductUsecase) Create(product *entity.Product) error {
	return uc.productRepo.Create(product)
}

func (uc *ProductUsecase) GetAll() ([]entity.Product, error) {
	return uc.productRepo.GetAll()
}

func (uc *ProductUsecase) GetByID(id uint) (*entity.Product, error) {
	return uc.productRepo.GetByID(id)
}

func (uc *ProductUsecase) GetByVariant(variant string) ([]entity.Product, error) {
	return uc.productRepo.GetByVariant(variant)
}

func (uc *ProductUsecase) GetByTitle(variant string) ([]entity.Product, error) {
	return uc.productRepo.GetByTitle(variant)
}

func (uc *ProductUsecase) Update(product *entity.Product) error {
	return uc.productRepo.Update(product)
}

func (uc *ProductUsecase) Delete(id uint) (*entity.Product, error) {
	return uc.productRepo.Delete(id)
}
