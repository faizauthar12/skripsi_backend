package order

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/faizauthar12/skripsi/order-gomod/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	EthAddress          *types.Transaction
	Status              string
}

type CartItem struct {
	ProductUUID       string
	ProductName       string
	ProductPic        string
	ProductQuantity   int64
	ProductPrice      int64
	ProductTotalPrice int64
}

type IDGenerator struct {
	ID   string `bson:"_id"`
	Next int64  `bson:"next"`
}

type DurianPayOrderItems struct {
	Name  string `json:"name"`
	Qty   int64  `json:"qty"`
	Price string `json:"price"`
	Logo  string `json:"logo"`
}

type DurianPayOrderCustomer struct {
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	GivenName    string `json:"given_name"`
	AddressLine1 string `json:"address_line_1"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postal_code"`
}

type DurianPayOrderSandboxOptions struct {
	ForceFail bool  `json:"force_fail"`
	Delay     int64 `json:"delay_ms"`
}

type DurianPayOrder struct {
	Amount                string                       `json:"amount"`
	Currency              string                       `json:"currency"`
	IsPaymentLink         bool                         `json:"is_payment_link"`
	IsLive                bool                         `json:"is_live"`
	IsNotificationEnabled bool                         `json:"is_notification_enabled"`
	SandboxOptions        DurianPayOrderSandboxOptions `json:"sandbox_options"`
	Customer              DurianPayOrderCustomer       `json:"customer"`
	Items                 []DurianPayOrderItems        `json:"items"`
}

type DurianPayorderResponseData struct {
	PaymentLinkURL string `json:"payment_link_url"`
}
type DurianPayOrderResponseBody struct {
	Data DurianPayorderResponseData `json:"data"`
}

const (
	MONGO_ERROR                        = "order: MongoError: "
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

func New(
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

	order := Order{
		CartItem:            cartItem,
		CartGrandTotal:      cartGrandTotal,
		CustomerUUID:        customerUUID,
		CustomerName:        customerName,
		CustomerEmail:       customerEmail,
		CustomerAddress:     customerAddress,
		CustomerPhoneNumber: customerPhoneNumber,
	}

	return order, nil
}

func DeployApi(auth *bind.TransactOpts, ethClient *ethclient.Client) (common.Address, error) {

	address, tx, instance, errorDeploy := api.DeployApi(auth, ethClient)
	if errorDeploy != nil {
		return common.Address{}, errorDeploy
	}

	fmt.Println("DeployApi: address->", address.Hex())
	fmt.Println("DeployApi: instance->", instance)
	fmt.Println("DeployApi: tx->", tx.Hash().Hex())

	return address, nil
}

func NewApi(address common.Address, ethClient *ethclient.Client) (*api.Api, error) {
	orderContract, errorConnectEth := api.NewApi(
		common.HexToAddress(address.Hex()),
		ethClient,
	)

	if errorConnectEth != nil {
		return nil, errorConnectEth
	}

	return orderContract, nil
}

func CheckBalance(ethClient *ethclient.Client, address common.Address) (*big.Int, error) {
	balancedAt, errorGetBalance := ethClient.BalanceAt(context.TODO(), common.HexToAddress(address.Hex()), nil)

	if errorGetBalance != nil {
		return nil, errorGetBalance
	}

	fmt.Println("CheckBalance: balancedAt: ", balancedAt)
	return balancedAt, nil
}

// function to create auth for any account from its private key
func GetAccountAuth(ethClient *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

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
		fmt.Sprint(order.ID),
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

func AddEthAddress(
	order *Order,
	ethAddress *types.Transaction,
) {
	order.EthAddress = ethAddress
}

func Create(
	client *mongo.Client,
	order *Order,
) error {
	coll := connect(client)

	// Obtain the next available ID from the IDGenerator collection
	idGenColl := client.Database(DATABASE).Collection("orderID")
	update := bson.D{{"$inc", bson.D{{"next", 1}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var idDoc IDGenerator
	errorGenColl := idGenColl.FindOneAndUpdate(context.TODO(), bson.D{}, update, opts).Decode(&idDoc)
	if errorGenColl != nil {
		return errors.New(ERROR_CREATING_DB)
	}

	for {
		order.ID = idDoc.Next

		_, errorCreatingOrder := coll.InsertOne(context.TODO(), order)

		if errorCreatingOrder == nil {
			return nil
		}

		if monggoError, ok := errorCreatingOrder.(mongo.WriteException); ok {
			for _, writeException := range monggoError.WriteErrors {
				if writeException.Code == 11000 {
					// Duplicate key error, rerun generateBookingID and try again
					continue
				}
			}
		}

		// if the error is not duplicate key
		return errors.New(MONGO_ERROR + errorCreatingOrder.Error())
	}
}

func GetMany(
	client *mongo.Client,
	numItems int64,
	pages int64,
) ([]Order, error) {

	coll := connect(client)

	filter := bson.D{{}}
	opts := options.Find().SetLimit(numItems).SetSkip((pages - 1) * numItems)

	cursor, errorFindQuery := coll.Find(
		context.TODO(),
		filter,
		opts,
	)

	if errorFindQuery != nil {
		if errorFindQuery == mongo.ErrNoDocuments {
			return []Order{}, nil
		}

		return []Order{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []customer
	var orders []Order
	for cursor.Next(context.TODO()) {
		var order Order
		errorDecode := cursor.Decode(&order)
		if errorDecode != nil {
			return nil, errorDecode
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetCount(
	client *mongo.Client,
) (int64, error) {

	coll := connect(client)

	counts, errorGetCount := coll.CountDocuments(context.TODO(), bson.D{})

	if errorGetCount != nil {
		if errorGetCount == mongo.ErrNilDocument {
			return 0, nil
		}

		return 0, errorGetCount
	}

	return counts, nil
}

func NewPaymentDurianPay(order Order) DurianPayOrder {

	durian := DurianPayOrder{
		Amount:                strconv.FormatInt(order.CartGrandTotal, 10),
		Currency:              "IDR",
		IsPaymentLink:         true,
		IsLive:                false,
		IsNotificationEnabled: true,
		SandboxOptions: DurianPayOrderSandboxOptions{
			ForceFail: false,
			Delay:     1,
		},
		Customer: DurianPayOrderCustomer{
			Email:        order.CustomerEmail,
			Mobile:       order.CustomerPhoneNumber,
			GivenName:    order.CustomerName,
			AddressLine1: order.CustomerAddress,
		},
	}

	return durian
}

func AddItemsDurianPay(durianPayOrder *DurianPayOrder, order *Order) {

	var items []DurianPayOrderItems
	for _, cartItem := range order.CartItem {
		item := DurianPayOrderItems{
			Name:  cartItem.ProductName,
			Qty:   cartItem.ProductQuantity,
			Price: strconv.FormatInt(cartItem.ProductPrice, 10),
			Logo:  cartItem.ProductPic,
		}

		items = append(items, item)
	}

	durianPayOrder.Items = items
}

func CreatePaymentDurianPay(durianPayOrder DurianPayOrder, DurianPayServerKey string) (string, error) {

	client := &http.Client{}

	jsonBody, errorJsonBody := json.Marshal(durianPayOrder)
	if errorJsonBody != nil {
		fmt.Println("Error converting to json: ", errorJsonBody)
		return "", errorJsonBody
	}

	bodyBuffer := bytes.NewBuffer(jsonBody)

	request, errorRequest := http.NewRequest("POST", "https://api.durianpay.id/v1/orders", bodyBuffer)
	if errorRequest != nil {
		fmt.Println("Error Creating request: ", errorRequest)
		return "", nil
	}

	// Set the Authorization header with the basic auth
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(DurianPayServerKey+":")))
	request.Header.Set("Content-Type", "application/json")

	response, errorDoRequest := client.Do(request)

	if response.StatusCode != http.StatusCreated {
		fmt.Println("Response DurianPayment: ", response.Body)
		return "", errorDoRequest
	}

	var durianPayOrderResponseBody DurianPayOrderResponseBody

	errorDecodeJson := json.NewDecoder(response.Body).Decode(&durianPayOrderResponseBody)
	if errorDecodeJson != nil {
		fmt.Println("Error decoding response: ", errorDecodeJson)
		return "", nil
	}

	return fmt.Sprintf("https://links.durianpay.id/payment/%s", durianPayOrderResponseBody.Data.PaymentLinkURL), nil
}

// func connectMongo() *mongo.Client {
// 	const URI = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, errConnect := mongo.Connect(ctx, options.Client().ApplyURI(URI))

// 	if errConnect != nil {
// 		panic(errConnect)
// 	}

// 	if errPing := client.Ping(ctx, readpref.Primary()); errPing != nil {
// 		panic(errPing)
// 	}

// 	return client
// }

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
