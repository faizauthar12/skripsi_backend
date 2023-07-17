package customer

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

type Customer struct {
	UUID        string
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	CreatedAt   int64
	UpdatedAt   int64
}

type UpdateCandidate struct {
	customer  Customer
	objective int
}

const (
	DATABASE   = "skripsi"
	COLLECTION = "customer"

	UPDATE_NAME        = 1
	UPDATE_EMAIL       = 2
	UPDATE_ADDRESS     = 3
	UPDATE_PHONENUMBER = 4
	UPDATE_UPDATED_AT  = 6
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	customerName string,
	customerEmail string,
	customerAddress string,
	customerPhoneNumber string,
) (Customer, error) {

	uuid := uuid.New().String()

	if customerName == "" {
		return Customer{}, errors.New("customer name cannot be blank")
	}

	if customerEmail == "" {
		return Customer{}, errors.New("customer email cannot be blank")
	}

	if customerAddress == "" {
		return Customer{}, errors.New("customer address cannot be blank")
	}

	if customerPhoneNumber == "" {
		return Customer{}, errors.New("customer phone address cannot be blank")
	}

	customer := Customer{
		UUID:        uuid,
		Name:        customerName,
		Email:       customerEmail,
		Address:     customerAddress,
		PhoneNumber: customerPhoneNumber,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	coll := connect(client)

	_, err := coll.InsertOne(context.TODO(), customer)

	if err != nil {
		return Customer{}, errors.New("error creating customer")
	}

	return customer, nil
}

func Get(
	client *mongo.Client,
	customerUUID string,
) (Customer, bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: customerUUID}}

	var findResult Customer
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		if findQueryError == mongo.ErrNoDocuments {
			return Customer{}, false, nil
		}
		return Customer{}, false, findQueryError
	}

	return findResult, true, nil
}

func GetMany(
	client *mongo.Client,
	numItems int64,
	pages int64,
) ([]Customer, error) {

	coll := connect(client)

	// filter := bson.D{{Key: "username", Value: userName}}
	opts := options.Find().SetLimit(numItems).SetSkip((pages - 1) * numItems)

	cursor, errorFindQuery := coll.Find(
		context.TODO(),
		opts,
	)

	if errorFindQuery != nil {
		if errorFindQuery == mongo.ErrNoDocuments {
			return []Customer{}, nil
		}

		return []Customer{}, errorFindQuery
	}

	defer cursor.Close(context.TODO())

	// Iterate over the result set and decode the documents into []customer
	var customers []Customer
	for cursor.Next(context.TODO()) {
		var customer Customer
		errorDecode := cursor.Decode(&customer)
		if errorDecode != nil {
			return nil, errorDecode
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func UpdateName(
	updateList []UpdateCandidate,
	newCustomerName string,
) ([]UpdateCandidate, error) {

	if newCustomerName == "" {
		return nil, errors.New("customer name cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			customer:  Customer{Name: newCustomerName},
			objective: UPDATE_NAME,
		},
	)

	return updateList, nil
}

func UpdateEmail(
	updateList []UpdateCandidate,
	newCustomerEmail string,
) ([]UpdateCandidate, error) {

	if newCustomerEmail == "" {
		return nil, errors.New("customer email cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			customer:  Customer{Email: newCustomerEmail},
			objective: UPDATE_EMAIL,
		},
	)

	return updateList, nil
}

func UpdateAddress(
	updateList []UpdateCandidate,
	newCustomerAddress string,
) ([]UpdateCandidate, error) {

	if newCustomerAddress == "" {
		return nil, errors.New("customer address cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			customer:  Customer{Address: newCustomerAddress},
			objective: UPDATE_ADDRESS,
		},
	)

	return updateList, nil
}

func UpdatePhoneNumber(
	updateList []UpdateCandidate,
	newCustomerPhoneNumber string,
) ([]UpdateCandidate, error) {

	if newCustomerPhoneNumber == "" {
		return nil, errors.New("customer Phone Number cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			customer:  Customer{PhoneNumber: newCustomerPhoneNumber},
			objective: UPDATE_PHONENUMBER,
		},
	)

	return updateList, nil
}

func ExecUpdate(
	client *mongo.Client,
	updateList []UpdateCandidate,
	customerUUID string,
) error {

	if customerUUID == "" {
		return errors.New("customer UUID cannot be blank")
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			customer:  Customer{UpdatedAt: time.Now().Unix()},
			objective: UPDATE_UPDATED_AT,
		},
	)

	var setList bson.D
	var unsetList bson.D

	if len(updateList) > 0 {
		for i := 0; i < len(updateList); i++ {
			updateField := reflect.ValueOf(updateList[i].customer)

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

	_, customerExist, errorFindCustomer := Get(
		client,
		customerUUID,
	)

	if !customerExist {
		if errorFindCustomer != nil {
			return errorFindCustomer
		} else {
			return errors.New("customer not found")
		}
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

	filterUpdate := bson.D{{Key: "uuid", Value: customerUUID}}

	_, err := coll.UpdateOne(context.TODO(), filterUpdate, update)

	if err != nil {
		return errors.New("update customer failed")
	}

	return nil
}

func Delete(
	client *mongo.Client,
	customerUUID string,
) (bool, error) {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: customerUUID}}

	var findResult Customer
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		return false, errors.New("customer not found")
	}

	_, errorDelete := coll.DeleteOne(context.TODO(), filter)
	if errorDelete != nil {
		return false, errorDelete
	}

	return true, nil
}
