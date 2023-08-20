package order

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreatePaymentLink(t *testing.T) {

	errorLoadEnv := godotenv.Load()
	if errorLoadEnv != nil {
		panic(errorLoadEnv)
	}

	durianPayAuth := os.Getenv("DURIANPAY_AUTH")
	customerEmail := os.Getenv("CUSTOMER_EMAIL")

	cartItem := []CartItem{
		{
			ProductUUID:     "c39d5871-ca59-4262-81a7-1b8be3972138",
			ProductName:     "Samsung Galaxy S23",
			ProductQuantity: 1,
			ProductPrice:    12000000,
			ProductPic:      "/static/tv_image.jpg",
		},
	}

	order, errorNewOrder := New(
		cartItem,
		12000000,
		"df3bad63-ca55-4309-8824-dec5ab300be3",
		"Pukisna",
		customerEmail,
		"Jl. Doang ga jadian.",
		"085155054855",
	)

	if errorNewOrder != nil {
		t.Log(errorNewOrder)
		t.FailNow()
	}

	durianPayOrder := NewPaymentDurianPay(order)

	AddItemsDurianPay(&durianPayOrder, &order)

	DurianPaymentLink, errorCreatePaymentLink := CreatePaymentDurianPay(durianPayOrder, durianPayAuth)

	if errorCreatePaymentLink != nil {
		t.Log(errorCreatePaymentLink)
		t.FailNow()
	}

	fmt.Println("DurianPaymentLink", DurianPaymentLink)
}
