package main

import (
	"github.com/gemm123/go-mongo/controllers"
	"github.com/gemm123/go-mongo/database"
	"github.com/gemm123/go-mongo/services"
	"github.com/gin-gonic/gin"
)

const uri = "mongodb://localhost:27017"

func main() {
	client, err := database.ConnectMongo(uri)
	if err != nil {
		panic(err)
	}
	defer database.Disconnect(client)

	userCollection := client.Database("go-mongo").Collection("users")
	service := services.NewUserService(userCollection)
	controller := controllers.NewUserController(service)

	r := gin.Default()

	r.POST("/users", controller.PostUser)
	r.GET("/users", controller.GetAllUser)
	r.GET("/users/:name", controller.GetUserByName)
	r.PUT("/users/:name", controller.UpdateUser)
	r.DELETE("users/:name", controller.DeleteUser)

	r.Run()
}
