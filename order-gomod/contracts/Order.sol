// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract OrderContract {
    struct Order {
        string OrderUUID;
        string[] ProductUUID;
        int64[] ProductQuantity;
        int64[] ProductTotalPrice;
        int64 CartGrandTotal;
        string CustomerUUID;
        string CustomerName;
        string CustomerEmail;
        string CustomerAddress;
        string CustomerPhoneNumber;
        string Status;
    }

    mapping(uint256 => Order) private orders;
    mapping(string => uint256) private orderIndexes; // Mapping order UUID to its index
    uint256 private orderCount;

    // Setter function for creating or updating an order
    function setOrder(
        string memory _orderUUID,
        string[] memory _productUUID,
        int64[] memory _productQuantity,
        int64[] memory _productTotalPrice,
        int64 _cartGrandTotal,
        string memory _customerUUID,
        string memory _customerName,
        string memory _customerEmail,
        string memory _customerAddress,
        string memory _customerPhoneNumber,
        string memory _status
    ) public {
        orders[orderCount] = Order(
            _orderUUID,
            _productUUID,
            _productQuantity,
            _productTotalPrice,
            _cartGrandTotal,
            _customerUUID,
            _customerName,
            _customerEmail,
            _customerAddress,
            _customerPhoneNumber,
            _status
        );

        orderIndexes[_orderUUID] = orderCount; // Update mapping for order UUID to its index
        orderCount++;
    }

    // Getter function to retrieve an order by its UUID
    function getOrder(
        string memory orderUUID
    ) public view returns (Order memory) {
        uint256 index = orderIndexes[orderUUID];
        require(index < orderCount, "Invalid order UUID");
        return orders[index];
    }

    // Getter function to get the total number of orders
    function getOrderCount() public view returns (uint256) {
        return orderCount;
    }
}
