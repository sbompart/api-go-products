# api-go-products
Api desarrollada con golang-gin y mongoDB

# api-products

Api CRUD desarrollada con golang-gin y mongoDB

# Arquitectura

Realizada con arquitectura hexagonal

docker run --privileged -it --rm api-products:latest /bin/sh

docker run -d -p 8080:8080 --name myapp-container api-products:latest

docker build  -t api-products:latest .