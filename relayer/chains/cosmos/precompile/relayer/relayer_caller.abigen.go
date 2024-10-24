// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayer

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

// RelayerCallerMetaData contains all meta data concerning the RelayerCaller contract.
var RelayerCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"}],\"name\":\"batchCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105bd8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063f456fa401461002d575b5f5ffd5b61004760048036038101906100429190610413565b610049565b005b5f5f90505b8151811015610185575f8383838151811061006c5761006b61046d565b5b6020026020010151604051602401610085929190610509565b6040516020818303038152906040527f70f040a4000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f606573ffffffffffffffffffffffffffffffffffffffff168260405161012b9190610571565b5f604051808303815f865af19150503d805f8114610164576040519150601f19603f3d011682016040523d82523d5f602084013e610169565b606091505b5050905080610176575f5ffd5b5050808060010191505061004e565b505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101c48261019b565b9050919050565b6101d4816101ba565b81146101de575f5ffd5b50565b5f813590506101ef816101cb565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61023f826101f9565b810181811067ffffffffffffffff8211171561025e5761025d610209565b5b80604052505050565b5f61027061018a565b905061027c8282610236565b919050565b5f67ffffffffffffffff82111561029b5761029a610209565b5b602082029050602081019050919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff8211156102ce576102cd610209565b5b6102d7826101f9565b9050602081019050919050565b828183375f83830152505050565b5f6103046102ff846102b4565b610267565b9050828152602081018484840111156103205761031f6102b0565b5b61032b8482856102e4565b509392505050565b5f82601f830112610347576103466101f5565b5b81356103578482602086016102f2565b91505092915050565b5f61037261036d84610281565b610267565b90508083825260208201905060208402830185811115610395576103946102ac565b5b835b818110156103dc57803567ffffffffffffffff8111156103ba576103b96101f5565b5b8086016103c78982610333565b85526020850194505050602081019050610397565b5050509392505050565b5f82601f8301126103fa576103f96101f5565b5b813561040a848260208601610360565b91505092915050565b5f5f6040838503121561042957610428610193565b5b5f610436858286016101e1565b925050602083013567ffffffffffffffff81111561045757610456610197565b5b610463858286016103e6565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6104a3816101ba565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6104db826104a9565b6104e581856104b3565b93506104f58185602086016104c3565b6104fe816101f9565b840191505092915050565b5f60408201905061051c5f83018561049a565b818103602083015261052e81846104d1565b90509392505050565b5f81905092915050565b5f61054b826104a9565b6105558185610537565b93506105658185602086016104c3565b80840191505092915050565b5f61057c8284610541565b91508190509291505056fea2646970667358221220bd33dee1e8ffb3d8c1d9e2d6f6a52bb57dd77375ca7e02e630b1755a6e6780d064736f6c634300081c0033",
}

// RelayerCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayerCallerMetaData.ABI instead.
var RelayerCallerABI = RelayerCallerMetaData.ABI

// RelayerCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RelayerCallerMetaData.Bin instead.
var RelayerCallerBin = RelayerCallerMetaData.Bin

// DeployRelayerCaller deploys a new Ethereum contract, binding an instance of RelayerCaller to it.
func DeployRelayerCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelayerCaller, error) {
	parsed, err := RelayerCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RelayerCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayerCaller{RelayerCallerCaller: RelayerCallerCaller{contract: contract}, RelayerCallerTransactor: RelayerCallerTransactor{contract: contract}, RelayerCallerFilterer: RelayerCallerFilterer{contract: contract}}, nil
}

// RelayerCaller is an auto generated Go binding around an Ethereum contract.
type RelayerCaller struct {
	RelayerCallerCaller     // Read-only binding to the contract
	RelayerCallerTransactor // Write-only binding to the contract
	RelayerCallerFilterer   // Log filterer for contract events
}

// RelayerCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerCallerSession struct {
	Contract     *RelayerCaller    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerCallerCallerSession struct {
	Contract *RelayerCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RelayerCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerCallerTransactorSession struct {
	Contract     *RelayerCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RelayerCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerCallerRaw struct {
	Contract *RelayerCaller // Generic contract binding to access the raw methods on
}

// RelayerCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerCallerCallerRaw struct {
	Contract *RelayerCallerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerCallerTransactorRaw struct {
	Contract *RelayerCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerCaller creates a new instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCaller(address common.Address, backend bind.ContractBackend) (*RelayerCaller, error) {
	contract, err := bindRelayerCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayerCaller{RelayerCallerCaller: RelayerCallerCaller{contract: contract}, RelayerCallerTransactor: RelayerCallerTransactor{contract: contract}, RelayerCallerFilterer: RelayerCallerFilterer{contract: contract}}, nil
}

// NewRelayerCallerCaller creates a new read-only instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerCaller(address common.Address, caller bind.ContractCaller) (*RelayerCallerCaller, error) {
	contract, err := bindRelayerCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerCaller{contract: contract}, nil
}

// NewRelayerCallerTransactor creates a new write-only instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerCallerTransactor, error) {
	contract, err := bindRelayerCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerTransactor{contract: contract}, nil
}

// NewRelayerCallerFilterer creates a new log filterer instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerCallerFilterer, error) {
	contract, err := bindRelayerCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerFilterer{contract: contract}, nil
}

// bindRelayerCaller binds a generic wrapper to an already deployed contract.
func bindRelayerCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelayerCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerCaller *RelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerCaller.Contract.RelayerCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerCaller *RelayerCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerCaller.Contract.RelayerCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerCaller *RelayerCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerCaller.Contract.RelayerCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerCaller *RelayerCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerCaller *RelayerCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerCaller *RelayerCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerCaller.Contract.contract.Transact(opts, method, params...)
}

// BatchCall is a paid mutator transaction binding the contract method 0xf456fa40.
//
// Solidity: function batchCall(address payee, bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerTransactor) BatchCall(opts *bind.TransactOpts, payee common.Address, payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.contract.Transact(opts, "batchCall", payee, payloads)
}

// BatchCall is a paid mutator transaction binding the contract method 0xf456fa40.
//
// Solidity: function batchCall(address payee, bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerSession) BatchCall(payee common.Address, payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, payee, payloads)
}

// BatchCall is a paid mutator transaction binding the contract method 0xf456fa40.
//
// Solidity: function batchCall(address payee, bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerTransactorSession) BatchCall(payee common.Address, payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, payee, payloads)
}
