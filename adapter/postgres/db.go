package postgres

import (
	"github.com/ncostamagna/go-sp-products/domain"
	"gorm.io/gorm"
)


type repository struct {
	db *gorm.DB
}

func NewRepository(dbGorm *gorm.DB) *repository {
	return &repository{db: dbGorm}
}

func (r *repository) Store(p domain.Product) (domain.Product, error) {
	if err := r.db.Create(&p).Error; err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]domain.Product) {
	var p []domain.Product
	if err := r.db.Find(&p).Error; err != nil {
		return []domain.Product{}
	}
	return p
}

func (r *repository) GetById(id string) (domain.Product, error) {
	var p domain.Product
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Product{}, ErrProductNotFound
		}
		return domain.Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id string, p domain.Product) (domain.Product, error) {
	result := r.db.Model(&p).Where("id = ?", id).Updates(&p)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.Product{}, ErrProductNotFound
	}
	return p, nil
}

func (r *repository) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&domain.Product{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}
