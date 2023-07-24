// OrderContract.sol
pragma solidity ^0.8.0;

contract OrderContract {
    struct Order {
        string UUID;
        string[] productUUIDs;
        int64[] productQuantities;
        int64[] productTotalPrices;
        int64 cartGrandTotal;
        string customerUUID;
        string customerName;
        string customerEmail;
        string customerAddress;
        string customerPhoneNumber;
        string status;
    }

    Order public orderInstance;

    function storeOrder(
        string memory _UUID,
        string[] memory _productUUIDs,
        int64[] memory _productQuantities,
        int64[] memory _productTotalPrices,
        int64 _cartGrandTotal,
        string memory _customerUUID,
        string memory _customerName,
        string memory _customerEmail,
        string memory _customerAddress,
        string memory _customerPhoneNumber,
        string memory _status
    ) public {
        orderInstance = Order(
            _UUID,
            _productUUIDs,
            _productQuantities,
            _productTotalPrices,
            _cartGrandTotal,
            _customerUUID,
            _customerName,
            _customerEmail,
            _customerAddress,
            _customerPhoneNumber,
            _status
        );
    }
}
