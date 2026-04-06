package product

import "github.com/ncostamagna/go-sp-products/domain"

type Repository interface {
	Store(p domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product)
	GetById(id string) (domain.Product, error)
	Update(id string, p domain.Product) error
	Delete(id string) error
}