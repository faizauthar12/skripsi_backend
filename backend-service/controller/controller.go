package controller

const (
	DEFAULT_NUM_ITEMS int64 = 10
	DEFAULT_PAGES     int64 = 1

	SERVER_MALFUNCTION_CANNOT_CREATE_USER    = "Server malfunction, cannot create user"
	SERVER_MALFUNCTION_CANNOT_CREATE_TOKEN   = "Server malfunction, cannot create token"
	SERVER_MALFUNCTION_CANNOT_UPDATE_USER    = "Server malfunction, cannot update user"
	SERVER_MALFUNCTION_CANNOT_CREATE_PRODUCT = "Server malfunction, cannot create product"
	SERVER_MALFUNCTION_CANNOT_UPDATE_PRODUCT = "Server malfunction, cannot update product"
	SERVER_MALFUNCTION_CANNOT_DELETE_PRODUCT = "Server malfunction, cannot delete product"
	SERVER_MALFUNCTION_CANNOT_GET_PRODUCT    = "Server malfunction, cannot get product"
	SUCCESS_DELETE_PRODUCT                   = "Successfully delete the product"

	UNAUTHORIZED = "Unauthorized"

	SUCCESS_CREATE_USER    = "Successfully create user"
	SUCCESS_LOGIN_USER     = "Successfully logged in"
	SUCCESS_UPDATE_USER    = "Successfully update user"
	SUCCESS_CREATE_PRODUCT = "Successfully create product"
	SUCCESS_UPDATE_PRODUCT = "Successfully update the product"
	SUCCESS_DELETE_SERVICE = "Successfully delete the product"
	SUCCESS_GET_PRODUCT    = "Successfully get product"

	USER_NOT_FOUND = "email or password not found"
)
