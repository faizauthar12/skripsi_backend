package cartitem

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartItem struct {
	UUID              string
	CustomerUUID      string
	ProductUUID       string
	ProductName       string
	ProductPic        string
	ProductQuantity   int64
	ProductPrice      int64
	ProductTotalPrice int64
	CreatedAt         int64
	UpdatedAt         int64
}

type UpdateCandidate struct {
	cartItem  CartItem
	objective int
}

const (
	DATABASE   = "skripsi"
	COLLECTION = "cartitem"

	CART_UUID_CANNOT_BLANK        = "Cart UUID cannot be blank"
	PRODUCT_UUID_CANNOT_BLANK     = "Product UUID cannot be blank"
	CART_TOTAL_PRICE_CANNOT_BLANK = "Cart total price cannot be blank"
	CUSTOMER_UUID_CANNOT_BLANK    = "Customer UUID cannot be blank"
	CART_ITEM_UUID_CANNOT_BLANK   = "Cart item UUID cannot be blank"

	CART_ITEM_NOT_FOUND     = "Cart item not found"
	FAILED_UPDATE_CART_ITEM = "Failed to update the cart item"

	QUANTITY_NOT_VALID      = "Product Quantity is not valid"
	CANNOT_CREATE_CART_ITEM = "Cannot create cart item"

	// UpdateCandidate for CartItem

	UPDATE_CART_ITEM_PRODUCT_QUANTITY    = 3
	UPDATE_CART_ITEM_PRODUCT_TOTAL_PRICE = 4
	UPDATE_CART_ITEM_PRODUCT_UPDATED_AT  = 6
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	customerUUID string,
	productUUID string,
	productPrice int64,
	productQuantity int64,
) (CartItem, error) {

	uuid := uuid.New().String()

	if customerUUID == "" {
		return CartItem{}, errors.New(CART_UUID_CANNOT_BLANK)
	}

	if productUUID == "" {
		return CartItem{}, errors.New(PRODUCT_UUID_CANNOT_BLANK)
	}

	if productQuantity < 0 {
		return CartItem{}, errors.New(QUANTITY_NOT_VALID)
	}

	productTotalPrice := productPrice * productQuantity

	cartItem := CartItem{
		UUID:              uuid,
		CustomerUUID:      customerUUID,
		ProductUUID:       productUUID,
		ProductQuantity:   productQuantity,
		ProductTotalPrice: productTotalPrice,
	}

	coll := connect(client)
	_, err := coll.InsertOne(context.TODO(), cartItem)

	if err != nil {
		return CartItem{}, errors.New(CANNOT_CREATE_CART_ITEM)
	}

	return cartItem, nil
}

func Get(
	client *mongo.Client,
	cartItemUUID string,
) (CartItem, bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: cartItemUUID}}

	var findResult CartItem
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		if findQueryError == mongo.ErrNoDocuments {
			return CartItem{}, false, nil
		}
		return CartItem{}, false, findQueryError
	}

	return findResult, true, nil
}

func GetManyByCustomerUUID(
	client *mongo.Client,
	customerUUID string,
	numItems int64,
	pages int64,
) ([]CartItem, error) {

	coll := connect(client)

	filter := bson.D{{Key: "customeruuid", Value: customerUUID}}
	opts := options.Find().SetLimit(numItems).SetSkip((pages - 1) * numItems)

	cursor, errorFindQuery := coll.Find(
		context.TODO(),
		filter,
		opts,
	)

	if errorFindQuery != nil {
		if errorFindQuery == mongo.ErrNoDocuments {
			return []CartItem{}, nil
		}

		return []CartItem{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []cartItem
	var cartItems []CartItem
	for cursor.Next(context.TODO()) {
		var cartItem CartItem
		errorDecode := cursor.Decode(&cartItem)
		if errorDecode != nil {
			return nil, errorDecode
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func UpdateProductQuantity(
	updateList []UpdateCandidate,
	newProductQuantity int64,
) ([]UpdateCandidate, error) {

	if newProductQuantity < 0 {
		return nil, errors.New(QUANTITY_NOT_VALID)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			cartItem:  CartItem{ProductQuantity: newProductQuantity},
			objective: UPDATE_CART_ITEM_PRODUCT_QUANTITY,
		},
	)

	return updateList, nil
}

func UpdateProducTotalPrice(
	updateList []UpdateCandidate,
	newProductTotalPrice int64,
) ([]UpdateCandidate, error) {

	if newProductTotalPrice < 0 {
		return nil, errors.New(CART_TOTAL_PRICE_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			cartItem:  CartItem{ProductQuantity: newProductTotalPrice},
			objective: UPDATE_CART_ITEM_PRODUCT_TOTAL_PRICE,
		},
	)

	return updateList, nil
}

func ExecUpdate(
	client *mongo.Client,
	updateList []UpdateCandidate,
	customerUUID string,
	cartItemUUID string,
) error {

	if cartItemUUID == "" {
		return errors.New(CART_ITEM_UUID_CANNOT_BLANK)
	}

	if customerUUID == "" {
		return errors.New(CUSTOMER_UUID_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			cartItem:  CartItem{UpdatedAt: time.Now().Unix()},
			objective: UPDATE_CART_ITEM_PRODUCT_UPDATED_AT,
		},
	)

	var setList bson.D
	var unsetList bson.D

	if len(updateList) > 0 {
		for i := 0; i < len(updateList); i++ {
			updateField := reflect.ValueOf(updateList[i].cartItem)

			fieldName := strings.ToLower(updateField.Type().Field(updateList[i].objective).Name)
			value := updateField.Field(updateList[i].objective).Interface()

			if value == "" {
				unsetList = append(unsetList, bson.E{fieldName, ""})
			} else {
				setList = append(setList, bson.E{fieldName, value})
			}
		}
	}

	coll := connect(client)

	cart, cartItemExist, errorFindCartItem := Get(
		client,
		cartItemUUID,
	)

	if !cartItemExist {
		if errorFindCartItem != nil {
			return errorFindCartItem
		} else {
			return errors.New(CART_ITEM_NOT_FOUND)
		}
	}

	if cartItemUUID != cart.CustomerUUID { // Validate product ownership
		return errors.New(CART_ITEM_NOT_FOUND)
	}

	var update bson.D

	// conclusive
	if len(unsetList) > 0 && len(setList) > 0 {
		update = bson.D{{"$set", setList}, {"$unset", unsetList}}
	} else if len(setList) > 0 {
		update = bson.D{{"$set", setList}}
	} else {
		update = bson.D{{"$unset", unsetList}}
	}

	filterUpdate := bson.D{{Key: "uuid", Value: cartItemUUID}}

	_, err := coll.UpdateOne(context.TODO(), filterUpdate, update)

	if err != nil {
		return errors.New(FAILED_UPDATE_CART_ITEM)
	}

	return nil
}
