package cart

import "go.mongodb.org/mongo-driver/mongo"

type Cart struct {
	UUID           string
	UserUUID       string
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

	// UpdateCandidate for Cart

	UPDATE_CART_GRAND_TOTAL = 2
	UPDATE_UPDATED_AT       = 4
)

func connect(client *mongo.Client) *mongo.Collection {
	return client.Database(DATABASE).Collection(COLLECTION)
}
