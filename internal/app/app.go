package app

import (
	"github.com/bersennaidoo/socialmedia/internal/service"
)

type App struct {
	//RS service.RecipeInterface
	US service.UserInterface
	//RC service.RecipeRedisInterface
}

func NewApp(us service.UserInterface) *App {
	return &App{
		US: us,
	}
}
