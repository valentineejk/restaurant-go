package controllers

import (
	"context"
	"net/http"
	"restaurant-go/database"
	"restaurant-go/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		// timmeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// get the foodid
		foodId := c.Param("food_id")

		// foodmodel
		var food models.Food

		//mongo food collectioon, find foood with this id
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)

		//defer
		defer cancel()

		//check error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "no data found",
				"data":    nil,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "no data found",
			"data":    food,
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
