package products

import (
	"api-products/internal/domain"
	"fmt"
)

func (s ProductService) Create(product domain.Product) (id interface{}, err error) {
	// insertId, err := s.Repo.Create(&product)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Printf("PRODUCT: %v", product)
	id = "12312341"
	return id, nil
}
