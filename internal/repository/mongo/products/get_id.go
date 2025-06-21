package products

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"api-products/internal/domain"
)

func (r *RepositoryMongo) GetByID(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	return &product, err
}
