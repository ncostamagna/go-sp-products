package product

import (
	"errors"
	"github.com/ncostamagna/go-sp-products/domain"
	"github.com/ncostamagna/go-sp-products/adapter/postgres"
	"github.com/google/uuid"
)


type Service interface {
	Store(p domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product)
	GetById(id string) (domain.Product, error)
	Update(id string, p domain.Product) error
	Delete(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Store(p domain.Product) (domain.Product, error) {
	if p.Name == nil || *p.Name == "" {
		return domain.Product{}, ErrNameRequired
	}
	if p.Price == nil {
		return domain.Product{}, ErrPriceRequired
	}
	if *p.Price < 0 {
		return domain.Product{}, ErrPriceNegative
	}
	p.ID = uuid.NewString()
	
	return s.repo.Store(p)

}

func (s *service) GetAll() ([]domain.Product) {
	return s.repo.GetAll()
}

func (s *service) GetById(id string) (domain.Product, error) {
	p, err := s.repo.GetById(id)
	    if err != nil {
			if errors.Is(err, postgres.ErrProductNotFound) {
				return domain.Product{}, ErrProductNotFound
			}
			return domain.Product{}, err
		}
		return p, nil
}

func (s *service) Update(id string, p domain.Product) error {
	if id == "" {
		return ErrIdRequired
	}

	if p.Name != nil && *p.Name == "" {
		return ErrNameRequired
	}
	if p.Price != nil && *p.Price < 0 {
		return ErrPriceNegative
	} 
	return s.repo.Update(id, p)
}

func (s *service) Delete(id string) error {
	if id == "" {
		return ErrIdRequired
	}
	return s.repo.Delete(id)
}