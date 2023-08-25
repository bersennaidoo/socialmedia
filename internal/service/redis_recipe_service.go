package service

import (
	"time"

	"github.com/go-redis/redis"
)

type RecipeRedisService struct {
	RC *redis.Client
}

func NewRecipeRedisService(rc *redis.Client) *RecipeRedisService {
	return &RecipeRedisService{
		RC: rc,
	}
}

func (rrs *RecipeRedisService) GetRecipes(recipe string) (string, error) {
	val, err := rrs.RC.Get(recipe).Result()
	if err == redis.Nil {
		return "", err
	}

	return val, nil
}

func (rrs *RecipeRedisService) SetRecipes(recipe string, r string, t time.Duration) {
	rrs.RC.Set(recipe, r, t)
}

func (rrs *RecipeRedisService) DeleteRecipes(recipes string) {
	rrs.RC.Del(recipes)
}
