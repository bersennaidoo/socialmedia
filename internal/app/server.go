package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
)

func (a *App) RunApi(addr string) {
	router := gin.Default()

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("users_api", store))

	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/signin", a.SignInHandler)
	router.GET("/signout", a.SignOutHandler)
	router.GET("/api/users", a.ListUserHandler)
	router.POST("/api/users", a.CreateUserHandler)

	authorized := router.Group("/api/")
	authorized.Use(a.AuthMiddleware())
	{

		authorized.GET("/users/:userId", a.UserByIdHandler)
		authorized.PUT("/users/:userId", a.UpdateUserHandler)
		authorized.DELETE("/users/:userId", a.DeleteUserHandler)
	}

	router.Use(static.ServeRoot("/", "./ui/socialmedia/build"))

	router.Run(addr)
}
