package app

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *App) CreateUserHandler(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h := sha256.New()

	ctx := context.Background()

	err := a.US.CreateUser(ctx, bson.M{
		"name":      user.Name,
		"email":     user.Email,
		"updatedAt": time.Now(),
		"password":  string(h.Sum([]byte(user.Password))),
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SignUp Successfull",
	})

}

func (a *App) ListUserHandler(c *gin.Context) {

	ctx := context.Background()
	users, err := a.US.ListUser(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (a *App) UpdateUserHandler(c *gin.Context) {
	id := c.Param("userId")
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := context.Background()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("%v", err)
	}

	filter := bson.M{"_id": objectId}
	update := bson.D{{"$set", bson.D{{"email", user.Email}}}}

	err = a.US.UpdateUser(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been updated",
	})
}
