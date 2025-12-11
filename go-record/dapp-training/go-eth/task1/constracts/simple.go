// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package constracts

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

// ConstractsMetaData contains all meta data concerning the Constracts contract.
var ConstractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060015f81905550610189806100235f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80634f2be91f14610043578063561892361461004d57806375794a3c1461006b575b5f5ffd5b61004b610089565b005b6100556100a1565b60405161006291906100c6565b60405180910390f35b6100736100a9565b60405161008091906100c6565b60405180910390f35b5f5f81548092919061009a9061010c565b9190505550565b5f5f54905090565b5f5481565b5f819050919050565b6100c0816100ae565b82525050565b5f6020820190506100d95f8301846100b7565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610116826100ae565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610148576101476100df565b5b60018201905091905056fea2646970667358221220fcd493d5a8215b537f606f5b2e5ed90f306cffa046d8c4e26c16bfc45b0af9dc64736f6c634300081f0033",
}

// ConstractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstractsMetaData.ABI instead.
var ConstractsABI = ConstractsMetaData.ABI

// ConstractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstractsMetaData.Bin instead.
var ConstractsBin = ConstractsMetaData.Bin

// DeployConstracts deploys a new Ethereum contract, binding an instance of Constracts to it.
func DeployConstracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constracts, error) {
	parsed, err := ConstractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constracts{ConstractsCaller: ConstractsCaller{contract: contract}, ConstractsTransactor: ConstractsTransactor{contract: contract}, ConstractsFilterer: ConstractsFilterer{contract: contract}}, nil
}

// Constracts is an auto generated Go binding around an Ethereum contract.
type Constracts struct {
	ConstractsCaller     // Read-only binding to the contract
	ConstractsTransactor // Write-only binding to the contract
	ConstractsFilterer   // Log filterer for contract events
}

// ConstractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstractsSession struct {
	Contract     *Constracts       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstractsCallerSession struct {
	Contract *ConstractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ConstractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstractsTransactorSession struct {
	Contract     *ConstractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ConstractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstractsRaw struct {
	Contract *Constracts // Generic contract binding to access the raw methods on
}

// ConstractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstractsCallerRaw struct {
	Contract *ConstractsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstractsTransactorRaw struct {
	Contract *ConstractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstracts creates a new instance of Constracts, bound to a specific deployed contract.
func NewConstracts(address common.Address, backend bind.ContractBackend) (*Constracts, error) {
	contract, err := bindConstracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constracts{ConstractsCaller: ConstractsCaller{contract: contract}, ConstractsTransactor: ConstractsTransactor{contract: contract}, ConstractsFilterer: ConstractsFilterer{contract: contract}}, nil
}

// NewConstractsCaller creates a new read-only instance of Constracts, bound to a specific deployed contract.
func NewConstractsCaller(address common.Address, caller bind.ContractCaller) (*ConstractsCaller, error) {
	contract, err := bindConstracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstractsCaller{contract: contract}, nil
}

// NewConstractsTransactor creates a new write-only instance of Constracts, bound to a specific deployed contract.
func NewConstractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstractsTransactor, error) {
	contract, err := bindConstracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstractsTransactor{contract: contract}, nil
}

// NewConstractsFilterer creates a new log filterer instance of Constracts, bound to a specific deployed contract.
func NewConstractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstractsFilterer, error) {
	contract, err := bindConstracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstractsFilterer{contract: contract}, nil
}

// bindConstracts binds a generic wrapper to an already deployed contract.
func bindConstracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConstractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constracts *ConstractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constracts.Contract.ConstractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constracts *ConstractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constracts.Contract.ConstractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constracts *ConstractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constracts.Contract.ConstractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constracts *ConstractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constracts *ConstractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constracts *ConstractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constracts.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Constracts *ConstractsCaller) GetCurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constracts.contract.Call(opts, &out, "getCurrentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Constracts *ConstractsSession) GetCurrentTokenId() (*big.Int, error) {
	return _Constracts.Contract.GetCurrentTokenId(&_Constracts.CallOpts)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_Constracts *ConstractsCallerSession) GetCurrentTokenId() (*big.Int, error) {
	return _Constracts.Contract.GetCurrentTokenId(&_Constracts.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Constracts *ConstractsCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constracts.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Constracts *ConstractsSession) NextTokenId() (*big.Int, error) {
	return _Constracts.Contract.NextTokenId(&_Constracts.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Constracts *ConstractsCallerSession) NextTokenId() (*big.Int, error) {
	return _Constracts.Contract.NextTokenId(&_Constracts.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Constracts *ConstractsTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constracts.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Constracts *ConstractsSession) Add() (*types.Transaction, error) {
	return _Constracts.Contract.Add(&_Constracts.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Constracts *ConstractsTransactorSession) Add() (*types.Transaction, error) {
	return _Constracts.Contract.Add(&_Constracts.TransactOpts)
}
