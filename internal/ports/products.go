package ports

import "api-products/internal/domain"

type ProductRepository interface {
	Create(product domain.Product) (interface{}, error)
	GetAll() ([]domain.Product, error)
}
