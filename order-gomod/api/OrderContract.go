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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"orderInstance\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"UUID\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"cartGrandTotal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"customerUUID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"customerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"customerEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"customerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"customerPhoneNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_UUID\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_productUUIDs\",\"type\":\"string[]\"},{\"internalType\":\"int64[]\",\"name\":\"_productQuantities\",\"type\":\"int64[]\"},{\"internalType\":\"int64[]\",\"name\":\"_productTotalPrices\",\"type\":\"int64[]\"},{\"internalType\":\"int64\",\"name\":\"_cartGrandTotal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"_customerUUID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_customerPhoneNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"}],\"name\":\"storeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610d078061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063b35fa0b114610038578063df138e581461005d575b5f80fd5b610040610072565b604051610054989796959493929190610780565b60405180910390f35b61007061006b3660046109f8565b61045b565b005b5f8054819061008090610b8f565b80601f01602080910402602001604051908101604052809291908181526020018280546100ac90610b8f565b80156100f75780601f106100ce576101008083540402835291602001916100f7565b820191905f5260205f20905b8154815290600101906020018083116100da57829003601f168201915b505050505090806004015f9054906101000a900460070b9080600501805461011e90610b8f565b80601f016020809104026020016040519081016040528092919081815260200182805461014a90610b8f565b80156101955780601f1061016c57610100808354040283529160200191610195565b820191905f5260205f20905b81548152906001019060200180831161017857829003601f168201915b5050505050908060060180546101aa90610b8f565b80601f01602080910402602001604051908101604052809291908181526020018280546101d690610b8f565b80156102215780601f106101f857610100808354040283529160200191610221565b820191905f5260205f20905b81548152906001019060200180831161020457829003601f168201915b50505050509080600701805461023690610b8f565b80601f016020809104026020016040519081016040528092919081815260200182805461026290610b8f565b80156102ad5780601f10610284576101008083540402835291602001916102ad565b820191905f5260205f20905b81548152906001019060200180831161029057829003601f168201915b5050505050908060080180546102c290610b8f565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee90610b8f565b80156103395780601f1061031057610100808354040283529160200191610339565b820191905f5260205f20905b81548152906001019060200180831161031c57829003601f168201915b50505050509080600901805461034e90610b8f565b80601f016020809104026020016040519081016040528092919081815260200182805461037a90610b8f565b80156103c55780601f1061039c576101008083540402835291602001916103c5565b820191905f5260205f20905b8154815290600101906020018083116103a857829003601f168201915b50505050509080600a0180546103da90610b8f565b80601f016020809104026020016040519081016040528092919081815260200182805461040690610b8f565b80156104515780601f1061042857610100808354040283529160200191610451565b820191905f5260205f20905b81548152906001019060200180831161043457829003601f168201915b5050505050905088565b6040518061016001604052808c81526020018b81526020018a81526020018981526020018860070b8152602001878152602001868152602001858152602001848152602001838152602001828152505f80820151815f0190816104be9190610c15565b5060208281015180516104d792600185019201906105c8565b50604082015180516104f391600284019160209091019061061c565b506060820151805161050f91600384019160209091019061061c565b50608082015160048201805467ffffffffffffffff191667ffffffffffffffff90921691909117905560a0820151600582019061054c9082610c15565b5060c082015160068201906105619082610c15565b5060e082015160078201906105769082610c15565b50610100820151600882019061058c9082610c15565b5061012082015160098201906105a29082610c15565b50610140820151600a8201906105b89082610c15565b5050505050505050505050505050565b828054828255905f5260205f2090810192821561060c579160200282015b8281111561060c57825182906105fc9082610c15565b50916020019190600101906105e6565b506106189291506106d3565b5090565b828054828255905f5260205f20906003016004900481019282156106c7579160200282015f5b8382111561069157835183826101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055509260200192600801602081600701049283019260010302610642565b80156106c55782816101000a81549067ffffffffffffffff0219169055600801602081600701049283019260010302610691565b505b506106189291506106ef565b80821115610618575f6106e68282610703565b506001016106d3565b5b80821115610618575f81556001016106f0565b50805461070f90610b8f565b5f825580601f1061071e575050565b601f0160209004905f5260205f209081019061073a91906106ef565b50565b5f81518084525f5b8181101561076157602081850181015186830182015201610745565b505f602082860101526020601f19601f83011685010191505092915050565b5f6101008083526107938184018c61073d565b90508960070b602084015282810360408401526107b0818a61073d565b905082810360608401526107c4818961073d565b905082810360808401526107d8818861073d565b905082810360a08401526107ec818761073d565b905082810360c0840152610800818661073d565b905082810360e0840152610814818561073d565b9b9a5050505050505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561086057610860610823565b604052919050565b5f82601f830112610877575f80fd5b813567ffffffffffffffff81111561089157610891610823565b6108a4601f8201601f1916602001610837565b8181528460208386010111156108b8575f80fd5b816020850160208301375f918101602001919091529392505050565b5f67ffffffffffffffff8211156108ed576108ed610823565b5060051b60200190565b5f82601f830112610906575f80fd5b8135602061091b610916836108d4565b610837565b82815260059290921b84018101918181019086841115610939575f80fd5b8286015b8481101561097857803567ffffffffffffffff81111561095c575f8081fd5b61096a8986838b0101610868565b84525091830191830161093d565b509695505050505050565b8035600781900b8114610994575f80fd5b919050565b5f82601f8301126109a8575f80fd5b813560206109b8610916836108d4565b82815260059290921b840181019181810190868411156109d6575f80fd5b8286015b84811015610978576109eb81610983565b83529183019183016109da565b5f805f805f805f805f805f6101608c8e031215610a13575f80fd5b67ffffffffffffffff808d351115610a29575f80fd5b610a368e8e358f01610868565b9b508060208e01351115610a48575f80fd5b610a588e60208f01358f016108f7565b9a508060408e01351115610a6a575f80fd5b610a7a8e60408f01358f01610999565b99508060608e01351115610a8c575f80fd5b610a9c8e60608f01358f01610999565b9850610aaa60808e01610983565b97508060a08e01351115610abc575f80fd5b610acc8e60a08f01358f01610868565b96508060c08e01351115610ade575f80fd5b610aee8e60c08f01358f01610868565b95508060e08e01351115610b00575f80fd5b610b108e60e08f01358f01610868565b9450806101008e01351115610b23575f80fd5b610b348e6101008f01358f01610868565b9350806101208e01351115610b47575f80fd5b610b588e6101208f01358f01610868565b9250806101408e01351115610b6b575f80fd5b50610b7d8d6101408e01358e01610868565b90509295989b509295989b9093969950565b600181811c90821680610ba357607f821691505b602082108103610bc157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610c10575f81815260208120601f850160051c81016020861015610bed5750805b601f850160051c820191505b81811015610c0c57828155600101610bf9565b5050505b505050565b815167ffffffffffffffff811115610c2f57610c2f610823565b610c4381610c3d8454610b8f565b84610bc7565b602080601f831160018114610c76575f8415610c5f5750858301515b5f19600386901b1c1916600185901b178555610c0c565b5f85815260208120601f198616915b82811015610ca457888601518255948401946001909101908401610c85565b5085821015610cc157878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220c839265a79119c4fab1a91620278f2741c434aa3800de2e23cf3d894d2444b5464736f6c63430008140033",
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

// OrderInstance is a free data retrieval call binding the contract method 0xb35fa0b1.
//
// Solidity: function orderInstance() view returns(string UUID, int64 cartGrandTotal, string customerUUID, string customerName, string customerEmail, string customerAddress, string customerPhoneNumber, string status)
func (_Api *ApiCaller) OrderInstance(opts *bind.CallOpts) (struct {
	UUID                string
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "orderInstance")

	outstruct := new(struct {
		UUID                string
		CartGrandTotal      int64
		CustomerUUID        string
		CustomerName        string
		CustomerEmail       string
		CustomerAddress     string
		CustomerPhoneNumber string
		Status              string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UUID = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CartGrandTotal = *abi.ConvertType(out[1], new(int64)).(*int64)
	outstruct.CustomerUUID = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.CustomerName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.CustomerEmail = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CustomerAddress = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.CustomerPhoneNumber = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// OrderInstance is a free data retrieval call binding the contract method 0xb35fa0b1.
//
// Solidity: function orderInstance() view returns(string UUID, int64 cartGrandTotal, string customerUUID, string customerName, string customerEmail, string customerAddress, string customerPhoneNumber, string status)
func (_Api *ApiSession) OrderInstance() (struct {
	UUID                string
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}, error) {
	return _Api.Contract.OrderInstance(&_Api.CallOpts)
}

// OrderInstance is a free data retrieval call binding the contract method 0xb35fa0b1.
//
// Solidity: function orderInstance() view returns(string UUID, int64 cartGrandTotal, string customerUUID, string customerName, string customerEmail, string customerAddress, string customerPhoneNumber, string status)
func (_Api *ApiCallerSession) OrderInstance() (struct {
	UUID                string
	CartGrandTotal      int64
	CustomerUUID        string
	CustomerName        string
	CustomerEmail       string
	CustomerAddress     string
	CustomerPhoneNumber string
	Status              string
}, error) {
	return _Api.Contract.OrderInstance(&_Api.CallOpts)
}

// StoreOrder is a paid mutator transaction binding the contract method 0xdf138e58.
//
// Solidity: function storeOrder(string _UUID, string[] _productUUIDs, int64[] _productQuantities, int64[] _productTotalPrices, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiTransactor) StoreOrder(opts *bind.TransactOpts, _UUID string, _productUUIDs []string, _productQuantities []int64, _productTotalPrices []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "storeOrder", _UUID, _productUUIDs, _productQuantities, _productTotalPrices, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}

// StoreOrder is a paid mutator transaction binding the contract method 0xdf138e58.
//
// Solidity: function storeOrder(string _UUID, string[] _productUUIDs, int64[] _productQuantities, int64[] _productTotalPrices, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiSession) StoreOrder(_UUID string, _productUUIDs []string, _productQuantities []int64, _productTotalPrices []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.Contract.StoreOrder(&_Api.TransactOpts, _UUID, _productUUIDs, _productQuantities, _productTotalPrices, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}

// StoreOrder is a paid mutator transaction binding the contract method 0xdf138e58.
//
// Solidity: function storeOrder(string _UUID, string[] _productUUIDs, int64[] _productQuantities, int64[] _productTotalPrices, int64 _cartGrandTotal, string _customerUUID, string _customerName, string _customerEmail, string _customerAddress, string _customerPhoneNumber, string _status) returns()
func (_Api *ApiTransactorSession) StoreOrder(_UUID string, _productUUIDs []string, _productQuantities []int64, _productTotalPrices []int64, _cartGrandTotal int64, _customerUUID string, _customerName string, _customerEmail string, _customerAddress string, _customerPhoneNumber string, _status string) (*types.Transaction, error) {
	return _Api.Contract.StoreOrder(&_Api.TransactOpts, _UUID, _productUUIDs, _productQuantities, _productTotalPrices, _cartGrandTotal, _customerUUID, _customerName, _customerEmail, _customerAddress, _customerPhoneNumber, _status)
}
