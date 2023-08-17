package main

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/faizauthar12/skripsi/backend-service/controller"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	PORT = ":8080"
)

type application struct {
	user     *controller.UserController
	product  *controller.ProductController
	customer *controller.CustomerController
	cart     *controller.CartController
	order    *controller.OrderController
	admin    *controller.AdminController
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

func connectEth() *ethclient.Client {
	ethClient, errorConnectEthereum := ethclient.Dial("http://127.0.0.1:8545")

	if errorConnectEthereum != nil {
		panic(errorConnectEthereum)
	}

	return ethClient
}

func main() {

	errorLoadEnv := godotenv.Load()
	if errorLoadEnv != nil {
		panic(errorLoadEnv)
	}

	client := connect()
	clientEth := connectEth()

	ethPrivateKet := os.Getenv("ETH_PRIVATE_KEY")

	app := application{
		user:     &controller.UserController{Client: client},
		product:  &controller.ProductController{Client: client},
		customer: &controller.CustomerController{Client: client},
		cart:     &controller.CartController{Client: client},
		order: &controller.OrderController{
			Client:        client,
			ClientEth:     clientEth,
			EthPrivateKey: ethPrivateKet,
		},
		admin: &controller.AdminController{Client: client},
	}

	route := app.routes()
	route.Run(PORT)
}
