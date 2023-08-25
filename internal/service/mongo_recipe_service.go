package service

import (
	"context"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeService struct {
	MC *mongo.Client
}

func NewRecipeService(mc *mongo.Client) *RecipeService {
	return &RecipeService{
		MC: mc,
	}
}

func (rs *RecipeService) ListRecipes(ctx context.Context, bs bson.M) ([]domain.Recipe, error) {
	collection := rs.MC.Database("demo").Collection("recipes")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	recipes := make([]domain.Recipe, 0)
	for cur.Next(ctx) {
		var recipe domain.Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (rs *RecipeService) NewRecipe(ctx context.Context, recipe domain.Recipe) (domain.Recipe, error) {
	collection := rs.MC.Database("demo").Collection("recipes")
	_, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		return domain.Recipe{}, err
	}

	return domain.Recipe{}, nil

}

func (rs *RecipeService) UpdateRecipe(ctx context.Context, bs bson.M, bd bson.D) error {
	collection := rs.MC.Database("demo").Collection("recipes")

	_, err := collection.UpdateOne(ctx, bs, bd)
	if err != nil {
		return err
	}

	return nil

	return nil
}

func (rs *RecipeService) DeleteRecipe(ctx context.Context, bs bson.M) error {
	collection := rs.MC.Database("demo").Collection("recipes")

	_, err := collection.DeleteOne(ctx, bs)
	if err != nil {
		return err
	}

	return nil
}

func (rs *RecipeService) GetRecipe(ctx context.Context, bs bson.M) (domain.Recipe, error) {
	collection := rs.MC.Database("demo").Collection("recipes")
	cur := collection.FindOne(ctx, bs)

	var recipe domain.Recipe
	err := cur.Decode(&recipe)
	if err != nil {
		return domain.Recipe{}, err
	}

	return recipe, nil
}
