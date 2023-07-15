package cartitem

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartItem struct {
	UUID              string
	CartUUID          string
	ProductUUID       string
	ProductQuantity   int64
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

	// UpdateCandidate for CartItem

	UPDATE_CART_ITEM_PRODUCT_QUANTITY    = 3
	UPDATE_CART_ITEM_PRODUCT_TOTAL_PRICE = 4
	UPDATE_CART_ITEM_PRODUCT_UPDATED_AT  = 6
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func create(
	client *mongo.Client,
	cartUUID string,
	productUUID string,
	productPrice int64,
	productQuantity int64,
) (CartItem, error) {

	uuid := uuid.New().String()

	if cartUUID == "" {
		return CartItem{}, errors.New("cart uuid cannot be blank")
	}

	if productUUID == "" {
		return CartItem{}, errors.New("product uuid cannot be blank")
	}

	if productQuantity < 0 {
		return CartItem{}, errors.New("invalid quantity")
	}

	productTotalPrice := productPrice * productQuantity

	cartItem := CartItem{
		UUID:              uuid,
		CartUUID:          cartUUID,
		ProductUUID:       productUUID,
		ProductQuantity:   productQuantity,
		ProductTotalPrice: productTotalPrice,
	}

	coll := connect(client)
	_, err := coll.InsertOne(context.TODO(), cartItem)

	if err != nil {
		return CartItem{}, errors.New("error creating cart item")
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

func GetManyByUserUUID(
	client *mongo.Client,
	cartUUID string,
	numItems int64,
	pages int64,
) ([]CartItem, error) {

	coll := connect(client)

	filter := bson.D{{Key: "cartuuid", Value: cartUUID}}
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
