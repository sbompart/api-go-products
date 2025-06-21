package products

import (
	"context"
	"fmt"

	"api-products/internal/domain"
)

func (r *RepositoryMongo) Insert(product *domain.Product) (interface{}, error) {

	insertResult, err := r.Collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, fmt.Errorf("Error al crear un producto: %w", err)
	}

	return insertResult.InsertedID, nil
}
