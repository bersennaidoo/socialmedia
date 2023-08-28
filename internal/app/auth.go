package app

import (
	"context"
	"crypto/sha256"
	"net/http"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *App) SignInHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h := sha256.New()

	ctx := context.Background()

	signinUser, err := a.US.SignIn(ctx, bson.M{
		"email":    user.Email,
		"password": string(h.Sum([]byte(user.Password))),
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("name", signinUser.Name)
	session.Set("token", sessionToken)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"token": session.Get("token"),
		"user": gin.H{
			"_id":   signinUser.ID,
			"name":  signinUser.Name,
			"email": signinUser.Email,
		}})
}

func (a *App) SignOutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed out...",
	})
}
