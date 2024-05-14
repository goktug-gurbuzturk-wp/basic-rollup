// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// RollupDataLayerMetaData contains all meta data concerning the RollupDataLayer contract.
var RollupDataLayerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_sequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRollupTransaction\",\"inputs\":[{\"name\":\"transaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sequencer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transactions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"}]",
}

// RollupDataLayerABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupDataLayerMetaData.ABI instead.
var RollupDataLayerABI = RollupDataLayerMetaData.ABI

// RollupDataLayer is an auto generated Go binding around an Ethereum contract.
type RollupDataLayer struct {
	RollupDataLayerCaller     // Read-only binding to the contract
	RollupDataLayerTransactor // Write-only binding to the contract
	RollupDataLayerFilterer   // Log filterer for contract events
}

// RollupDataLayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupDataLayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupDataLayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupDataLayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupDataLayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupDataLayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupDataLayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupDataLayerSession struct {
	Contract     *RollupDataLayer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupDataLayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupDataLayerCallerSession struct {
	Contract *RollupDataLayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RollupDataLayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupDataLayerTransactorSession struct {
	Contract     *RollupDataLayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RollupDataLayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupDataLayerRaw struct {
	Contract *RollupDataLayer // Generic contract binding to access the raw methods on
}

// RollupDataLayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupDataLayerCallerRaw struct {
	Contract *RollupDataLayerCaller // Generic read-only contract binding to access the raw methods on
}

// RollupDataLayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupDataLayerTransactorRaw struct {
	Contract *RollupDataLayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupDataLayer creates a new instance of RollupDataLayer, bound to a specific deployed contract.
func NewRollupDataLayer(address common.Address, backend bind.ContractBackend) (*RollupDataLayer, error) {
	contract, err := bindRollupDataLayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupDataLayer{RollupDataLayerCaller: RollupDataLayerCaller{contract: contract}, RollupDataLayerTransactor: RollupDataLayerTransactor{contract: contract}, RollupDataLayerFilterer: RollupDataLayerFilterer{contract: contract}}, nil
}

// NewRollupDataLayerCaller creates a new read-only instance of RollupDataLayer, bound to a specific deployed contract.
func NewRollupDataLayerCaller(address common.Address, caller bind.ContractCaller) (*RollupDataLayerCaller, error) {
	contract, err := bindRollupDataLayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupDataLayerCaller{contract: contract}, nil
}

// NewRollupDataLayerTransactor creates a new write-only instance of RollupDataLayer, bound to a specific deployed contract.
func NewRollupDataLayerTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupDataLayerTransactor, error) {
	contract, err := bindRollupDataLayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupDataLayerTransactor{contract: contract}, nil
}

// NewRollupDataLayerFilterer creates a new log filterer instance of RollupDataLayer, bound to a specific deployed contract.
func NewRollupDataLayerFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupDataLayerFilterer, error) {
	contract, err := bindRollupDataLayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupDataLayerFilterer{contract: contract}, nil
}

// bindRollupDataLayer binds a generic wrapper to an already deployed contract.
func bindRollupDataLayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupDataLayerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupDataLayer *RollupDataLayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupDataLayer.Contract.RollupDataLayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupDataLayer *RollupDataLayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.RollupDataLayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupDataLayer *RollupDataLayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.RollupDataLayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupDataLayer *RollupDataLayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupDataLayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupDataLayer *RollupDataLayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupDataLayer *RollupDataLayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.contract.Transact(opts, method, params...)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_RollupDataLayer *RollupDataLayerCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupDataLayer.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_RollupDataLayer *RollupDataLayerSession) Sequencer() (common.Address, error) {
	return _RollupDataLayer.Contract.Sequencer(&_RollupDataLayer.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_RollupDataLayer *RollupDataLayerCallerSession) Sequencer() (common.Address, error) {
	return _RollupDataLayer.Contract.Sequencer(&_RollupDataLayer.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes)
func (_RollupDataLayer *RollupDataLayerCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _RollupDataLayer.contract.Call(opts, &out, "transactions", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes)
func (_RollupDataLayer *RollupDataLayerSession) Transactions(arg0 *big.Int) ([]byte, error) {
	return _RollupDataLayer.Contract.Transactions(&_RollupDataLayer.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bytes)
func (_RollupDataLayer *RollupDataLayerCallerSession) Transactions(arg0 *big.Int) ([]byte, error) {
	return _RollupDataLayer.Contract.Transactions(&_RollupDataLayer.CallOpts, arg0)
}

// AddRollupTransaction is a paid mutator transaction binding the contract method 0xec0e7793.
//
// Solidity: function addRollupTransaction(bytes transaction) returns()
func (_RollupDataLayer *RollupDataLayerTransactor) AddRollupTransaction(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _RollupDataLayer.contract.Transact(opts, "addRollupTransaction", transaction)
}

// AddRollupTransaction is a paid mutator transaction binding the contract method 0xec0e7793.
//
// Solidity: function addRollupTransaction(bytes transaction) returns()
func (_RollupDataLayer *RollupDataLayerSession) AddRollupTransaction(transaction []byte) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.AddRollupTransaction(&_RollupDataLayer.TransactOpts, transaction)
}

// AddRollupTransaction is a paid mutator transaction binding the contract method 0xec0e7793.
//
// Solidity: function addRollupTransaction(bytes transaction) returns()
func (_RollupDataLayer *RollupDataLayerTransactorSession) AddRollupTransaction(transaction []byte) (*types.Transaction, error) {
	return _RollupDataLayer.Contract.AddRollupTransaction(&_RollupDataLayer.TransactOpts, transaction)
}
