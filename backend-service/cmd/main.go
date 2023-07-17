package main

import (
	"context"
	"time"

	User "github.com/faizauthar12/skripsi/user-gomod"

	"github.com/faizauthar12/skripsi/backend-service/controller"
	"github.com/faizauthar12/skripsi/backend-service/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	userInfo *User.User
)

const (
	PORT = ":8080"
)

type application struct {
	user        *controller.UserController
	product     *controller.ProductController
	middlewares *middlewares.Middlewares
}

func connect() *mongo.Client {
	const URI = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, errConnect := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	if errConnect != nil {
		panic(errConnect)
	}

	if errPing := client.Ping(ctx, readpref.Primary()); errPing != nil {
		panic(errPing)
	}

	return client
}

func main() {

	client := connect()

	app := application{
		user: &controller.UserController{Client: client},
		product: &controller.ProductController{
			Client:   client,
			UserInfo: userInfo,
		},
		middlewares: &middlewares.Middlewares{UserInfo: userInfo},
	}

	route := app.routes()
	route.Run(PORT)
}
