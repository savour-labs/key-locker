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
)

// KeyLockerMetaData contains all meta data concerning the KeyLocker contract.
var KeyLockerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_uuid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_keys\",\"type\":\"bytes[]\"}],\"name\":\"keyLockerAppend\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_uuid\",\"type\":\"bytes32\"}],\"name\":\"getSocialKey\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_uuid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_keys\",\"type\":\"bytes[]\"}],\"name\":\"setSocialKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socialKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610ad6806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806388369c9a1161005b57806388369c9a146100a75780638da5cb5b146100d0578063e168952b146100eb578063f2fde38b1461010b57600080fd5b8063715018a61461008257806374aa77891461008c5780638129fc1c1461009f575b600080fd5b61008a61011e565b005b61008a61009a36600461079f565b610132565b61008a6101dd565b6100ba6100b53660046108c3565b6102ee565b6040516100c79190610981565b60405180910390f35b6033546040516001600160a01b0390911681526020016100c7565b6100fe6100f936600461099b565b6103da565b6040516100c791906109bd565b61008a6101193660046109d0565b610493565b610126610509565b6101306000610563565b565b61013a610509565b60008151116101805760405162461bcd60e51b815260206004820152600d60248201526c6b65797320697320656d70747960981b60448201526064015b60405180910390fd5b6000828152606560209081526040909120825161019f9284019061060f565b507fdfbfd50cc6a471f5290791e8d90e214ce3b67ffe9b0d9afbd7b80280f7b2891782826040516101d19291906109f9565b60405180910390a15050565b600054610100900460ff16158080156101fd5750600054600160ff909116105b806102175750303b158015610217575060005460ff166001145b61027a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610177565b6000805460ff19166001179055801561029d576000805461ff0019166101001790555b6102a56105b5565b80156102eb576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b606060656000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156103cf57838290600052602060002001805461034290610a1a565b80601f016020809104026020016040519081016040528092919081815260200182805461036e90610a1a565b80156103bb5780601f10610390576101008083540402835291602001916103bb565b820191906000526020600020905b81548152906001019060200180831161039e57829003601f168201915b505050505081526020019060010190610323565b505050509050919050565b606560205281600052604060002081815481106103f657600080fd5b9060005260206000200160009150915050805461041290610a1a565b80601f016020809104026020016040519081016040528092919081815260200182805461043e90610a1a565b801561048b5780601f106104605761010080835404028352916020019161048b565b820191906000526020600020905b81548152906001019060200180831161046e57829003601f168201915b505050505081565b61049b610509565b6001600160a01b0381166105005760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610177565b6102eb81610563565b6033546001600160a01b031633146101305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610177565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166105dc5760405162461bcd60e51b815260040161017790610a55565b610130600054610100900460ff166106065760405162461bcd60e51b815260040161017790610a55565b61013033610563565b82805482825590600052602060002090810192821561065c579160200282015b8281111561065c578251805161064c91849160209091019061066c565b509160200191906001019061062f565b506106689291506106ec565b5090565b82805461067890610a1a565b90600052602060002090601f01602090048101928261069a57600085556106e0565b82601f106106b357805160ff19168380011785556106e0565b828001600101855582156106e0579182015b828111156106e05782518255916020019190600101906106c5565b50610668929150610709565b80821115610668576000610700828261071e565b506001016106ec565b5b80821115610668576000815560010161070a565b50805461072a90610a1a565b6000825580601f1061073a575050565b601f0160209004906000526020600020908101906102eb9190610709565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561079757610797610758565b604052919050565b60008060408084860312156107b357600080fd5b8335925060208085013567ffffffffffffffff808211156107d357600080fd5b8187019150601f88818401126107e857600080fd5b8235828111156107fa576107fa610758565b8060051b61080986820161076e565b918252848101860191868101908c84111561082357600080fd5b87870192505b838310156108b0578235868111156108415760008081fd5b8701603f81018e136108535760008081fd5b888101358781111561086757610867610758565b610878818801601f19168b0161076e565b8181528f8c83850101111561088d5760008081fd5b818c84018c83013760009181018b01919091528352509187019190870190610829565b8099505050505050505050509250929050565b6000602082840312156108d557600080fd5b5035919050565b6000815180845260005b81811015610902576020818501810151868301820152016108e6565b81811115610914576000602083870101525b50601f01601f19169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b8481101561097457601f198684030189526109628383516108dc565b98840198925090830190600101610946565b5090979650505050505050565b6020815260006109946020830184610929565b9392505050565b600080604083850312156109ae57600080fd5b50508035926020909101359150565b60208152600061099460208301846108dc565b6000602082840312156109e257600080fd5b81356001600160a01b038116811461099457600080fd5b828152604060208201526000610a126040830184610929565b949350505050565b600181811c90821680610a2e57607f821691505b60208210811415610a4f57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea2646970667358221220adf285877ef73d06b1b4b52dfd4ef7b1a151a4c028550d4121f10040db975c7164736f6c634300080c0033",
}

// KeyLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyLockerMetaData.ABI instead.
var KeyLockerABI = KeyLockerMetaData.ABI

// KeyLockerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyLockerMetaData.Bin instead.
var KeyLockerBin = KeyLockerMetaData.Bin

// DeployKeyLocker deploys a new Ethereum contract, binding an instance of KeyLocker to it.
func DeployKeyLocker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyLocker, error) {
	parsed, err := KeyLockerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyLockerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyLocker{KeyLockerCaller: KeyLockerCaller{contract: contract}, KeyLockerTransactor: KeyLockerTransactor{contract: contract}, KeyLockerFilterer: KeyLockerFilterer{contract: contract}}, nil
}

// KeyLocker is an auto generated Go binding around an Ethereum contract.
type KeyLocker struct {
	KeyLockerCaller     // Read-only binding to the contract
	KeyLockerTransactor // Write-only binding to the contract
	KeyLockerFilterer   // Log filterer for contract events
}

// KeyLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyLockerSession struct {
	Contract     *KeyLocker        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyLockerCallerSession struct {
	Contract *KeyLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KeyLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyLockerTransactorSession struct {
	Contract     *KeyLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KeyLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyLockerRaw struct {
	Contract *KeyLocker // Generic contract binding to access the raw methods on
}

// KeyLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyLockerCallerRaw struct {
	Contract *KeyLockerCaller // Generic read-only contract binding to access the raw methods on
}

// KeyLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyLockerTransactorRaw struct {
	Contract *KeyLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyLocker creates a new instance of KeyLocker, bound to a specific deployed contract.
func NewKeyLocker(address common.Address, backend bind.ContractBackend) (*KeyLocker, error) {
	contract, err := bindKeyLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyLocker{KeyLockerCaller: KeyLockerCaller{contract: contract}, KeyLockerTransactor: KeyLockerTransactor{contract: contract}, KeyLockerFilterer: KeyLockerFilterer{contract: contract}}, nil
}

// NewKeyLockerCaller creates a new read-only instance of KeyLocker, bound to a specific deployed contract.
func NewKeyLockerCaller(address common.Address, caller bind.ContractCaller) (*KeyLockerCaller, error) {
	contract, err := bindKeyLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyLockerCaller{contract: contract}, nil
}

// NewKeyLockerTransactor creates a new write-only instance of KeyLocker, bound to a specific deployed contract.
func NewKeyLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyLockerTransactor, error) {
	contract, err := bindKeyLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyLockerTransactor{contract: contract}, nil
}

// NewKeyLockerFilterer creates a new log filterer instance of KeyLocker, bound to a specific deployed contract.
func NewKeyLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyLockerFilterer, error) {
	contract, err := bindKeyLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyLockerFilterer{contract: contract}, nil
}

// bindKeyLocker binds a generic wrapper to an already deployed contract.
func bindKeyLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyLockerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyLocker *KeyLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyLocker.Contract.KeyLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyLocker *KeyLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyLocker.Contract.KeyLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyLocker *KeyLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyLocker.Contract.KeyLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyLocker *KeyLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyLocker *KeyLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyLocker *KeyLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyLocker.Contract.contract.Transact(opts, method, params...)
}

// GetSocialKey is a free data retrieval call binding the contract method 0x88369c9a.
//
// Solidity: function getSocialKey(bytes32 _uuid) view returns(bytes[])
func (_KeyLocker *KeyLockerCaller) GetSocialKey(opts *bind.CallOpts, _uuid [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _KeyLocker.contract.Call(opts, &out, "getSocialKey", _uuid)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetSocialKey is a free data retrieval call binding the contract method 0x88369c9a.
//
// Solidity: function getSocialKey(bytes32 _uuid) view returns(bytes[])
func (_KeyLocker *KeyLockerSession) GetSocialKey(_uuid [32]byte) ([][]byte, error) {
	return _KeyLocker.Contract.GetSocialKey(&_KeyLocker.CallOpts, _uuid)
}

// GetSocialKey is a free data retrieval call binding the contract method 0x88369c9a.
//
// Solidity: function getSocialKey(bytes32 _uuid) view returns(bytes[])
func (_KeyLocker *KeyLockerCallerSession) GetSocialKey(_uuid [32]byte) ([][]byte, error) {
	return _KeyLocker.Contract.GetSocialKey(&_KeyLocker.CallOpts, _uuid)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyLocker *KeyLockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyLocker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyLocker *KeyLockerSession) Owner() (common.Address, error) {
	return _KeyLocker.Contract.Owner(&_KeyLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyLocker *KeyLockerCallerSession) Owner() (common.Address, error) {
	return _KeyLocker.Contract.Owner(&_KeyLocker.CallOpts)
}

// SocialKeys is a free data retrieval call binding the contract method 0xe168952b.
//
// Solidity: function socialKeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyLocker *KeyLockerCaller) SocialKeys(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _KeyLocker.contract.Call(opts, &out, "socialKeys", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SocialKeys is a free data retrieval call binding the contract method 0xe168952b.
//
// Solidity: function socialKeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyLocker *KeyLockerSession) SocialKeys(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _KeyLocker.Contract.SocialKeys(&_KeyLocker.CallOpts, arg0, arg1)
}

// SocialKeys is a free data retrieval call binding the contract method 0xe168952b.
//
// Solidity: function socialKeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyLocker *KeyLockerCallerSession) SocialKeys(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _KeyLocker.Contract.SocialKeys(&_KeyLocker.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_KeyLocker *KeyLockerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyLocker.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_KeyLocker *KeyLockerSession) Initialize() (*types.Transaction, error) {
	return _KeyLocker.Contract.Initialize(&_KeyLocker.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_KeyLocker *KeyLockerTransactorSession) Initialize() (*types.Transaction, error) {
	return _KeyLocker.Contract.Initialize(&_KeyLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyLocker *KeyLockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyLocker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyLocker *KeyLockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyLocker.Contract.RenounceOwnership(&_KeyLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyLocker *KeyLockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyLocker.Contract.RenounceOwnership(&_KeyLocker.TransactOpts)
}

// SetSocialKey is a paid mutator transaction binding the contract method 0x74aa7789.
//
// Solidity: function setSocialKey(bytes32 _uuid, bytes[] _keys) returns()
func (_KeyLocker *KeyLockerTransactor) SetSocialKey(opts *bind.TransactOpts, _uuid [32]byte, _keys [][]byte) (*types.Transaction, error) {
	return _KeyLocker.contract.Transact(opts, "setSocialKey", _uuid, _keys)
}

// SetSocialKey is a paid mutator transaction binding the contract method 0x74aa7789.
//
// Solidity: function setSocialKey(bytes32 _uuid, bytes[] _keys) returns()
func (_KeyLocker *KeyLockerSession) SetSocialKey(_uuid [32]byte, _keys [][]byte) (*types.Transaction, error) {
	return _KeyLocker.Contract.SetSocialKey(&_KeyLocker.TransactOpts, _uuid, _keys)
}

// SetSocialKey is a paid mutator transaction binding the contract method 0x74aa7789.
//
// Solidity: function setSocialKey(bytes32 _uuid, bytes[] _keys) returns()
func (_KeyLocker *KeyLockerTransactorSession) SetSocialKey(_uuid [32]byte, _keys [][]byte) (*types.Transaction, error) {
	return _KeyLocker.Contract.SetSocialKey(&_KeyLocker.TransactOpts, _uuid, _keys)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyLocker *KeyLockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyLocker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyLocker *KeyLockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyLocker.Contract.TransferOwnership(&_KeyLocker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyLocker *KeyLockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyLocker.Contract.TransferOwnership(&_KeyLocker.TransactOpts, newOwner)
}

// KeyLockerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the KeyLocker contract.
type KeyLockerInitializedIterator struct {
	Event *KeyLockerInitialized // Event containing the contract specifics and raw log

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
func (it *KeyLockerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyLockerInitialized)
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
		it.Event = new(KeyLockerInitialized)
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
func (it *KeyLockerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyLockerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyLockerInitialized represents a Initialized event raised by the KeyLocker contract.
type KeyLockerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KeyLocker *KeyLockerFilterer) FilterInitialized(opts *bind.FilterOpts) (*KeyLockerInitializedIterator, error) {

	logs, sub, err := _KeyLocker.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &KeyLockerInitializedIterator{contract: _KeyLocker.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KeyLocker *KeyLockerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *KeyLockerInitialized) (event.Subscription, error) {

	logs, sub, err := _KeyLocker.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyLockerInitialized)
				if err := _KeyLocker.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KeyLocker *KeyLockerFilterer) ParseInitialized(log types.Log) (*KeyLockerInitialized, error) {
	event := new(KeyLockerInitialized)
	if err := _KeyLocker.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyLockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyLocker contract.
type KeyLockerOwnershipTransferredIterator struct {
	Event *KeyLockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyLockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyLockerOwnershipTransferred)
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
		it.Event = new(KeyLockerOwnershipTransferred)
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
func (it *KeyLockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyLockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyLockerOwnershipTransferred represents a OwnershipTransferred event raised by the KeyLocker contract.
type KeyLockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyLocker *KeyLockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyLockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyLocker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyLockerOwnershipTransferredIterator{contract: _KeyLocker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyLocker *KeyLockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyLockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyLocker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyLockerOwnershipTransferred)
				if err := _KeyLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyLocker *KeyLockerFilterer) ParseOwnershipTransferred(log types.Log) (*KeyLockerOwnershipTransferred, error) {
	event := new(KeyLockerOwnershipTransferred)
	if err := _KeyLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyLockerKeyLockerAppendIterator is returned from FilterKeyLockerAppend and is used to iterate over the raw logs and unpacked data for KeyLockerAppend events raised by the KeyLocker contract.
type KeyLockerKeyLockerAppendIterator struct {
	Event *KeyLockerKeyLockerAppend // Event containing the contract specifics and raw log

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
func (it *KeyLockerKeyLockerAppendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyLockerKeyLockerAppend)
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
		it.Event = new(KeyLockerKeyLockerAppend)
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
func (it *KeyLockerKeyLockerAppendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyLockerKeyLockerAppendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyLockerKeyLockerAppend represents a KeyLockerAppend event raised by the KeyLocker contract.
type KeyLockerKeyLockerAppend struct {
	Uuid [32]byte
	Keys [][]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyLockerAppend is a free log retrieval operation binding the contract event 0xdfbfd50cc6a471f5290791e8d90e214ce3b67ffe9b0d9afbd7b80280f7b28917.
//
// Solidity: event keyLockerAppend(bytes32 _uuid, bytes[] _keys)
func (_KeyLocker *KeyLockerFilterer) FilterKeyLockerAppend(opts *bind.FilterOpts) (*KeyLockerKeyLockerAppendIterator, error) {

	logs, sub, err := _KeyLocker.contract.FilterLogs(opts, "keyLockerAppend")
	if err != nil {
		return nil, err
	}
	return &KeyLockerKeyLockerAppendIterator{contract: _KeyLocker.contract, event: "keyLockerAppend", logs: logs, sub: sub}, nil
}

// WatchKeyLockerAppend is a free log subscription operation binding the contract event 0xdfbfd50cc6a471f5290791e8d90e214ce3b67ffe9b0d9afbd7b80280f7b28917.
//
// Solidity: event keyLockerAppend(bytes32 _uuid, bytes[] _keys)
func (_KeyLocker *KeyLockerFilterer) WatchKeyLockerAppend(opts *bind.WatchOpts, sink chan<- *KeyLockerKeyLockerAppend) (event.Subscription, error) {

	logs, sub, err := _KeyLocker.contract.WatchLogs(opts, "keyLockerAppend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyLockerKeyLockerAppend)
				if err := _KeyLocker.contract.UnpackLog(event, "keyLockerAppend", log); err != nil {
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

// ParseKeyLockerAppend is a log parse operation binding the contract event 0xdfbfd50cc6a471f5290791e8d90e214ce3b67ffe9b0d9afbd7b80280f7b28917.
//
// Solidity: event keyLockerAppend(bytes32 _uuid, bytes[] _keys)
func (_KeyLocker *KeyLockerFilterer) ParseKeyLockerAppend(log types.Log) (*KeyLockerKeyLockerAppend, error) {
	event := new(KeyLockerKeyLockerAppend)
	if err := _KeyLocker.contract.UnpackLog(event, "keyLockerAppend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
