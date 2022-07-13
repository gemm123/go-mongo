package main

import (
	"github.com/gemm123/go-mongo/controllers"
	"github.com/gemm123/go-mongo/database"
	"github.com/gin-gonic/gin"
)

const uri = "mongodb://localhost:27017"

func main() {
	client, err := database.ConnectMongo(uri)
	if err != nil {
		panic(err)
	}
	defer database.Disconnect(client)

	controller := controllers.NewController(client)

	r := gin.Default()
	r.POST("/users", controller.PostUser)
	r.Run()
}
