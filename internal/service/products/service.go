package products

import (
	"api-products/internal/domain"
	"api-products/internal/ports"
)

type ProductService struct {
	Repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(product domain.Product) (interface{}, error) {
	id, err := s.Repo.Create(product)
	return id, err
}

func (s *ProductService) GetProduct(id string) ([]domain.Product, error) {
	return s.Repo.GetAll()
}

/* func (s *ProductService) UpdateProduct(product *domain.Product) error {
	return s.Repo.Update(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.Delete(id)
}
*/
