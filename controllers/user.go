package controllers

import (
	"context"
	"net/http"

	"github.com/gemm123/go-mongo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type controller struct {
	client *mongo.Client
}

func NewController(client *mongo.Client) *controller {
	return &controller{
		client: client,
	}
}

// func (ctr *controller) GetUser(c *gin.Context) {
// 	name := c.Param("name")

// }

func (ctr *controller) PostUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	user.ID = primitive.NewObjectID()

	coll := ctr.client.Database("go-mongo").Collection("users")

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success insert user"})
}
