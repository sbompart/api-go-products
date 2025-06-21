package products

import (
	"api-products/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryMongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// Create implements ports.ProductRepository.
func (r *RepositoryMongo) Create(product domain.Product) (interface{}, error) {
	panic("unimplemented")
}

// func NewRepositoryMongo(uriMongo, dbName string) *RepositoryMongo {
// 	client, err := repoMongo.ConnectClient(uriMongo)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	collection := client.Database(dbName).Collection("products")

// 	return &RepositoryMongo{
// 		Collection: collection,
// 		Client:     client,
// 	}
// }

func NewRepositoryMongo(client *mongo.Client, dbName, collectionName string) *RepositoryMongo {
	return &RepositoryMongo{
		Client:     client,
		Collection: client.Database(dbName).Collection(collectionName),
	}
}
