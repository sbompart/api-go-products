package products

import (
	"api-products/internal/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepositoryMongo) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
