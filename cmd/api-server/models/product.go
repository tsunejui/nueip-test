package models

import (
	pkgModel "nueip/pkg/lib/model"
	"time"
)

const defaultModelProductVersion = 1

type Product struct {
	pkgModel.Header
	Name        string    `gorm:"index;not null" json:"name" validate:"required"`
	Cost        int       `gorm:"not null;default:0" json:"cost" validate:"required"`
	Price       int       `gorm:"not null;default:0" json:"price" validate:"required"`
	Description string    `json:"description"`
	State       int       `gorm:"not null;default:0" json:"state"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

func (dao *DAO) GetProducts() ([]*Product, error) {
	var products []*Product
	if err := dao.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (dao *DAO) GetProduct(id uint32) (*Product, error) {
	p := &Product{}
	if err := dao.DB.Where("id = ?", id).First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (dao *DAO) CreateProduct(product *Product) (*Product, error) {
	product.Version = defaultModelProductVersion
	if err := dao.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (dao *DAO) UpdateProduct(id uint32, product *Product) (*Product, error) {
	p, err := dao.GetProduct(id)
	if err != nil {
		return nil, err
	}
	p.Version = defaultModelProductVersion
	p.Name = product.Name
	p.Cost = product.Cost
	p.Price = product.Price
	p.Description = product.Description
	p.State = product.State
	p.StartTime = product.StartTime
	p.EndTime = product.EndTime
	if err := dao.DB.Save(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (dao *DAO) DeleteProduct(id uint32) error {
	if err := dao.DB.Where("id = ?", id).Delete(Product{}).Error; err != nil {
		return err
	}
	return nil
}
