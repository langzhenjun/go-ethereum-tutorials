// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examples

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenERC20ABI is the input ABI used to generate the binding from.
const TokenERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// TokenERC20Bin is the compiled bytecode used for deploying new contracts.
const TokenERC20Bin = `0x60606040526002805460ff19166012179055341561001c57600080fd5b604051610a1c380380610a1c833981016040528080519190602001805182019190602001805160025460ff16600a0a85026003819055600160a060020a03331660009081526004602052604081209190915592019190508280516100849291602001906100a1565b5060018180516100989291602001906100a1565b5050505061013c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100e257805160ff191683800117855561010f565b8280016001018555821561010f579182015b8281111561010f5782518255916020019190600101906100f4565b5061011b92915061011f565b5090565b61013991905b8082111561011b5760008155600101610125565b90565b6108d18061014b6000396000f3006060604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461017e57806323b872dd146101a3578063313ce567146101cb57806342966c68146101f457806370a082311461020a57806379cc67901461022957806395d89b411461024b578063a9059cbb1461025e578063cae9ca5114610282578063dd62ed3e146102e7575b600080fd5b34156100c957600080fd5b6100d161030c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561010d5780820151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015357600080fd5b61016a600160a060020a03600435166024356103aa565b604051901515815260200160405180910390f35b341561018957600080fd5b6101916103da565b60405190815260200160405180910390f35b34156101ae57600080fd5b61016a600160a060020a03600435811690602435166044356103e0565b34156101d657600080fd5b6101de610457565b60405160ff909116815260200160405180910390f35b34156101ff57600080fd5b61016a600435610460565b341561021557600080fd5b610191600160a060020a03600435166104eb565b341561023457600080fd5b61016a600160a060020a03600435166024356104fd565b341561025657600080fd5b6100d16105d9565b341561026957600080fd5b610280600160a060020a0360043516602435610644565b005b341561028d57600080fd5b61016a60048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061065395505050505050565b34156102f257600080fd5b610191600160a060020a0360043581169060243516610781565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a25780601f10610377576101008083540402835291602001916103a2565b820191906000526020600020905b81548152906001019060200180831161038557829003601f168201915b505050505081565b600160a060020a033381166000908152600560209081526040808320938616835292905220819055600192915050565b60035481565b600160a060020a0380841660009081526005602090815260408083203390941683529290529081205482111561041557600080fd5b600160a060020a038085166000908152600560209081526040808320339094168352929052208054839003905561044d84848461079e565b5060019392505050565b60025460ff1681565b600160a060020a0333166000908152600460205260408120548290101561048657600080fd5b600160a060020a03331660008181526004602052604090819020805485900390556003805485900390557fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a2506001919050565b60046020526000908152604090205481565b600160a060020a0382166000908152600460205260408120548290101561052357600080fd5b600160a060020a038084166000908152600560209081526040808320339094168352929052205482111561055657600080fd5b600160a060020a038084166000818152600460209081526040808320805488900390556005825280832033909516835293905282902080548590039055600380548590039055907fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a250600192915050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a25780601f10610377576101008083540402835291602001916103a2565b61064f33838361079e565b5050565b60008361066081856103aa565b156107795780600160a060020a0316638f4ffcb1338630876040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156107165780820151838201526020016106fe565b50505050905090810190601f1680156107435780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b151561076457600080fd5b5af1151561077157600080fd5b505050600191505b509392505050565b600560209081526000928352604080842090915290825290205481565b6000600160a060020a03831615156107b557600080fd5b600160a060020a038416600090815260046020526040902054829010156107db57600080fd5b600160a060020a038316600090815260046020526040902054828101101561080257600080fd5b50600160a060020a0380831660008181526004602052604080822080549488168084528284208054888103909155938590528154870190915591909301927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3600160a060020a0380841660009081526004602052604080822054928716825290205401811461089f57fe5b505050505600a165627a7a723058203015e14d3496243c9dd77983148097167b857ccf179ae4657448d39b964523a00029`

// DeployTokenERC20 deploys a new Ethereum contract, binding an instance of TokenERC20 to it.
func DeployTokenERC20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *TokenERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenERC20Bin), backend, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// TokenERC20 is an auto generated Go binding around an Ethereum contract.
type TokenERC20 struct {
	TokenERC20Caller     // Read-only binding to the contract
	TokenERC20Transactor // Write-only binding to the contract
	TokenERC20Filterer   // Log filterer for contract events
}

// TokenERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC20Session struct {
	Contract     *TokenERC20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC20CallerSession struct {
	Contract *TokenERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC20TransactorSession struct {
	Contract     *TokenERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC20Raw struct {
	Contract *TokenERC20 // Generic contract binding to access the raw methods on
}

// TokenERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC20CallerRaw struct {
	Contract *TokenERC20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC20TransactorRaw struct {
	Contract *TokenERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC20 creates a new instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20(address common.Address, backend bind.ContractBackend) (*TokenERC20, error) {
	contract, err := bindTokenERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC20{TokenERC20Caller: TokenERC20Caller{contract: contract}, TokenERC20Transactor: TokenERC20Transactor{contract: contract}, TokenERC20Filterer: TokenERC20Filterer{contract: contract}}, nil
}

// NewTokenERC20Caller creates a new read-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Caller(address common.Address, caller bind.ContractCaller) (*TokenERC20Caller, error) {
	contract, err := bindTokenERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Caller{contract: contract}, nil
}

// NewTokenERC20Transactor creates a new write-only instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC20Transactor, error) {
	contract, err := bindTokenERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Transactor{contract: contract}, nil
}

// NewTokenERC20Filterer creates a new log filterer instance of TokenERC20, bound to a specific deployed contract.
func NewTokenERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC20Filterer, error) {
	contract, err := bindTokenERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC20Filterer{contract: contract}, nil
}

// bindTokenERC20 binds a generic wrapper to an already deployed contract.
func bindTokenERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.TokenERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.TokenERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC20 *TokenERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC20 *TokenERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC20 *TokenERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TokenERC20 *TokenERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TokenERC20 *TokenERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.Allowance(&_TokenERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TokenERC20 *TokenERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TokenERC20 *TokenERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TokenERC20.Contract.BalanceOf(&_TokenERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenERC20 *TokenERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenERC20 *TokenERC20Session) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenERC20 *TokenERC20CallerSession) Decimals() (uint8, error) {
	return _TokenERC20.Contract.Decimals(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenERC20 *TokenERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenERC20 *TokenERC20Session) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Name() (string, error) {
	return _TokenERC20.Contract.Name(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenERC20 *TokenERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenERC20 *TokenERC20Session) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenERC20 *TokenERC20CallerSession) Symbol() (string, error) {
	return _TokenERC20.Contract.Symbol(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenERC20 *TokenERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenERC20 *TokenERC20Session) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenERC20 *TokenERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenERC20.Contract.TotalSupply(&_TokenERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Approve(&_TokenERC20.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TokenERC20 *TokenERC20Transactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TokenERC20 *TokenERC20Session) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TokenERC20.Contract.ApproveAndCall(&_TokenERC20.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TokenERC20 *TokenERC20TransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TokenERC20.Contract.ApproveAndCall(&_TokenERC20.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Transactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Session) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Burn(&_TokenERC20.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20TransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Burn(&_TokenERC20.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Transactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Session) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.BurnFrom(&_TokenERC20.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20TransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.BurnFrom(&_TokenERC20.TransactOpts, _from, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TokenERC20 *TokenERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TokenERC20 *TokenERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TokenERC20 *TokenERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.Transfer(&_TokenERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TokenERC20 *TokenERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenERC20.Contract.TransferFrom(&_TokenERC20.TransactOpts, _from, _to, _value)
}

// TokenERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the TokenERC20 contract.
type TokenERC20BurnIterator struct {
	Event *TokenERC20Burn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Burn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC20Burn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Burn represents a Burn event raised by the TokenERC20 contract.
type TokenERC20Burn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_TokenERC20 *TokenERC20Filterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*TokenERC20BurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20BurnIterator{contract: _TokenERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(from indexed address, value uint256)
func (_TokenERC20 *TokenERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *TokenERC20Burn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Burn)
				if err := _TokenERC20.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenERC20 contract.
type TokenERC20TransferIterator struct {
	Event *TokenERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC20Transfer represents a Transfer event raised by the TokenERC20 contract.
type TokenERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TokenERC20 *TokenERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC20TransferIterator{contract: _TokenERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TokenERC20 *TokenERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC20Transfer)
				if err := _TokenERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenRecipientABI is the input ABI used to generate the binding from.
const TokenRecipientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenRecipientBin is the compiled bytecode used for deploying new contracts.
const TokenRecipientBin = `0x`

// DeployTokenRecipient deploys a new Ethereum contract, binding an instance of TokenRecipient to it.
func DeployTokenRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRecipient, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// TokenRecipient is an auto generated Go binding around an Ethereum contract.
type TokenRecipient struct {
	TokenRecipientCaller     // Read-only binding to the contract
	TokenRecipientTransactor // Write-only binding to the contract
	TokenRecipientFilterer   // Log filterer for contract events
}

// TokenRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRecipientSession struct {
	Contract     *TokenRecipient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRecipientCallerSession struct {
	Contract *TokenRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRecipientTransactorSession struct {
	Contract     *TokenRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRecipientRaw struct {
	Contract *TokenRecipient // Generic contract binding to access the raw methods on
}

// TokenRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRecipientCallerRaw struct {
	Contract *TokenRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRecipientTransactorRaw struct {
	Contract *TokenRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRecipient creates a new instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipient(address common.Address, backend bind.ContractBackend) (*TokenRecipient, error) {
	contract, err := bindTokenRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// NewTokenRecipientCaller creates a new read-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientCaller(address common.Address, caller bind.ContractCaller) (*TokenRecipientCaller, error) {
	contract, err := bindTokenRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientCaller{contract: contract}, nil
}

// NewTokenRecipientTransactor creates a new write-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRecipientTransactor, error) {
	contract, err := bindTokenRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientTransactor{contract: contract}, nil
}

// NewTokenRecipientFilterer creates a new log filterer instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRecipientFilterer, error) {
	contract, err := bindTokenRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientFilterer{contract: contract}, nil
}

// bindTokenRecipient binds a generic wrapper to an already deployed contract.
func bindTokenRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.TokenRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactor) ReceiveApproval(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.contract.Transact(opts, "receiveApproval", _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactorSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}
