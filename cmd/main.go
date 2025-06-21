package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	handlers "api-products/cmd/api/handlers"
	repository "api-products/internal/repository/mongo/products"
	service "api-products/internal/service/products"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnvConfig struct {
	Port   string
	DbName string
	DbURI  string
}

func main() {

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	/* // router := gin.Default()

	fmt.Printf("Variables de entorno, PORT: %v URI: %v DB: %v \n", cfg.Port, cfg.DbURI, cfg.DbName)

	const uri = "mongodb://saul:saul123@localhost:27017/?maxPoolSize=2&authSource=test1"
	const dbName = "test1" */

	clientOptions := options.Client().ApplyURI(cfg.DbURI)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error al conectarse a MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error al hacer ping a MongoDB: %v", err)
	}

	repo := repository.NewRepositoryMongo(client, cfg.DbName, "products")

	router := gin.Default()
	router.GET("/health", health)

	serviceProduct := service.NewProductService(repo)
	handlers.NewProductHandler(serviceProduct, router)

	log.Fatalln(router.Run(cfg.Port))

}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Ok"})
}

func LoadConfig() (*EnvConfig, error) {
	config := &EnvConfig{
		Port:   ":8080",
		DbName: "test1",
		DbURI:  "mongodb://saul:saul123@localhost:27017/?maxPoolSize=2&authSource=test1",
	}
	if config.DbName == "" {
		return nil, errors.New("DATABASE_NAME is required")
	}
	if config.Port == "" {
		return nil, errors.New("PORT is required")
	}
	if config.DbURI == "" {
		return nil, errors.New("MONGO_URI is required")
	}
	return config, nil
}
