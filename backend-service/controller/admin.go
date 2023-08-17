package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	Customer "github.com/faizauthar12/skripsi/customer-gomod"
	Order "github.com/faizauthar12/skripsi/order-gomod"
	Product "github.com/faizauthar12/skripsi/product-gomod"
	"github.com/gin-gonic/gin"
)

func HandlePagination(pages int64, pageLink string, itemCount int64) (int64, string, string) {
	var maxPages int64

	if itemCount > 10 {
		maxPages = int64(math.Ceil(float64(itemCount) / 10))
	} else {
		maxPages = 1
	}

	var previousPageLink string

	nextPageLink := fmt.Sprintf("/admin/%s?pages=%d", pageLink, pages+1)
	if pages > 0 {
		previousPageLink = fmt.Sprintf("/admin/%s?pages=%d", pageLink, pages-1)
	}

	return maxPages, previousPageLink, nextPageLink
}

func (controller *Controller) HomePage(c *gin.Context) {

	productCount, errorGetProductCount := Product.GetCount(
		controller.Client,
	)

	if errorGetProductCount != nil {
		fmt.Println("Admin: Error(): ", errorGetProductCount)
	}

	customerCount, errorGetCustomerCount := Customer.GetCount(
		controller.Client,
	)

	if errorGetCustomerCount != nil {
		fmt.Println("Admin: Error(): ", errorGetCustomerCount)
	}

	orderCount, errorGetOrderCount := Order.GetCount(
		controller.Client,
	)

	if errorGetOrderCount != nil {
		fmt.Println("Admin: Error(): ", errorGetOrderCount)
	}

	c.HTML(http.StatusOK, "homepage.tmpl", gin.H{
		"title":         "Admin Welcome page",
		"ProdukTotal":   productCount,
		"CustomerTotal": customerCount,
		"OrderTotal":    orderCount,
	})
}

func (controller *Controller) ProductPage(c *gin.Context) {
	numItems, errorParsingNumItems := strconv.ParseInt(c.Query("numItems"), 10, 64)
	pages, errorParsingPages := strconv.ParseInt(c.Query("pages"), 10, 64)
	category := c.Query("category")

	if errorParsingNumItems != nil {
		numItems = DEFAULT_NUM_ITEMS
	}

	if errorParsingPages != nil {
		pages = DEFAULT_PAGES
	}

	productCount, errorGetProductCount := Product.GetCount(
		controller.Client,
	)

	if errorGetProductCount != nil {
		fmt.Println("Admin: Product: Error(): ", errorGetProductCount)
	}

	products, errorGetProducts := Product.GetMany(
		controller.Client,
		category,
		numItems,
		pages,
	)

	if errorGetProducts != nil {
		fmt.Println("Admin: Product: Error(): ", errorGetProducts)
	}

	maxPages, previousPageLink, nextPageLink := HandlePagination(pages, "product", productCount)

	c.HTML(http.StatusOK, "productpage.tmpl", gin.H{
		"title":        "Product Page",
		"Products":     products,
		"pages":        pages,
		"maxPages":     maxPages,
		"nextPage":     nextPageLink,
		"previousPage": previousPageLink,
	})
}

func (controller *Controller) OrderPage(c *gin.Context) {
	numItems, errorParsingNumItems := strconv.ParseInt(c.Query("numItems"), 10, 64)
	pages, errorParsingPages := strconv.ParseInt(c.Query("pages"), 10, 64)

	if errorParsingNumItems != nil {
		numItems = DEFAULT_NUM_ITEMS
	}

	if errorParsingPages != nil {
		pages = DEFAULT_PAGES
	}

	orderCount, errorGetOrderCount := Order.GetCount(
		controller.Client,
	)

	if errorGetOrderCount != nil {
		fmt.Println("Admin: Order: Error(): ", errorGetOrderCount)
	}

	orders, errorGetOrders := Order.GetMany(
		controller.Client,
		numItems,
		pages,
	)

	if errorGetOrders != nil {
		fmt.Println("Admin: Order: Error(): ", errorGetOrders)
	}

	maxPages, previousPageLink, nextPageLink := HandlePagination(pages, "order", orderCount)

	c.HTML(http.StatusOK, "orderpage.tmpl", gin.H{
		"title":        "Order Page",
		"Orders":       orders,
		"pages":        pages,
		"maxPages":     maxPages,
		"nextPage":     nextPageLink,
		"previousPage": previousPageLink,
	})
}

func (controller *Controller) CustomerPage(c *gin.Context) {
	numItems, errorParsingNumItems := strconv.ParseInt(c.Query("numItems"), 10, 64)
	pages, errorParsingPages := strconv.ParseInt(c.Query("pages"), 10, 64)

	if errorParsingNumItems != nil {
		numItems = DEFAULT_NUM_ITEMS
	}

	if errorParsingPages != nil {
		pages = DEFAULT_PAGES
	}

	customerCount, errorGetCustomerCount := Customer.GetCount(
		controller.Client,
	)

	if errorGetCustomerCount != nil {
		fmt.Println("Admin: Customer: Error(): ", errorGetCustomerCount)
	}

	customers, errorGetCustomers := Customer.GetMany(
		controller.Client,
		numItems,
		pages,
	)

	if errorGetCustomers != nil {
		fmt.Println("Admin: Customers: Error(): ", errorGetCustomers)
	}

	maxPages, previousPageLink, nextPageLink := HandlePagination(pages, "customer", customerCount)

	c.HTML(http.StatusOK, "customerpage.tmpl", gin.H{
		"title":        "Customer Page",
		"Customers":    customers,
		"pages":        pages,
		"maxPages":     maxPages,
		"nextPage":     nextPageLink,
		"previousPage": previousPageLink,
	})
}
