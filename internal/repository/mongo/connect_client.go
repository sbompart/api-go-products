package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbUri string) (client *mongo.Client, err error) {
	clientOptions := options.Client().ApplyURI(dbUri)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Error al conectarse a MongoDB: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Error al hacer ping a MongoDB: %v", err)
		return nil, err
	}
	return client, nil
}
