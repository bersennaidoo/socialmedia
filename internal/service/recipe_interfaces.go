package service

import (
	"context"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

/*type RecipeInterface interface {
	ListRecipes(ctx context.Context, bs bson.M) ([]domain.Recipe, error)
	NewRecipe(ctx context.Context, recipe domain.Recipe) (domain.Recipe, error)
	UpdateRecipe(ctx context.Context, bs bson.M, bd bson.D) error
	DeleteRecipe(ctx context.Context, bs bson.M) error
	GetRecipe(ctx context.Context, bs bson.M) (domain.Recipe, error)
}*/

/*type RecipeRedisInterface interface {
	GetRecipes(recipe string) (string, error)
	SetRecipes(recipe string, r string, t time.Duration)
	DeleteRecipes(recipes string)
}*/

type UserInterface interface {
	//SignIn(ctx context.Context, bs bson.M) error
	CreateUser(ctx context.Context, bs bson.M) error
	ListUser(ctx context.Context, bs bson.M) ([]domain.User, error)
	UpdateUser(ctx context.Context, bs bson.M, bd bson.D) error
}
