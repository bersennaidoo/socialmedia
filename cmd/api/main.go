package main

import (
	"context"
	"log"
	"time"

	"github.com/bersennaidoo/socialmedia/internal/app"
	"github.com/bersennaidoo/socialmedia/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	/*redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	fmt.Printf("REdis Ping %v\n", status)*/

	//rrs := service.NewRecipeRedisService(redisClient)
	us := service.NewUserService(client)
	//rs := service.NewRecipeService(client)
	ap := app.NewApp(us)
	ap.RunApi(":8080")
}
