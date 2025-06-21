package products

import "api-products/internal/domain"

func (s ProductService) Get(id string) (product *domain.Product, err error) {
	// product, err := s.Repo.GetByID() //Corregir
	// if err != nil {
	// 	return nil, err
	// }
	product = &domain.Product{
		ID:   "123455",
		Name: "Taza",
	}
	return product, nil
}
