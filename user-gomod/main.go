package user

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/xdg-go/pbkdf2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UUID         string
	Name         string
	Email        string
	PasswordHash string
	PasswordSalt string
	CreatedAt    int64
	UpdatedAt    int64
}

type UpdateCandidate struct {
	user      User
	objective int
}

const (
	DATABASE   = "skripsi"
	COLLECTION = "user"

	TOKEN_LIFESPAN int = 1
	API_SECRET         = "H@zama12"

	ERROR_CREATE_USER         = "User: Error on creating data on DB"
	TOKEN_CANNOT_BLANK        = "User: Token cannot be blank"
	BEARER_TOKEN_CANNOT_BLANK = "User: Bearertoken cannot be blank"
	UNKNOWN_SIGNING_METHOD    = "User: Signing method invalid"
	AVATAR_CANNOT_BLANK       = "User: Avatar cannot be blank"
	NAME_CANNOT_BLANK         = "User: Name cannot be blank"
	USERNAME_CANNOT_BLANK     = "User: Username cannot be blank"
	EMAIL_CANNOT_BLANK        = "User: Email cannot be blank"
	PASSWORD_CANNOT_BLANK     = "User: Password cannot be blank"
	USER_EMAIL_CANNOT_BLANK   = "User: User email cannot be blank"
	USER_NOT_FOUND            = "User: User Not found"
	FAILED_UPDATE_USER        = "User: Failed update user, details : "

	UPDATE_NAME          = 1
	UPDATE_EMAIL         = 2
	UPDATE_PASSWORD_HASH = 3
	UPDATE_PASSWORD_SALT = 4
	UPDATE_UPDATED_AT    = 5
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}

func Create(
	client *mongo.Client,
	name string,
	email string,
	password string,
) (string, error) {

	coll := connect(client)

	var passwordHash []byte
	var user User

	uuid := uuid.New().String()

	salt := uniuri.New()
	passwordHash = pbkdf2.Key([]byte(password), []byte(salt), 10000, 64, sha1.New)

	user = User{
		UUID:         uuid,
		Name:         name,
		Email:        email,
		PasswordHash: hex.EncodeToString(passwordHash),
		PasswordSalt: salt,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	_, errorInsertUser := coll.InsertOne(context.TODO(), user)

	if errorInsertUser != nil {
		return uuid, errorInsertUser
	}

	return uuid, nil
}

func GenerateToken(user User, apiSecret string, expiredToken bool) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uuid"] = user.UUID
	claims["name"] = user.Name
	claims["email"] = user.Email

	// for testing expiredToken
	if expiredToken {
		claims["exp"] = time.Now().Add(time.Hour * -2).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(TOKEN_LIFESPAN)).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// if atribut "" then use default ApiSecret, to test GenerateToken with random API SecretKey.
	if apiSecret == "" {
		apiSecret = API_SECRET
	}

	return token.SignedString([]byte(API_SECRET))
}

func ExtractToken(bearerToken string) (User, error) {

	bearerToken = ExtractBearerToken(bearerToken)

	jwtToken, errorVerifyToken := VerifyTokenValid(bearerToken)

	if errorVerifyToken != nil {
		return User{}, errorVerifyToken
	}

	// claims uuid from token
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		uuid := fmt.Sprint(claims["uuid"])
		name := fmt.Sprint(claims["name"])
		email := fmt.Sprint(claims["email"])

		user := User{
			UUID:  uuid,
			Name:  name,
			Email: email,
		}

		return user, nil
	}

	return User{}, errorVerifyToken
}

func VerifyTokenValid(bearerToken string) (*jwt.Token, error) {

	bearerToken = ExtractBearerToken(bearerToken)

	// check if the token valid
	jwtToken, errorJwtParsing := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(UNKNOWN_SIGNING_METHOD)
		}
		return []byte(API_SECRET), nil
	})

	if errorJwtParsing != nil {
		return nil, errorJwtParsing
	}

	return jwtToken, nil
}

func ExtractBearerToken(bearerToken string) string {
	// Extract token from bearerToken
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	}

	return bearerToken
}

func NativeAuthenticate(
	client *mongo.Client,
	email string,
	password string,
) (User, bool) {

	var result User

	query := bson.D{{Key: "email", Value: email}}

	coll := connect(client)

	errorFindAccount := coll.FindOne(
		context.TODO(),
		query,
	).Decode(&result)

	if errorFindAccount != nil {
		return User{}, false
	}

	insertedPasswordHash := hex.EncodeToString(
		pbkdf2.Key(
			[]byte(password),
			[]byte(result.PasswordSalt), 10000, 64, sha1.New),
	)

	if insertedPasswordHash != result.PasswordHash {
		return User{}, false
	}

	return result, true
}

func Get(
	client *mongo.Client,
	userEmail string,
) (User, bool, error) {

	coll := connect(client)

	var filter bson.D
	if userEmail != "" {
		filter = bson.D{{Key: "email", Value: userEmail}}
	}

	var findResult User
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		if findQueryError == mongo.ErrNoDocuments {
			return User{}, false, nil
		}
		return User{}, false, findQueryError
	}

	// remove the sensitive data!
	findResult = User{
		UUID:      findResult.UUID,
		Name:      findResult.Name,
		Email:     findResult.Email,
		CreatedAt: findResult.CreatedAt,
		UpdatedAt: findResult.UpdatedAt,
	}

	return findResult, true, nil
}

func UpdateEmail(
	updateList []UpdateCandidate,
	newEmail string,
) ([]UpdateCandidate, error) {

	if newEmail == "" {
		return nil, errors.New(EMAIL_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			user:      User{Email: newEmail},
			objective: UPDATE_EMAIL,
		},
	)

	return updateList, nil
}

func UpdateName(
	updateList []UpdateCandidate,
	newName string,
) ([]UpdateCandidate, error) {

	if newName == "" {
		return nil, errors.New(NAME_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			user:      User{Name: newName},
			objective: UPDATE_NAME,
		},
	)

	return updateList, nil
}

func UpdatePassword(
	updateList []UpdateCandidate,
	newPassword string,
) ([]UpdateCandidate, error) {

	if newPassword == "" {
		return nil, errors.New(PASSWORD_CANNOT_BLANK)
	}

	salt := uniuri.New()
	passwordHash := pbkdf2.Key([]byte(newPassword), []byte(salt), 10000, 64, sha1.New)

	updateList = append(
		updateList,
		UpdateCandidate{
			user:      User{PasswordHash: hex.EncodeToString(passwordHash)},
			objective: UPDATE_PASSWORD_HASH,
		},
	)

	updateList = append(
		updateList,
		UpdateCandidate{
			user:      User{PasswordSalt: salt},
			objective: UPDATE_PASSWORD_SALT,
		},
	)

	return updateList, nil
}

func ExecUpdate(
	client *mongo.Client,
	updateList []UpdateCandidate,
	userEmail string,
) error {

	coll := connect(client)

	if userEmail == "" {
		return errors.New(USER_EMAIL_CANNOT_BLANK)
	}

	updateList = append(
		updateList,
		UpdateCandidate{
			user:      User{UpdatedAt: time.Now().Unix()},
			objective: UPDATE_UPDATED_AT,
		},
	)

	var setList bson.D

	if len(updateList) > 0 {
		for i := 0; i < len(updateList); i++ {
			updateField := reflect.ValueOf(updateList[i].user)

			fieldName := strings.ToLower(
				updateField.
					Type().
					Field(updateList[i].objective).
					Name,
			)

			value := updateField.Field(updateList[i].objective).Interface()

			setList = append(setList, bson.E{fieldName, value})
		}
	}

	filter := bson.D{{"email", userEmail}}
	update := bson.D{{"$set", setList}}

	_, errorUpdateUser := coll.UpdateOne(
		context.TODO(),
		filter,
		update,
	)

	if errorUpdateUser != nil {
		return errors.New(FAILED_UPDATE_USER + errorUpdateUser.Error())
	}

	return nil
}

func Delete(
	client *mongo.Client,
	userUUID string,
) error {

	coll := connect(client)

	filter := bson.D{{Key: "uuid", Value: userUUID}}

	var findResult User
	findQueryError := coll.FindOne(
		context.TODO(),
		filter).Decode(&findResult)

	if findQueryError != nil {
		return errors.New(USER_NOT_FOUND)
	}

	_, errorDeleteUser := coll.DeleteOne(context.TODO(), filter)
	if errorDeleteUser != nil {
		return errorDeleteUser
	}

	return nil
}
