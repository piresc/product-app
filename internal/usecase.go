package internal

import (
	"github.com/piresc/product-app/internal/entity"
)

type ProductUsecaseItf interface {
	Create(product *entity.Product) error
	GetAll() ([]entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	GetByVariant(variant string) ([]entity.Product, error)
	GetByTitle(title string) ([]entity.Product, error)
	Update(product *entity.Product) error
	Delete(id uint) (*entity.Product, error)
}
