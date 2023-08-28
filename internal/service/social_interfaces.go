package service

import (
	"context"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type UserInterface interface {
	SignIn(ctx context.Context, bs bson.M) (domain.User, error)
	CreateUser(ctx context.Context, bs bson.M) error
	ListUser(ctx context.Context, bs bson.M) ([]domain.User, error)
	UpdateUser(ctx context.Context, bs bson.M, bd bson.D) error
	UserById(ctx context.Context, bs bson.M) (domain.User, error)
	DeleteUser(ctx context.Context, bs bson.M) error
}
