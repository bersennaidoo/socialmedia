package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) RunApi(addr string) {
	router := gin.Default()

	//store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("recipes_api", store))

	//router.GET("/recipes", a.ListRecipesHandler)
	//router.POST("/signin", a.SignInHandler)
	//router.POST("/refresh", a.RefreshHandler)
	//router.POST("/signout", a.SignOutHandler)
	router.GET("/api/users", a.ListUserHandler)
	router.POST("/api/users", a.CreateUserHandler)
	router.GET("/api/users/:userId", a.UserByIdHandler)
	router.PUT("/api/users/:userId", a.UpdateUserHandler)
	router.DELETE("/api/users/:userId", a.DeleteUserHandler)

	/*authorized := router.Group("/")
	authorized.Use(a.AuthMiddleware())
	{

		authorized.POST("/recipes", a.NewRecipeHandler)
		authorized.PUT("/recipes/:id", a.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", a.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", a.GetRecipeHandler)
		//authorized.GET("/recipes/search", a.SearchRecipesHandler)
	}*/

	router.Run(addr)
}
