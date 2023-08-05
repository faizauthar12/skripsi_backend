package product

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

type Product struct {
	UUID               string
	UserUUID           string
	UserName           string
	ProductName        string
	ProductDescription string
	ProductCategory    string
	ProductPrice       int64
	ProductStock       int64
	CreatedAt          int64
	UpdatedAt          int64
}

type UpdateCandidate struct {
	product   Product
	objective int
}

const (
	DATABASE   = "skripsi"
	COLLECTION = "product"

	UUID_DOESNT_MATCH = "Product: Unauthorized Update: UserUUID doesn't match"
	PRODUCT_NOT_FOUND = "Product: Product not found"

	UPDATE_USER_NAME           = 2
	UPDATE_PRODUCT_NAME        = 3
	UPDATE_PRODUCT_DESCRIPTION = 4
	UPDATE_PRODUCT_CATEGORY    = 5
	UPDATE_PRODUCT_PRICE       = 6
	UPDATE_PRODUCT_STOCK       = 7
	UPDATE_UPDATED_AT          = 9
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	userUUID string,
	userName string,
	name string,
	description string,
	category string,
	price int64,
	quantity int64,
) (Product, error) {

	uuid := uuid.New().String()

	if userUUID == "" {
		return Product{}, errors.New("userUUID cannot be blank")
	}

	if userName == "" {
		return Product{}, errors.New("user name cannot be blank")
	}

	if name == "" {
		return Product{}, errors.New("name cannot be blank")
	}

	if description == "" {
		return Product{}, errors.New("description cannot be blank")
	}

	if category == "" {
		return Product{}, errors.New("category cannot be blank")
	}

	if price < 0 {
		return Product{}, errors.New("price cannot be blank")
	}

	if quantity < 0 {
		return Product{}, errors.New("quantity cannot be blank")
	}

	product := Product{
		UUID:               uuid,
		UserUUID:           userUUID,
		UserName:           userName,
		ProductName:        name,
		ProductDescription: description,
		ProductCategory:    category,
		ProductPrice:       price,
		ProductStock:       quantity,
		CreatedAt:          time.Now().Unix(),
		UpdatedAt:          time.Now().Unix(),
	}

	coll := connect(client)
	_, err := coll.InsertOne(context.TODO(), product)

	if err != nil {
		return Product{}, errors.New("error creating product")
	}

	return product, nil
}

func Get(
	client *mongo.Client,
	productUUID string,
) (Product, bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: productUUID}}

	var findResult Product
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		if findQueryError == mongo.ErrNoDocuments {
			return Product{}, false, nil
		}
		return Product{}, false, findQueryError
	}

	return findResult, true, nil
}

func GetMany(
	client *mongo.Client,
	category string,
	numItems int64,
	pages int64,
) ([]Product, error) {

	coll := connect(client)

	var filter bson.D
	if category != "" {
		filter = bson.D{{Key: "productcategory", Value: category}}
	} else {
		filter = bson.D{}
	}
	opts := options.Find().SetLimit(numItems).SetSkip((pages - 1) * numItems)

	cursor, errorFindQuery := coll.Find(
		context.TODO(),
		filter,
		opts,
	)

	if errorFindQuery != nil {
		if errorFindQuery == mongo.ErrNoDocuments {
			return []Product{}, nil
		}

		return []Product{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []product
	var products []Product
	for cursor.Next(context.TODO()) {
		var product Product
		errorDecode := cursor.Decode(&product)
		if errorDecode != nil {
			return nil, errorDecode
		}
		products = append(products, product)
	}

	return products, nil
}

func GetManyByUserName(
	client *mongo.Client,
	userName string,
	numItems int64,
	pages int64,
) ([]Product, error) {

	coll := connect(client)

	filter := bson.D{{Key: "username", Value: userName}}
	opts := options.Find().SetLimit(numItems).SetSkip((pages - 1) * numItems)

	cursor, errorFindQuery := coll.Find(
		context.TODO(),
		filter,
		opts,
	)

	if errorFindQuery != nil {
		if errorFindQuery == mongo.ErrNoDocuments {
			return []Product{}, nil
		}

		return []Product{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []product
	var products []Product
	for cursor.Next(context.TODO()) {
		var product Product
		errorDecode := cursor.Decode(&product)
		if errorDecode != nil {
			return nil, errorDecode
		}
		products = append(products, product)
	}

	return products, nil
}

func UpdateName(
	updateList []UpdateCandidate,
	newProductName string,
) ([]UpdateCandidate, error) {

	if newProductName == "" {
		return nil, errors.New("product name cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{ProductName: newProductName},
			objective: UPDATE_PRODUCT_NAME,
		},
	)

	return updateList, nil
}

func UpdateDescription(
	updateList []UpdateCandidate,
	newProductDescription string,
) ([]UpdateCandidate, error) {

	if newProductDescription == "" {
		return nil, errors.New("product description cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{ProductDescription: newProductDescription},
			objective: UPDATE_PRODUCT_DESCRIPTION,
		},
	)

	return updateList, nil
}

func UpdateCategory(
	updateList []UpdateCandidate,
	newProductCategory string,
) ([]UpdateCandidate, error) {

	if newProductCategory == "" {
		return nil, errors.New("product category cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{ProductCategory: newProductCategory},
			objective: UPDATE_PRODUCT_CATEGORY,
		},
	)

	return updateList, nil
}

func UpdatePrice(
	updateList []UpdateCandidate,
	newProductPrice int64,
) ([]UpdateCandidate, error) {

	if newProductPrice < 0 {
		return nil, errors.New("product category cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{ProductPrice: newProductPrice},
			objective: UPDATE_PRODUCT_PRICE,
		},
	)

	return updateList, nil
}

func UpdateStock(
	updateList []UpdateCandidate,
	newProductStock int64,
) ([]UpdateCandidate, error) {

	if newProductStock < 0 {
		return nil, errors.New("product stock cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{ProductStock: newProductStock},
			objective: UPDATE_PRODUCT_STOCK,
		},
	)

	return updateList, nil
}

func ExecUpdate(
	client *mongo.Client,
	updateList []UpdateCandidate,
	userUUID string,
	productUUID string,
) error {

	if productUUID == "" {
		return errors.New("product UUID cannot be blank")
	}

	if userUUID == "" {
		return errors.New("user UUID cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			product:   Product{UpdatedAt: time.Now().Unix()},
			objective: UPDATE_UPDATED_AT,
		},
	)

	var setList bson.D
	var unsetList bson.D

	if len(updateList) > 0 {
		for i := 0; i < len(updateList); i++ {
			updateField := reflect.ValueOf(updateList[i].product)

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

	product, productExist, errorFindProduct := Get(
		client,
		productUUID,
	)

	if !productExist {
		if errorFindProduct != nil {
			return errorFindProduct
		} else {
			return errors.New(PRODUCT_NOT_FOUND)
		}
	}

	if productUUID != product.UserUUID { // Validate product ownership
		return errors.New(UUID_DOESNT_MATCH)
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

	filterUpdate := bson.D{{Key: "uuid", Value: productUUID}}

	_, err := coll.UpdateOne(context.TODO(), filterUpdate, update)

	if err != nil {
		return errors.New("update product failed")
	}

	return nil
}

func Delete(
	client *mongo.Client,
	userUUID string,
	productUUID string,
) (bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: productUUID}}

	var findResult Product
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		return false, errors.New(PRODUCT_NOT_FOUND)
	}

	// Validate product ownership
	if userUUID != findResult.UserUUID {
		return false, errors.New(UUID_DOESNT_MATCH)
	}

	_, errorDelete := coll.DeleteOne(context.TODO(), filter)
	if errorDelete != nil {
		return false, errorDelete
	}

	return true, nil
}

func DeleteByUseruuid(
	client *mongo.Client,
	userUUID string,
) error {

	coll := connect(client)

	filter := bson.D{{Key: "useruuid", Value: userUUID}}

	_, errorDeleteBulk := coll.DeleteMany(context.TODO(), filter)
	if errorDeleteBulk != nil {
		return errorDeleteBulk
	}

	return nil
}
