package service

import (
	"context"
	"log"

	"github.com/bersennaidoo/socialmedia/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	MC *mongo.Client
}

func NewUserService(mc *mongo.Client) *UserService {
	return &UserService{
		MC: mc,
	}
}

/*func (us *UserService) SignIn(ctx context.Context, bs bson.M) error {
	collection := us.MC.Database("demo").Collection("users")

	cur := collection.FindOne(ctx, bs)
	if cur.Err() != nil {
		return cur.Err()
	}
	return nil
}*/

func (us *UserService) CreateUser(ctx context.Context, bs bson.M) error {
	collection := us.MC.Database("social").Collection("users")

	_, err := collection.InsertOne(ctx, bs)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) ListUser(ctx context.Context, bs bson.M) ([]domain.User, error) {
	collection := us.MC.Database("social").Collection("users")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	users := make([]domain.User, 0)
	for cur.Next(ctx) {
		var user domain.User
		cur.Decode(&user)
		user.Password = ""
		users = append(users, user)
	}

	return users, nil
}

func (us *UserService) UpdateUser(ctx context.Context, bs bson.M, bd bson.D) error {
	collection := us.MC.Database("social").Collection("users")

	_, err := collection.UpdateOne(ctx, bs, bd)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
