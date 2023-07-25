package order

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/faizauthar12/skripsi/order-gomod/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Order struct {
	ID                  int64 // Auto-incremented ID
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

type IDGenerator struct {
	ID   string `bson:"_id"`
	Next int64  `bson:"next"`
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

	privateKey = "8675ae8a067522d06bc7c78e6e84c945fb70329042bd04d86ec38a05de372008"
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

	// Obtain the next available ID from the IDGenerator collection
	idGenColl := client.Database(DATABASE).Collection("orderID")
	update := bson.D{{"$inc", bson.D{{"next", 1}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var idDoc IDGenerator
	errorGenColl := idGenColl.FindOneAndUpdate(context.TODO(), bson.D{}, update, opts).Decode(&idDoc)
	if errorGenColl != nil {
		return Order{}, errors.New(ERROR_CREATING_DB)
	}

	order := Order{
		ID:                  idDoc.Next,
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

func StoreDataToEth(order Order, orderContract *api.Api, auth *bind.TransactOpts) (*types.Transaction, error) {
	var productUUID []string
	var productQuantity []int64
	var productTotalPrice []int64
	for _, item := range order.CartItem {
		productUUID = append(productUUID, item.ProductUUID)
		productQuantity = append(productQuantity, item.ProductQuantity)
		productTotalPrice = append(productTotalPrice, item.ProductTotalPrice)
	}

	tx, errorStoringData := orderContract.SetOrder(
		auth,
		string(order.ID),
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
		return &types.Transaction{}, nil
	}

	return tx, nil

}

func ReadDataFromEth(address common.Address, orderContract *api.Api, auth *bind.TransactOpts, orderUUID string) {

	fmt.Println("Read: From: ", auth.From)
	fmt.Println("Read: Nonce: ", auth.Nonce)

	result, errorResult := orderContract.GetOrder(&bind.CallOpts{}, orderUUID)

	if errorResult != nil {
		panic(errorResult)
	}

	id, _ := strconv.ParseInt(result.OrderID, 10, 64)

	retrivedOrder := Order{
		ID:                  id,
		CartGrandTotal:      result.CartGrandTotal,
		CustomerUUID:        result.CustomerUUID,
		CustomerName:        result.CustomerName,
		CustomerEmail:       result.CustomerEmail,
		CustomerAddress:     result.CustomerAddress,
		CustomerPhoneNumber: result.CustomerPhoneNumber,
		Status:              result.Status,
	}

	fmt.Println("Read: Result: ", result)
	fmt.Println("Read: retrievedOrder: ", retrivedOrder)
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

// func main() {

// 	cartItem := CartItem{
// 		ProductUUID:       "asdasdasd",
// 		ProductQuantity:   2,
// 		ProductTotalPrice: 24000000,
// 	}

// 	// mongodb
// 	clientMongo := connectMongo()

// 	order, errorCreateOrder := Create(
// 		clientMongo,
// 		[]CartItem{cartItem},
// 		24000000,
// 		"asdasdasdasd",
// 		"faiz",
// 		"faiz@gmail.com",
// 		"disini sana sini",
// 		"085155054855",
// 	)

// 	if errorCreateOrder != nil {
// 		panic(errorCreateOrder)
// 	}

// 	ethClient, errorConnectEthereum := ethclient.Dial("HTTP://127.0.0.1:8545")
// 	if errorConnectEthereum != nil {
// 		panic(errorConnectEthereum)
// 	}

// 	auth := getAccountAuth(ethClient, privateKey)

// 	address, tx, instance, errorDeploy := api.DeployApi(auth, ethClient)
// 	if errorDeploy != nil {
// 		panic(errorDeploy)
// 	}

// 	fmt.Println("main: address->", address.Hex())

// 	_, _ = instance, tx
// 	fmt.Println("main: instance->", instance)
// 	fmt.Println("main: tx->", tx.Hash().Hex())

// 	orderContract, errorConnectEth := api.NewApi(
// 		common.HexToAddress(address.Hex()),
// 		ethClient,
// 	)

// 	if errorConnectEth != nil {
// 		panic(errorConnectEth)
// 	}

// 	balancedAt, errorGetBalance := ethClient.BalanceAt(context.TODO(), common.HexToAddress(address.Hex()), nil)

// 	if errorGetBalance != nil {
// 		panic(errorGetBalance)
// 	}

// 	fmt.Println("main: balancedAt: ", balancedAt)

// 	tx, errorStoreData := StoreDataToEth(order, orderContract, auth)
// 	if errorStoreData != nil {
// 		panic(errorStoreData)
// 	}

// 	fmt.Println("main: tx of stored data: ", tx)

// 	ReadDataFromEth(common.HexToAddress(address.Hex()), orderContract, auth, string(rune(order.ID)))
// }
