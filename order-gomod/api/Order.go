// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OrderContractOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderContractOrder struct {
	OrderID             string
	ProductUUID         []string
	ProductQuantity     []int64
	ProductTotalPrice   []int64
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"orderID\",\"type\":\"string\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"OrderID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ProductUUID\",\"type\":\"string[]\"},{\"internalType\":\"int64[]\",\"name\":\"ProductQuantity\",\"type\":\"int64[]\"},{\"internalType\":\"int64[]\",\"name\":\"ProductTotalPrice\",\"type\":\"int64[]\"},{\"internalType\":\"int64\",\"name\":\"CartGrandTotal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"CustomerUUID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"CustomerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"CustomerEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"CustomerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"CustomerPhoneNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Status\",\"type\":\"string\"}],\"internalType\":\"structOrderContract.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orderID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_productUUID\",\"type\":\"string[]\"},{\"internalType\":\"int64[]\",\"name\":\"_productQuantity\",\"type\":\"int64[]\"},{\"internalType\":\"int64[]\",\"name\":\"_productTotalPrice\",\"type\":\"int64[]\"},{\"internalType\":\"int64\",\"name\":\"_cartGrandTotal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"_customerUUID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerPhoneNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"}],\"name\":\"setOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506111f08061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80636cdaeec014610043578063712ca0f8146100585780638d0a5fbb14610081575b5f80fd5b610056610051366004610c40565b610092565b005b61006b610066366004610dd7565b610248565b6040516100789190610ef0565b60405180910390f35b600254604051908152602001610078565b6040518061016001604052808c81526020018b81526020018a81526020018981526020018860070b8152602001878152602001868152602001858152602001848152602001838152602001828152505f8060025481526020019081526020015f205f820151815f01908161010691906110bf565b50602082810151805161011f92600185019201906108f6565b506040820151805161013b91600284019160209091019061094a565b506060820151805161015791600384019160209091019061094a565b50608082015160048201805467ffffffffffffffff191667ffffffffffffffff90921691909117905560a0820151600582019061019490826110bf565b5060c082015160068201906101a990826110bf565b5060e082015160078201906101be90826110bf565b5061010082015160088201906101d490826110bf565b5061012082015160098201906101ea90826110bf565b50610140820151600a82019061020090826110bf565b5090505060025460018c604051610217919061117b565b9081526040519081900360200190205560028054905f61023683611196565b91905055505050505050505050505050565b6102a7604051806101600160405280606081526020016060815260200160608152602001606081526020015f60070b81526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b5f6001836040516102b8919061117b565b908152602001604051809103902054905060025481106103115760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081bdc99195c88125160821b604482015260640160405180910390fd5b5f81815260208190526040908190208151610160810190925280548290829061033990611039565b80601f016020809104026020016040519081016040528092919081815260200182805461036590611039565b80156103b05780601f10610387576101008083540402835291602001916103b0565b820191905f5260205f20905b81548152906001019060200180831161039357829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020015f905b82821015610485578382905f5260205f200180546103fa90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461042690611039565b80156104715780601f1061044857610100808354040283529160200191610471565b820191905f5260205f20905b81548152906001019060200180831161045457829003601f168201915b5050505050815260200190600101906103dd565b505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156104fc57602002820191905f5260205f20905f905b825461010083900a900460070b81526020600f83018190049384019360010360089093019290920291018084116104c85790505b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561057457602002820191905f5260205f20905f905b825461010083900a900460070b81526020600f83018190049384019360010360089093019290920291018084116105405790505b5050509183525050600482015460070b602082015260058201805460409092019161059e90611039565b80601f01602080910402602001604051908101604052809291908181526020018280546105ca90611039565b80156106155780601f106105ec57610100808354040283529160200191610615565b820191905f5260205f20905b8154815290600101906020018083116105f857829003601f168201915b5050505050815260200160068201805461062e90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461065a90611039565b80156106a55780601f1061067c576101008083540402835291602001916106a5565b820191905f5260205f20905b81548152906001019060200180831161068857829003601f168201915b505050505081526020016007820180546106be90611039565b80601f01602080910402602001604051908101604052809291908181526020018280546106ea90611039565b80156107355780601f1061070c57610100808354040283529160200191610735565b820191905f5260205f20905b81548152906001019060200180831161071857829003601f168201915b5050505050815260200160088201805461074e90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461077a90611039565b80156107c55780601f1061079c576101008083540402835291602001916107c5565b820191905f5260205f20905b8154815290600101906020018083116107a857829003601f168201915b505050505081526020016009820180546107de90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461080a90611039565b80156108555780601f1061082c57610100808354040283529160200191610855565b820191905f5260205f20905b81548152906001019060200180831161083857829003601f168201915b50505050508152602001600a8201805461086e90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461089a90611039565b80156108e55780601f106108bc576101008083540402835291602001916108e5565b820191905f5260205f20905b8154815290600101906020018083116108c857829003601f168201915b505050505081525050915050919050565b828054828255905f5260205f2090810192821561093a579160200282015b8281111561093a578251829061092a90826110bf565b5091602001919060010190610914565b50610946929150610a01565b5090565b828054828255905f5260205f20906003016004900481019282156109f5579160200282015f5b838211156109bf57835183826101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055509260200192600801602081600701049283019260010302610970565b80156109f35782816101000a81549067ffffffffffffffff02191690556008016020816007010492830192600103026109bf565b505b50610946929150610a1d565b80821115610946575f610a148282610a31565b50600101610a01565b5b80821115610946575f8155600101610a1e565b508054610a3d90611039565b5f825580601f10610a4c575050565b601f0160209004905f5260205f2090810190610a689190610a1d565b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610aa857610aa8610a6b565b604052919050565b5f82601f830112610abf575f80fd5b813567ffffffffffffffff811115610ad957610ad9610a6b565b610aec601f8201601f1916602001610a7f565b818152846020838601011115610b00575f80fd5b816020850160208301375f918101602001919091529392505050565b5f67ffffffffffffffff821115610b3557610b35610a6b565b5060051b60200190565b5f82601f830112610b4e575f80fd5b81356020610b63610b5e83610b1c565b610a7f565b82815260059290921b84018101918181019086841115610b81575f80fd5b8286015b84811015610bc057803567ffffffffffffffff811115610ba4575f8081fd5b610bb28986838b0101610ab0565b845250918301918301610b85565b509695505050505050565b8035600781900b8114610bdc575f80fd5b919050565b5f82601f830112610bf0575f80fd5b81356020610c00610b5e83610b1c565b82815260059290921b84018101918181019086841115610c1e575f80fd5b8286015b84811015610bc057610c3381610bcb565b8352918301918301610c22565b5f805f805f805f805f805f6101608c8e031215610c5b575f80fd5b67ffffffffffffffff808d351115610c71575f80fd5b610c7e8e8e358f01610ab0565b9b508060208e01351115610c90575f80fd5b610ca08e60208f01358f01610b3f565b9a508060408e01351115610cb2575f80fd5b610cc28e60408f01358f01610be1565b99508060608e01351115610cd4575f80fd5b610ce48e60608f01358f01610be1565b9850610cf260808e01610bcb565b97508060a08e01351115610d04575f80fd5b610d148e60a08f01358f01610ab0565b96508060c08e01351115610d26575f80fd5b610d368e60c08f01358f01610ab0565b95508060e08e01351115610d48575f80fd5b610d588e60e08f01358f01610ab0565b9450806101008e01351115610d6b575f80fd5b610d7c8e6101008f01358f01610ab0565b9350806101208e01351115610d8f575f80fd5b610da08e6101208f01358f01610ab0565b9250806101408e01351115610db3575f80fd5b50610dc58d6101408e01358e01610ab0565b90509295989b509295989b9093969950565b5f60208284031215610de7575f80fd5b813567ffffffffffffffff811115610dfd575f80fd5b610e0984828501610ab0565b949350505050565b5f5b83811015610e2b578181015183820152602001610e13565b50505f910152565b5f8151808452610e4a816020860160208601610e11565b601f01601f19169290920160200192915050565b5f82825180855260208086019550808260051b8401018186015f5b84811015610ea757601f19868403018952610e95838351610e33565b98840198925090830190600101610e79565b5090979650505050505050565b5f8151808452602080850194508084015f5b83811015610ee557815160070b87529582019590820190600101610ec6565b509495945050505050565b602081525f8251610160806020850152610f0e610180850183610e33565b91506020850151601f1980868503016040870152610f2c8483610e5e565b93506040870151915080868503016060870152610f498483610eb4565b93506060870151915080868503016080870152610f668483610eb4565b935060808701519150610f7e60a087018360070b9052565b60a08701519150808685030160c0870152610f998483610e33565b935060c08701519150808685030160e0870152610fb68483610e33565b935060e08701519150610100818786030181880152610fd58584610e33565b945080880151925050610120818786030181880152610ff48584610e33565b9450808801519250506101408187860301818801526110138584610e33565b90880151878203909201848801529350905061102f8382610e33565b9695505050505050565b600181811c9082168061104d57607f821691505b60208210810361106b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156110ba575f81815260208120601f850160051c810160208610156110975750805b601f850160051c820191505b818110156110b6578281556001016110a3565b5050505b505050565b815167ffffffffffffffff8111156110d9576110d9610a6b565b6110ed816110e78454611039565b84611071565b602080601f831160018114611120575f84156111095750858301515b5f19600386901b1c1916600185901b1785556110b6565b5f85815260208120601f198616915b8281101561114e5788860151825594840194600190910190840161112f565b508582101561116b57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f825161118c818460208701610e11565b9190910192915050565b5f600182016111b357634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220c122786532b993591a97935c64a053c342ce1900f223c3e67452affa5bb18cb164736f6c63430008150033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetOrder is a free data retrieval call binding the contract method 0x712ca0f8.
//
// Solidity: function getOrder(string orderID) view returns((string,string[],int64[],int64[],int64,string,string,string,string,string,string))
func (_Api *ApiCaller) GetOrder(opts *bind.CallOpts, orderID string) (OrderContractOrder, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOrder", orderID)

	if err != nil {
		return *new(OrderContractOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderContractOrder)).(*OrderContractOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x712ca0f8.
//
// Solidity: function getOrder(string orderID) view returns((string,string[],int64[],int64[],int64,string,string,string,string,string,string))
func (_Api *ApiSession) GetOrder(orderID string) (OrderContractOrder, error) {
	return _Api.Contract.GetOrder(&_Api.CallOpts, orderID)
}

// GetOrder is a free data retrieval call binding the contract method 0x712ca0f8.
//
// Solidity: function getOrder(string orderID) view returns((string,string[],int64[],int64[],int64,string,string,string,string,string,string))
func (_Api *ApiCallerSession) GetOrder(orderID string) (OrderContractOrder, error) {
	return _Api.Contract.GetOrder(&_Api.CallOpts, orderID)
}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Api *ApiCaller) GetOrderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOrderCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Api *ApiSession) GetOrderCount() (*big.Int, error) {
	return _Api.Contract.GetOrderCount(&_Api.CallOpts)
}

// GetOrderCount is a free data retrieval call binding the contract method 0x8d0a5fbb.
//
// Solidity: function getOrderCount() view returns(uint256)
func (_Api *ApiCallerSession) GetOrderCount() (*big.Int, error) {
	return _Api.Contract.GetOrderCount(&_Api.CallOpts)
}

// SetOrder is a paid mutator transaction binding the contract method 0x6cdaeec0.
//
// Solidity: function setOrder(string _orderID, string[] _productUUID, int64[] _productQuantity, int64[] _productTotalPrice, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiTransactor) SetOrder(opts *bind.TransactOpts, _orderID string, _productUUID []string, _productQuantity []int64, _productTotalPrice []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setOrder", _orderID, _productUUID, _productQuantity, _productTotalPrice, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}

// SetOrder is a paid mutator transaction binding the contract method 0x6cdaeec0.
//
// Solidity: function setOrder(string _orderID, string[] _productUUID, int64[] _productQuantity, int64[] _productTotalPrice, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiSession) SetOrder(_orderID string, _productUUID []string, _productQuantity []int64, _productTotalPrice []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.Contract.SetOrder(&_Api.TransactOpts, _orderID, _productUUID, _productQuantity, _productTotalPrice, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}

// SetOrder is a paid mutator transaction binding the contract method 0x6cdaeec0.
//
// Solidity: function setOrder(string _orderID, string[] _productUUID, int64[] _productQuantity, int64[] _productTotalPrice, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiTransactorSession) SetOrder(_orderID string, _productUUID []string, _productQuantity []int64, _productTotalPrice []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.Contract.SetOrder(&_Api.TransactOpts, _orderID, _productUUID, _productQuantity, _productTotalPrice, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}
