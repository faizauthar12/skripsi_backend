// SPDX-License-Identifier: MIT
pragma solidity >=0.4.16 <0.9.0;

contract Order {
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

    function set(
        string memory orderUUID,
        string[] memory productUUID,
        int64[] memory productQuantity,
        int64[] memory productTotalPrice,
        int64 cartGrandTotal,
        string memory customerUUID,
        string memory customerName,
        string memory customerEmail,
        string memory customerAddress,
        string memory customerPhoneNumber,
        string memory status
    ) public {
        OrderUUID = orderUUID;
        ProductUUID = productUUID;
        ProductQuantity = productQuantity;
        ProductTotalPrice = productTotalPrice;
        CartGrandTotal = cartGrandTotal;
        CustomerUUID = customerUUID;
        CustomerName = customerName;
        CustomerEmail = customerEmail;
        CustomerAddress = customerAddress;
        CustomerPhoneNumber = customerPhoneNumber;
        Status = status;
    }

    function get()
        public
        view
        returns (
            string memory,
            string[] memory,
            int64[] memory,
            int64[] memory,
            int64,
            string memory,
            string memory,
            string memory,
            string memory,
            string memory,
            string memory
        )
    {
        return (
            OrderUUID,
            ProductUUID,
            ProductQuantity,
            ProductTotalPrice,
            CartGrandTotal,
            CustomerUUID,
            CustomerName,
            CustomerEmail,
            CustomerAddress,
            CustomerPhoneNumber,
            Status
        );
    }
}
