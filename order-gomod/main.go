package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/faizauthar12/skripsi/order-gomod/api"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Order struct {
	UUID                string
	CartItem            []CartItem
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}

type CartItem struct {
	ProductUUID       string
	ProductQuantity   int64
	ProductTotalPrice int64
}

const (
	CART_ITEM_CANNOT_BLANK             = "cart item cannot be blank"
	CART_GRAND_TOTAL_CANNOT_BLANK      = "cart grand total cannot be blank"
	CUSTOMER_UUID_CANNOT_BLANK         = "customer uuid cannot be blank"
	CUSTOMER_NAME_CANNOT_BLANK         = "customer name cannot be blank"
	CUSTOMER_EMAIL_CANNOT_BLANK        = "customer email cannot be blank"
	CUSTOMER_ADDRESS_CANNOT_BLANK      = "customer address cannot be blank"
	CUSTOMER_PHONE_NUMBER_CANNOT_BLANK = "customer phone number cannot blank"

	ERROR_CREATING_DB = "error creating order"

	DATABASE   = "skripsi"
	COLLECTION = "order"
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	cartItem []CartItem,
	cartGrandTotal int64,
	customerUUID string,
	customerName string,
	customerEmail string,
	customerAddress string,
	customerPhoneNumber string,
) (Order, error) {

	uuid := uuid.New().String()

	if cartItem == nil {
		return Order{}, errors.New(CART_ITEM_CANNOT_BLANK)
	}

	if cartGrandTotal < 0 {
		return Order{}, errors.New(CART_GRAND_TOTAL_CANNOT_BLANK)
	}

	if customerUUID == "" {
		return Order{}, errors.New(CUSTOMER_UUID_CANNOT_BLANK)
	}

	if customerName == "" {
		return Order{}, errors.New(CUSTOMER_NAME_CANNOT_BLANK)
	}

	if customerAddress == "" {
		return Order{}, errors.New(CUSTOMER_ADDRESS_CANNOT_BLANK)
	}

	if customerPhoneNumber == "" {
		return Order{}, errors.New(CUSTOMER_PHONE_NUMBER_CANNOT_BLANK)
	}

	order := Order{
		UUID:                uuid,
		CartItem:            cartItem,
		CartGrandTotal:      cartGrandTotal,
		CustomerUUID:        customerUUID,
		CustomerName:        customerName,
		CustomerEmail:       customerEmail,
		CustomerAddress:     customerAddress,
		CustomerPhoneNumber: customerPhoneNumber,
	}

	coll := connect(client)

	_, err := coll.InsertOne(context.TODO(), order)

	if err != nil {
		return Order{}, errors.New(ERROR_CREATING_DB)
	}

	return order, nil
}

// function to create auth for any account from its private key
func getAccountAuth(ethClient *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, errorGetPrivateGet := crypto.HexToECDSA(privateKeyAddress)
	if errorGetPrivateGet != nil {
		fmt.Println("errorGetPrivateKey")
		panic(errorGetPrivateGet)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("publicKeyECDSA")
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(context.TODO(), fromAddress)
	if err != nil {
		fmt.Println("nonce")
		panic(err)
	}
	fmt.Println("nounce=", nonce)

	gasPrice, errorSuggesGasPrice := ethClient.SuggestGasPrice(context.TODO())
	if errorSuggesGasPrice != nil {
		log.Fatal(errorSuggesGasPrice)
	}

	chainID, err := ethClient.ChainID(context.TODO())
	if err != nil {
		fmt.Println("chainID")
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		fmt.Println("Auth")
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth
}

const privateKey = "69f5d51743cb8451258fa506b16f0c70dab6473477afa85c3f0ddf7bbe7ede0e"
const accountAddress = "0x184987B87f7fDe80Ca9eF42911AAF6F4D9BE0537"

func StoreDataToEth(order Order, orderContract *api.Api, auth *bind.TransactOpts) {
	var productUUID []string
	var productQuantity []int64
	var productTotalPrice []int64
	for _, item := range order.CartItem {
		productUUID = append(productUUID, item.ProductUUID)
		productQuantity = append(productQuantity, item.ProductQuantity)
		productTotalPrice = append(productTotalPrice, item.ProductTotalPrice)
	}

	tx, errorStoringData := orderContract.StoreOrder(
		auth,
		order.UUID,
		productUUID,
		productQuantity,
		productTotalPrice,
		order.CartGrandTotal,
		order.CustomerUUID,
		order.CustomerName,
		order.CustomerEmail,
		order.CustomerAddress,
		order.CustomerPhoneNumber,
		order.Status,
	)

	if errorStoringData != nil {
		panic(errorStoringData)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

}

func ReadDataFromEth(orderContract *api.Api, auth *bind.TransactOpts) {
	// readData, errorReadData := connectionEthereum.OrderInstance(
	// 	&bind.CallOpts{
	// 		Pending:     false,
	// 		From:        auth.From,
	// 		BlockNumber: auth.Nonce,
	// 		Context:     context.TODO(),
	// 	},
	// )

	// if errorReadData != nil {
	// 	panic(errorReadData)
	// }

	// retrivedOrder := Order{
	// 	UUID:                readData.UUID,
	// 	CartGrandTotal:      readData.CartGrandTotal,
	// 	CustomerUUID:        readData.CustomerUUID,
	// 	CustomerName:        readData.CustomerName,
	// 	CustomerEmail:       readData.CustomerEmail,
	// 	CustomerAddress:     readData.CustomerAddress,
	// 	CustomerPhoneNumber: readData.CustomerPhoneNumber,
	// 	Status:              readData.Status,
	// }

	// fmt.Println(retrivedOrder)

	// fmt.Println("Read: From: ", auth.From)
	// fmt.Println("Read: Nonce: ", auth.Nonce)

	opts := &bind.CallOpts{
		Pending:     false,
		From:        auth.From,
		BlockNumber: auth.Nonce,
		Context:     context.TODO(),
	}

	result, errorResult := orderContract.OrderInstance(opts)

	if errorResult != nil {
		panic(errorResult)
	}

	fmt.Println(result)
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

func main() {

	cartItem := CartItem{
		ProductUUID:       "asdasdasd",
		ProductQuantity:   2,
		ProductTotalPrice: 24000000,
	}

	// mongodb
	clientMongo := connectMongo()

	order, errorCreateOrder := Create(
		clientMongo,
		[]CartItem{cartItem},
		24000000,
		"asdasdasdasd",
		"faiz",
		"faiz@gmail.com",
		"disini sana sini",
		"085155054855",
	)

	if errorCreateOrder != nil {
		panic(errorCreateOrder)
	}

	ethClient, errorConnectEthereum := ethclient.Dial("HTTP://127.0.0.1:7545")
	if errorConnectEthereum != nil {
		panic(errorConnectEthereum)
	}

	auth := getAccountAuth(ethClient, privateKey)

	// address, tx, instance, errorDeploy := api.DeployApi(auth, ethClient)
	// if errorDeploy != nil {
	// 	panic(errorDeploy)
	// }

	// fmt.Println(address.Hex())
	// _, _ = instance, tx
	// fmt.Println("instance->", instance)
	// fmt.Println("tx->", tx.Hash().Hex())

	commonAddress := common.HexToAddress(accountAddress)

	orderContract, errorConnectEth := api.NewApi(
		commonAddress,
		ethClient,
	)

	if errorConnectEth != nil {
		panic(errorConnectEth)
	}

	balancedAt, errorGetBalance := ethClient.BalanceAt(context.TODO(), commonAddress, nil)

	if errorGetBalance != nil {
		panic(errorGetBalance)
	}

	fmt.Println(balancedAt)

	StoreDataToEth(order, orderContract, auth)

	// ReadDataFromEth(orderContract, auth)
}
