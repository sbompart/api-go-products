package products

import "api-products/internal/domain"

func (s ProductService) GetAll() ([]domain.Product, error) {
	products, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
