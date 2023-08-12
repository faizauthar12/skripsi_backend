package cart

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

type Cart struct {
	UUID           string
	CustomerUUID   string
	CartGrandTotal int64
	CreatedAt      int64
	UpdatedAt      int64
}

// type Update
type UpdateCandidate struct {
	cart      Cart
	objective int
}

const (
	DATABASE   = "skripsi"
	COLLECTION = "cart"

	USER_UUID_CANNOT_BLANK        = "User UUID cannot be blank"
	PRODUCT_UUID_CANNOT_BLANK     = "Product UUID cannot be blank"
	CART_GRAND_TOTAL_CANNOT_BLANK = "Cart grand total cannot be blank"

	PRODUCT_NOT_FOUND            = "Product not found"
	CART_UUID_DOES_NOT_MATCH     = "Cart UUID does not match"
	CUSTOMER_UUID_DOES_NOT_MATCH = "Customer UUID does not match"

	FAILED_CREATE_CART = "Failed to create cart"
	FAILED_UPDATE_CART = "Failed to update the cart"

	// UpdateCandidate for Cart

	UPDATE_CART_GRAND_TOTAL = 2
	UPDATE_UPDATED_AT       = 4
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	customerUUID string,
	cartGrandTotal int64,
) (Cart, error) {

	uuid := uuid.New().String()

	if customerUUID == "" {
		return Cart{}, errors.New(USER_UUID_CANNOT_BLANK)
	}

	if cartGrandTotal < 0 {
		return Cart{}, errors.New(CART_GRAND_TOTAL_CANNOT_BLANK)
	}

	cart := Cart{
		UUID:           uuid,
		CustomerUUID:   customerUUID,
		CartGrandTotal: cartGrandTotal,
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
	}

	coll := connect(client)

	_, err := coll.InsertOne(context.TODO(), cart)

	if err != nil {
		return Cart{}, errors.New(FAILED_CREATE_CART)
	}

	return cart, nil
}

func Get(
	client *mongo.Client,
	customerUUID string,
) (Cart, bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "customeruuid", Value: customerUUID}}

	var findResult Cart
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		if findQueryError == mongo.ErrNoDocuments {
			return Cart{}, false, nil
		}
		return Cart{}, false, findQueryError
	}

	return findResult, true, nil
}

func GetManyByCustomerUUID(
	client *mongo.Client,
	customerUUID string,
	numItems int64,
	pages int64,
) ([]Cart, error) {

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
			return []Cart{}, nil
		}

		return []Cart{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []product
	var carts []Cart
	for cursor.Next(context.TODO()) {
		var cart Cart
		errorDecode := cursor.Decode(&cart)
		if errorDecode != nil {
			return nil, errorDecode
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func UpdateCartGrandTotal(
	updateList []UpdateCandidate,
	newCartGrandTotal int64,
) ([]UpdateCandidate, error) {

	if newCartGrandTotal < 0 {
		return nil, errors.New(CART_GRAND_TOTAL_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			cart:      Cart{CartGrandTotal: newCartGrandTotal},
			objective: UPDATE_CART_GRAND_TOTAL,
		},
	)

	return updateList, nil
}

func ExecUpdate(
	client *mongo.Client,
	updateList []UpdateCandidate,
	customerUUID string,
	cartUUID string,
) error {

	if cartUUID == "" {
		return errors.New(PRODUCT_UUID_CANNOT_BLANK)
	}

	if customerUUID == "" {
		return errors.New(USER_UUID_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			cart:      Cart{UpdatedAt: time.Now().Unix()},
			objective: UPDATE_UPDATED_AT,
		},
	)

	var setList bson.D
	var unsetList bson.D

	if len(updateList) > 0 {
		for i := 0; i < len(updateList); i++ {
			updateField := reflect.ValueOf(updateList[i].cart)

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

	cart, cartExist, errorFindCart := Get(
		client,
		cartUUID,
	)

	if !cartExist {
		if errorFindCart != nil {
			return errorFindCart
		} else {
			return errors.New(PRODUCT_NOT_FOUND)
		}
	}

	if cartUUID != cart.CustomerUUID { // Validate product ownership
		return errors.New(CART_UUID_DOES_NOT_MATCH)
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

	filterUpdate := bson.D{{Key: "uuid", Value: cartUUID}}

	_, err := coll.UpdateOne(context.TODO(), filterUpdate, update)

	if err != nil {
		return errors.New(FAILED_UPDATE_CART)
	}

	return nil
}

func Delete(
	client *mongo.Client,
	customerUUID string,
	cartUUID string,
) (bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: cartUUID}}

	var findResult Cart
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		return false, errors.New(PRODUCT_NOT_FOUND)
	}

	// Validate product ownership
	if customerUUID != findResult.CustomerUUID {
		return false, errors.New(CUSTOMER_UUID_DOES_NOT_MATCH)
	}

	_, errorDelete := coll.DeleteOne(context.TODO(), filter)
	if errorDelete != nil {
		return false, errorDelete
	}

	return true, nil
}

func DeleteByCustomerUUID(
	client *mongo.Client,
	customerUUID string,
) error {

	coll := connect(client)

	filter := bson.D{{Key: "customeruuid", Value: customerUUID}}

	_, errorDeleteBulk := coll.DeleteMany(context.TODO(), filter)
	if errorDeleteBulk != nil {
		return errorDeleteBulk
	}

	return nil
}
