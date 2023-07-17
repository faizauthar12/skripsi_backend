package main

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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

func main() {

}
