package controller

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/form/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DEFAULT_NUM_ITEMS int64 = 10
	DEFAULT_PAGES     int64 = 1

	SERVER_MALFUNCTION_CANNOT_CREATE_USER               = "Server malfunction, cannot create user"
	SERVER_MALFUNCTION_CANNOT_CREATE_TOKEN              = "Server malfunction, cannot create token"
	SERVER_MALFUNCTION_CANNOT_UPDATE_USER               = "Server malfunction, cannot update user"
	SERVER_MALFUNCTION_CANNOT_CREATE_GET                = "Server malfunction, cannot get product"
	SERVER_MALFUNCTION_CANNOT_CREATE_PRODUCT            = "Server malfunction, cannot create product"
	SERVER_MALFUNCTION_CANNOT_UPDATE_PRODUCT            = "Server malfunction, cannot update product"
	SERVER_MALFUNCTION_CANNOT_DELETE_PRODUCT            = "Server malfunction, cannot delete product"
	SERVER_MALFUNCTION_CANNOT_GET_PRODUCT               = "Server malfunction, cannot get product"
	SERVER_MALFUNCTION_CANNOT_CREATE_CART_ITEM          = "Server malfunction, cannot create cart item"
	SERVER_MALFUNCTION_CANNOT_STORE_SMART_CONTRACT      = "Server malfunction, cannot store smart contract"
	SERVER_MALFUNCTION_CANNOT_INITIALIZE_SMART_CONTRACT = "Server malfunction, cannot initialize smart contract"
	SERVER_MALFUNCTION_CANNOT_CREATE_CUSTOMER           = "Server malfunction, cannot create customer"
	SERVER_MALFUNCTION_CANNOT_GET_MANY_CART_ITEM        = "Server malfunction, cannot get many cart item"
	SERVER_MALFUNCTION_CANNOT_CREATE_CART               = "Server malfunction, cannot create cart"
	SERVER_MALFUNCTION_CANNOT_CREATE_ORDER              = "Server malfunction, cannot create order"
	SERVER_MALFUNCTION_CANNOT_CREATE_PAYMENT            = "Server malfunction, cannot create durian pay payment"

	UNAUTHORIZED = "Unauthorized"

	SUCCESS_CREATE_USER     = "Successfully create user"
	SUCCESS_LOGIN_USER      = "Successfully logged in"
	SUCCESS_UPDATE_USER     = "Successfully update user"
	SUCCESS_CREATE_PRODUCT  = "Successfully create product"
	SUCCESS_UPDATE_PRODUCT  = "Successfully update the product"
	SUCCESS_DELETE_SERVICE  = "Successfully delete the product"
	SUCCESS_GET_PRODUCT     = "Successfully get product"
	SUCCESS_DELETE_PRODUCT  = "Successfully delete the product"
	SUCCESS_CREATE_CART     = "Successfully create cart"
	SUCCESS_CREATE_CUSTOMER = "Successfully create customer"
	SUCCESS_CREATE_ORDER    = "Successfully create order"

	NOTHING_TO_UPDATE = "There is no item to be updated"

	USER_NOT_FOUND = "email or password not found"
)

type Controller struct {
	ClientMongo   *mongo.Client
	ClientEth     *ethclient.Client
	EthPrivateKey string
	DurianPayAuth string
	FormDecoder   *form.Decoder
}
