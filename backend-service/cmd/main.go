package main

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/faizauthar12/skripsi/backend-service/controller"
	"github.com/go-playground/form/v4"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	PORT = ":8080"
)

type application struct {
	controller *controller.Controller
}

func connectMongo() *mongo.Client {
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

	clientMongo := connectMongo()
	clientEth := connectEth()

	ethPrivateKey := os.Getenv("ETH_PRIVATE_KEY")
	durianPayAuth := os.Getenv("DURIANPAY_AUTH")

	formDecoder := form.NewDecoder()

	app := application{
		controller: &controller.Controller{
			ClientMongo:   clientMongo,
			ClientEth:     clientEth,
			EthPrivateKey: ethPrivateKey,
			DurianPayAuth: durianPayAuth,
			FormDecoder:   formDecoder,
		},
	}

	route := app.routes()
	route.Run(PORT)
}
