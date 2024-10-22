// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rootFounders

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

// Project is an auto generated low-level Go binding around an user-defined struct.
type Project struct {
	Id                  *big.Int
	Owner               common.Address
	DetailsLocationType uint8
	DetailsLocation     string
	ShortName           string
	TipJar              common.Address
}

// RootFoundersMetaData contains all meta data concerning the RootFounders contract.
var RootFoundersMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addTeammate\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTo\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"comment\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"content\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProject\",\"inputs\":[{\"name\":\"detailsLocationType\",\"type\":\"uint8\",\"internalType\":\"enumDetailsLocationType\"},{\"name\":\"detailsLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"shortName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProject\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"project\",\"type\":\"tuple\",\"internalType\":\"structProject\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"detailsLocationType\",\"type\":\"uint8\",\"internalType\":\"enumDetailsLocationType\"},{\"name\":\"detailsLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"shortName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tipJar\",\"type\":\"address\",\"internalType\":\"contractPaymentSplitter\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeTeammate\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"team\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"teammateRemoveSelf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Applied\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Commented\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"comment\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fallback\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"JoinedTeam\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LeftTeam\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"who\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProjectCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"project\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structProject\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"detailsLocationType\",\"type\":\"uint8\",\"internalType\":\"enumDetailsLocationType\"},{\"name\":\"detailsLocation\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"shortName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tipJar\",\"type\":\"address\",\"internalType\":\"contractPaymentSplitter\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// RootFoundersABI is the input ABI used to generate the binding from.
// Deprecated: Use RootFoundersMetaData.ABI instead.
var RootFoundersABI = RootFoundersMetaData.ABI

// RootFounders is an auto generated Go binding around an Ethereum contract.
type RootFounders struct {
	RootFoundersCaller     // Read-only binding to the contract
	RootFoundersTransactor // Write-only binding to the contract
	RootFoundersFilterer   // Log filterer for contract events
}

// RootFoundersCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootFoundersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootFoundersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootFoundersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootFoundersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootFoundersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootFoundersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootFoundersSession struct {
	Contract     *RootFounders     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootFoundersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootFoundersCallerSession struct {
	Contract *RootFoundersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RootFoundersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootFoundersTransactorSession struct {
	Contract     *RootFoundersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RootFoundersRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootFoundersRaw struct {
	Contract *RootFounders // Generic contract binding to access the raw methods on
}

// RootFoundersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootFoundersCallerRaw struct {
	Contract *RootFoundersCaller // Generic read-only contract binding to access the raw methods on
}

// RootFoundersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootFoundersTransactorRaw struct {
	Contract *RootFoundersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootFounders creates a new instance of RootFounders, bound to a specific deployed contract.
func NewRootFounders(address common.Address, backend bind.ContractBackend) (*RootFounders, error) {
	contract, err := bindRootFounders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootFounders{RootFoundersCaller: RootFoundersCaller{contract: contract}, RootFoundersTransactor: RootFoundersTransactor{contract: contract}, RootFoundersFilterer: RootFoundersFilterer{contract: contract}}, nil
}

// NewRootFoundersCaller creates a new read-only instance of RootFounders, bound to a specific deployed contract.
func NewRootFoundersCaller(address common.Address, caller bind.ContractCaller) (*RootFoundersCaller, error) {
	contract, err := bindRootFounders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootFoundersCaller{contract: contract}, nil
}

// NewRootFoundersTransactor creates a new write-only instance of RootFounders, bound to a specific deployed contract.
func NewRootFoundersTransactor(address common.Address, transactor bind.ContractTransactor) (*RootFoundersTransactor, error) {
	contract, err := bindRootFounders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootFoundersTransactor{contract: contract}, nil
}

// NewRootFoundersFilterer creates a new log filterer instance of RootFounders, bound to a specific deployed contract.
func NewRootFoundersFilterer(address common.Address, filterer bind.ContractFilterer) (*RootFoundersFilterer, error) {
	contract, err := bindRootFounders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootFoundersFilterer{contract: contract}, nil
}

// bindRootFounders binds a generic wrapper to an already deployed contract.
func bindRootFounders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RootFoundersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootFounders *RootFoundersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootFounders.Contract.RootFoundersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootFounders *RootFoundersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootFounders.Contract.RootFoundersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootFounders *RootFoundersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootFounders.Contract.RootFoundersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootFounders *RootFoundersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootFounders.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootFounders *RootFoundersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootFounders.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootFounders *RootFoundersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootFounders.Contract.contract.Transact(opts, method, params...)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 id) view returns((uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersCaller) GetProject(opts *bind.CallOpts, id *big.Int) (Project, error) {
	var out []interface{}
	err := _RootFounders.contract.Call(opts, &out, "getProject", id)

	if err != nil {
		return *new(Project), err
	}

	out0 := *abi.ConvertType(out[0], new(Project)).(*Project)

	return out0, err

}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 id) view returns((uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersSession) GetProject(id *big.Int) (Project, error) {
	return _RootFounders.Contract.GetProject(&_RootFounders.CallOpts, id)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(uint256 id) view returns((uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersCallerSession) GetProject(id *big.Int) (Project, error) {
	return _RootFounders.Contract.GetProject(&_RootFounders.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootFounders *RootFoundersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RootFounders.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootFounders *RootFoundersSession) Owner() (common.Address, error) {
	return _RootFounders.Contract.Owner(&_RootFounders.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RootFounders *RootFoundersCallerSession) Owner() (common.Address, error) {
	return _RootFounders.Contract.Owner(&_RootFounders.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x197ebd53.
//
// Solidity: function team(uint256 projectId) view returns(address[] addresses)
func (_RootFounders *RootFoundersCaller) Team(opts *bind.CallOpts, projectId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RootFounders.contract.Call(opts, &out, "team", projectId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Team is a free data retrieval call binding the contract method 0x197ebd53.
//
// Solidity: function team(uint256 projectId) view returns(address[] addresses)
func (_RootFounders *RootFoundersSession) Team(projectId *big.Int) ([]common.Address, error) {
	return _RootFounders.Contract.Team(&_RootFounders.CallOpts, projectId)
}

// Team is a free data retrieval call binding the contract method 0x197ebd53.
//
// Solidity: function team(uint256 projectId) view returns(address[] addresses)
func (_RootFounders *RootFoundersCallerSession) Team(projectId *big.Int) ([]common.Address, error) {
	return _RootFounders.Contract.Team(&_RootFounders.CallOpts, projectId)
}

// AddTeammate is a paid mutator transaction binding the contract method 0xb0c7542a.
//
// Solidity: function addTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersTransactor) AddTeammate(opts *bind.TransactOpts, projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "addTeammate", projectId, mate)
}

// AddTeammate is a paid mutator transaction binding the contract method 0xb0c7542a.
//
// Solidity: function addTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersSession) AddTeammate(projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.AddTeammate(&_RootFounders.TransactOpts, projectId, mate)
}

// AddTeammate is a paid mutator transaction binding the contract method 0xb0c7542a.
//
// Solidity: function addTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersTransactorSession) AddTeammate(projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.AddTeammate(&_RootFounders.TransactOpts, projectId, mate)
}

// ApplyTo is a paid mutator transaction binding the contract method 0x37a057b8.
//
// Solidity: function applyTo(uint256 projectId) returns(bool)
func (_RootFounders *RootFoundersTransactor) ApplyTo(opts *bind.TransactOpts, projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "applyTo", projectId)
}

// ApplyTo is a paid mutator transaction binding the contract method 0x37a057b8.
//
// Solidity: function applyTo(uint256 projectId) returns(bool)
func (_RootFounders *RootFoundersSession) ApplyTo(projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.ApplyTo(&_RootFounders.TransactOpts, projectId)
}

// ApplyTo is a paid mutator transaction binding the contract method 0x37a057b8.
//
// Solidity: function applyTo(uint256 projectId) returns(bool)
func (_RootFounders *RootFoundersTransactorSession) ApplyTo(projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.ApplyTo(&_RootFounders.TransactOpts, projectId)
}

// Comment is a paid mutator transaction binding the contract method 0x1805c523.
//
// Solidity: function comment(uint256 id, string content) returns()
func (_RootFounders *RootFoundersTransactor) Comment(opts *bind.TransactOpts, id *big.Int, content string) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "comment", id, content)
}

// Comment is a paid mutator transaction binding the contract method 0x1805c523.
//
// Solidity: function comment(uint256 id, string content) returns()
func (_RootFounders *RootFoundersSession) Comment(id *big.Int, content string) (*types.Transaction, error) {
	return _RootFounders.Contract.Comment(&_RootFounders.TransactOpts, id, content)
}

// Comment is a paid mutator transaction binding the contract method 0x1805c523.
//
// Solidity: function comment(uint256 id, string content) returns()
func (_RootFounders *RootFoundersTransactorSession) Comment(id *big.Int, content string) (*types.Transaction, error) {
	return _RootFounders.Contract.Comment(&_RootFounders.TransactOpts, id, content)
}

// CreateProject is a paid mutator transaction binding the contract method 0xdd5fdc60.
//
// Solidity: function createProject(uint8 detailsLocationType, string detailsLocation, string shortName) returns(uint256 id)
func (_RootFounders *RootFoundersTransactor) CreateProject(opts *bind.TransactOpts, detailsLocationType uint8, detailsLocation string, shortName string) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "createProject", detailsLocationType, detailsLocation, shortName)
}

// CreateProject is a paid mutator transaction binding the contract method 0xdd5fdc60.
//
// Solidity: function createProject(uint8 detailsLocationType, string detailsLocation, string shortName) returns(uint256 id)
func (_RootFounders *RootFoundersSession) CreateProject(detailsLocationType uint8, detailsLocation string, shortName string) (*types.Transaction, error) {
	return _RootFounders.Contract.CreateProject(&_RootFounders.TransactOpts, detailsLocationType, detailsLocation, shortName)
}

// CreateProject is a paid mutator transaction binding the contract method 0xdd5fdc60.
//
// Solidity: function createProject(uint8 detailsLocationType, string detailsLocation, string shortName) returns(uint256 id)
func (_RootFounders *RootFoundersTransactorSession) CreateProject(detailsLocationType uint8, detailsLocation string, shortName string) (*types.Transaction, error) {
	return _RootFounders.Contract.CreateProject(&_RootFounders.TransactOpts, detailsLocationType, detailsLocation, shortName)
}

// RemoveTeammate is a paid mutator transaction binding the contract method 0x839f76d9.
//
// Solidity: function removeTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersTransactor) RemoveTeammate(opts *bind.TransactOpts, projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "removeTeammate", projectId, mate)
}

// RemoveTeammate is a paid mutator transaction binding the contract method 0x839f76d9.
//
// Solidity: function removeTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersSession) RemoveTeammate(projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.RemoveTeammate(&_RootFounders.TransactOpts, projectId, mate)
}

// RemoveTeammate is a paid mutator transaction binding the contract method 0x839f76d9.
//
// Solidity: function removeTeammate(uint256 projectId, address mate) returns()
func (_RootFounders *RootFoundersTransactorSession) RemoveTeammate(projectId *big.Int, mate common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.RemoveTeammate(&_RootFounders.TransactOpts, projectId, mate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootFounders *RootFoundersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootFounders *RootFoundersSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootFounders.Contract.RenounceOwnership(&_RootFounders.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RootFounders *RootFoundersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RootFounders.Contract.RenounceOwnership(&_RootFounders.TransactOpts)
}

// TeammateRemoveSelf is a paid mutator transaction binding the contract method 0xc118e742.
//
// Solidity: function teammateRemoveSelf(uint256 projectId) returns()
func (_RootFounders *RootFoundersTransactor) TeammateRemoveSelf(opts *bind.TransactOpts, projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "teammateRemoveSelf", projectId)
}

// TeammateRemoveSelf is a paid mutator transaction binding the contract method 0xc118e742.
//
// Solidity: function teammateRemoveSelf(uint256 projectId) returns()
func (_RootFounders *RootFoundersSession) TeammateRemoveSelf(projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.TeammateRemoveSelf(&_RootFounders.TransactOpts, projectId)
}

// TeammateRemoveSelf is a paid mutator transaction binding the contract method 0xc118e742.
//
// Solidity: function teammateRemoveSelf(uint256 projectId) returns()
func (_RootFounders *RootFoundersTransactorSession) TeammateRemoveSelf(projectId *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.TeammateRemoveSelf(&_RootFounders.TransactOpts, projectId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootFounders *RootFoundersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootFounders *RootFoundersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.TransferOwnership(&_RootFounders.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RootFounders *RootFoundersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RootFounders.Contract.TransferOwnership(&_RootFounders.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RootFounders *RootFoundersTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _RootFounders.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RootFounders *RootFoundersSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.Withdraw(&_RootFounders.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_RootFounders *RootFoundersTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _RootFounders.Contract.Withdraw(&_RootFounders.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RootFounders *RootFoundersTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RootFounders.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RootFounders *RootFoundersSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RootFounders.Contract.Fallback(&_RootFounders.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RootFounders *RootFoundersTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RootFounders.Contract.Fallback(&_RootFounders.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootFounders *RootFoundersTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootFounders.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootFounders *RootFoundersSession) Receive() (*types.Transaction, error) {
	return _RootFounders.Contract.Receive(&_RootFounders.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootFounders *RootFoundersTransactorSession) Receive() (*types.Transaction, error) {
	return _RootFounders.Contract.Receive(&_RootFounders.TransactOpts)
}

// RootFoundersAppliedIterator is returned from FilterApplied and is used to iterate over the raw logs and unpacked data for Applied events raised by the RootFounders contract.
type RootFoundersAppliedIterator struct {
	Event *RootFoundersApplied // Event containing the contract specifics and raw log

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
func (it *RootFoundersAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersApplied)
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
		it.Event = new(RootFoundersApplied)
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
func (it *RootFoundersAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersApplied represents a Applied event raised by the RootFounders contract.
type RootFoundersApplied struct {
	ProjectId *big.Int
	Who       common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApplied is a free log retrieval operation binding the contract event 0x8e536842b0bedb99bbe3d9b2f4d7cd3c1d89955698940f266371b4d4f5caa257.
//
// Solidity: event Applied(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) FilterApplied(opts *bind.FilterOpts, projectId []*big.Int, who []common.Address) (*RootFoundersAppliedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "Applied", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersAppliedIterator{contract: _RootFounders.contract, event: "Applied", logs: logs, sub: sub}, nil
}

// WatchApplied is a free log subscription operation binding the contract event 0x8e536842b0bedb99bbe3d9b2f4d7cd3c1d89955698940f266371b4d4f5caa257.
//
// Solidity: event Applied(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) WatchApplied(opts *bind.WatchOpts, sink chan<- *RootFoundersApplied, projectId []*big.Int, who []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "Applied", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersApplied)
				if err := _RootFounders.contract.UnpackLog(event, "Applied", log); err != nil {
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

// ParseApplied is a log parse operation binding the contract event 0x8e536842b0bedb99bbe3d9b2f4d7cd3c1d89955698940f266371b4d4f5caa257.
//
// Solidity: event Applied(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) ParseApplied(log types.Log) (*RootFoundersApplied, error) {
	event := new(RootFoundersApplied)
	if err := _RootFounders.contract.UnpackLog(event, "Applied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersCommentedIterator is returned from FilterCommented and is used to iterate over the raw logs and unpacked data for Commented events raised by the RootFounders contract.
type RootFoundersCommentedIterator struct {
	Event *RootFoundersCommented // Event containing the contract specifics and raw log

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
func (it *RootFoundersCommentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersCommented)
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
		it.Event = new(RootFoundersCommented)
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
func (it *RootFoundersCommentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersCommentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersCommented represents a Commented event raised by the RootFounders contract.
type RootFoundersCommented struct {
	From      common.Address
	ProjectId *big.Int
	Comment   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommented is a free log retrieval operation binding the contract event 0x97f4ab7401ce70b8ac852732b9204e5ef745964fc4e8577abb5fe1bbaf2fcd9c.
//
// Solidity: event Commented(address indexed from, uint256 indexed projectId, string comment)
func (_RootFounders *RootFoundersFilterer) FilterCommented(opts *bind.FilterOpts, from []common.Address, projectId []*big.Int) (*RootFoundersCommentedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "Commented", fromRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersCommentedIterator{contract: _RootFounders.contract, event: "Commented", logs: logs, sub: sub}, nil
}

// WatchCommented is a free log subscription operation binding the contract event 0x97f4ab7401ce70b8ac852732b9204e5ef745964fc4e8577abb5fe1bbaf2fcd9c.
//
// Solidity: event Commented(address indexed from, uint256 indexed projectId, string comment)
func (_RootFounders *RootFoundersFilterer) WatchCommented(opts *bind.WatchOpts, sink chan<- *RootFoundersCommented, from []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "Commented", fromRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersCommented)
				if err := _RootFounders.contract.UnpackLog(event, "Commented", log); err != nil {
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

// ParseCommented is a log parse operation binding the contract event 0x97f4ab7401ce70b8ac852732b9204e5ef745964fc4e8577abb5fe1bbaf2fcd9c.
//
// Solidity: event Commented(address indexed from, uint256 indexed projectId, string comment)
func (_RootFounders *RootFoundersFilterer) ParseCommented(log types.Log) (*RootFoundersCommented, error) {
	event := new(RootFoundersCommented)
	if err := _RootFounders.contract.UnpackLog(event, "Commented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersFallbackIterator is returned from FilterFallback and is used to iterate over the raw logs and unpacked data for Fallback events raised by the RootFounders contract.
type RootFoundersFallbackIterator struct {
	Event *RootFoundersFallback // Event containing the contract specifics and raw log

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
func (it *RootFoundersFallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersFallback)
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
		it.Event = new(RootFoundersFallback)
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
func (it *RootFoundersFallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersFallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersFallback represents a Fallback event raised by the RootFounders contract.
type RootFoundersFallback struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFallback is a free log retrieval operation binding the contract event 0xfbf15a1bae5e021d024841007b692b167afd2a281a4ff0b44f47387eb388205c.
//
// Solidity: event Fallback(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) FilterFallback(opts *bind.FilterOpts) (*RootFoundersFallbackIterator, error) {

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "Fallback")
	if err != nil {
		return nil, err
	}
	return &RootFoundersFallbackIterator{contract: _RootFounders.contract, event: "Fallback", logs: logs, sub: sub}, nil
}

// WatchFallback is a free log subscription operation binding the contract event 0xfbf15a1bae5e021d024841007b692b167afd2a281a4ff0b44f47387eb388205c.
//
// Solidity: event Fallback(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) WatchFallback(opts *bind.WatchOpts, sink chan<- *RootFoundersFallback) (event.Subscription, error) {

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "Fallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersFallback)
				if err := _RootFounders.contract.UnpackLog(event, "Fallback", log); err != nil {
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

// ParseFallback is a log parse operation binding the contract event 0xfbf15a1bae5e021d024841007b692b167afd2a281a4ff0b44f47387eb388205c.
//
// Solidity: event Fallback(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) ParseFallback(log types.Log) (*RootFoundersFallback, error) {
	event := new(RootFoundersFallback)
	if err := _RootFounders.contract.UnpackLog(event, "Fallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersJoinedTeamIterator is returned from FilterJoinedTeam and is used to iterate over the raw logs and unpacked data for JoinedTeam events raised by the RootFounders contract.
type RootFoundersJoinedTeamIterator struct {
	Event *RootFoundersJoinedTeam // Event containing the contract specifics and raw log

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
func (it *RootFoundersJoinedTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersJoinedTeam)
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
		it.Event = new(RootFoundersJoinedTeam)
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
func (it *RootFoundersJoinedTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersJoinedTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersJoinedTeam represents a JoinedTeam event raised by the RootFounders contract.
type RootFoundersJoinedTeam struct {
	ProjectId *big.Int
	Who       common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJoinedTeam is a free log retrieval operation binding the contract event 0x4b543b179253227f7b5766055f0cde50bb871401f45ac8a0f6aa75015ebe12db.
//
// Solidity: event JoinedTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) FilterJoinedTeam(opts *bind.FilterOpts, projectId []*big.Int, who []common.Address) (*RootFoundersJoinedTeamIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "JoinedTeam", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersJoinedTeamIterator{contract: _RootFounders.contract, event: "JoinedTeam", logs: logs, sub: sub}, nil
}

// WatchJoinedTeam is a free log subscription operation binding the contract event 0x4b543b179253227f7b5766055f0cde50bb871401f45ac8a0f6aa75015ebe12db.
//
// Solidity: event JoinedTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) WatchJoinedTeam(opts *bind.WatchOpts, sink chan<- *RootFoundersJoinedTeam, projectId []*big.Int, who []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "JoinedTeam", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersJoinedTeam)
				if err := _RootFounders.contract.UnpackLog(event, "JoinedTeam", log); err != nil {
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

// ParseJoinedTeam is a log parse operation binding the contract event 0x4b543b179253227f7b5766055f0cde50bb871401f45ac8a0f6aa75015ebe12db.
//
// Solidity: event JoinedTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) ParseJoinedTeam(log types.Log) (*RootFoundersJoinedTeam, error) {
	event := new(RootFoundersJoinedTeam)
	if err := _RootFounders.contract.UnpackLog(event, "JoinedTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersLeftTeamIterator is returned from FilterLeftTeam and is used to iterate over the raw logs and unpacked data for LeftTeam events raised by the RootFounders contract.
type RootFoundersLeftTeamIterator struct {
	Event *RootFoundersLeftTeam // Event containing the contract specifics and raw log

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
func (it *RootFoundersLeftTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersLeftTeam)
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
		it.Event = new(RootFoundersLeftTeam)
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
func (it *RootFoundersLeftTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersLeftTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersLeftTeam represents a LeftTeam event raised by the RootFounders contract.
type RootFoundersLeftTeam struct {
	ProjectId *big.Int
	Who       common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLeftTeam is a free log retrieval operation binding the contract event 0x3e2c70cc3a3b9ab2bc8e5fbc796e0d6eb6617e77b1d17043aea5e44c9cc06ac6.
//
// Solidity: event LeftTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) FilterLeftTeam(opts *bind.FilterOpts, projectId []*big.Int, who []common.Address) (*RootFoundersLeftTeamIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "LeftTeam", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersLeftTeamIterator{contract: _RootFounders.contract, event: "LeftTeam", logs: logs, sub: sub}, nil
}

// WatchLeftTeam is a free log subscription operation binding the contract event 0x3e2c70cc3a3b9ab2bc8e5fbc796e0d6eb6617e77b1d17043aea5e44c9cc06ac6.
//
// Solidity: event LeftTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) WatchLeftTeam(opts *bind.WatchOpts, sink chan<- *RootFoundersLeftTeam, projectId []*big.Int, who []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "LeftTeam", projectIdRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersLeftTeam)
				if err := _RootFounders.contract.UnpackLog(event, "LeftTeam", log); err != nil {
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

// ParseLeftTeam is a log parse operation binding the contract event 0x3e2c70cc3a3b9ab2bc8e5fbc796e0d6eb6617e77b1d17043aea5e44c9cc06ac6.
//
// Solidity: event LeftTeam(uint256 indexed projectId, address indexed who)
func (_RootFounders *RootFoundersFilterer) ParseLeftTeam(log types.Log) (*RootFoundersLeftTeam, error) {
	event := new(RootFoundersLeftTeam)
	if err := _RootFounders.contract.UnpackLog(event, "LeftTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RootFounders contract.
type RootFoundersOwnershipTransferredIterator struct {
	Event *RootFoundersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RootFoundersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersOwnershipTransferred)
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
		it.Event = new(RootFoundersOwnershipTransferred)
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
func (it *RootFoundersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersOwnershipTransferred represents a OwnershipTransferred event raised by the RootFounders contract.
type RootFoundersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootFounders *RootFoundersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RootFoundersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersOwnershipTransferredIterator{contract: _RootFounders.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RootFounders *RootFoundersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RootFoundersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersOwnershipTransferred)
				if err := _RootFounders.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RootFounders *RootFoundersFilterer) ParseOwnershipTransferred(log types.Log) (*RootFoundersOwnershipTransferred, error) {
	event := new(RootFoundersOwnershipTransferred)
	if err := _RootFounders.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the RootFounders contract.
type RootFoundersProjectCreatedIterator struct {
	Event *RootFoundersProjectCreated // Event containing the contract specifics and raw log

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
func (it *RootFoundersProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersProjectCreated)
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
		it.Event = new(RootFoundersProjectCreated)
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
func (it *RootFoundersProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersProjectCreated represents a ProjectCreated event raised by the RootFounders contract.
type RootFoundersProjectCreated struct {
	Id      *big.Int
	Owner   common.Address
	Project Project
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0xe5631e3ccead5c67d9805b558477152cd2f010c3080044a7efdec83dcae4aa88.
//
// Solidity: event ProjectCreated(uint256 indexed id, address indexed owner, (uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersFilterer) FilterProjectCreated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*RootFoundersProjectCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "ProjectCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RootFoundersProjectCreatedIterator{contract: _RootFounders.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0xe5631e3ccead5c67d9805b558477152cd2f010c3080044a7efdec83dcae4aa88.
//
// Solidity: event ProjectCreated(uint256 indexed id, address indexed owner, (uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *RootFoundersProjectCreated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "ProjectCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersProjectCreated)
				if err := _RootFounders.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
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

// ParseProjectCreated is a log parse operation binding the contract event 0xe5631e3ccead5c67d9805b558477152cd2f010c3080044a7efdec83dcae4aa88.
//
// Solidity: event ProjectCreated(uint256 indexed id, address indexed owner, (uint256,address,uint8,string,string,address) project)
func (_RootFounders *RootFoundersFilterer) ParseProjectCreated(log types.Log) (*RootFoundersProjectCreated, error) {
	event := new(RootFoundersProjectCreated)
	if err := _RootFounders.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RootFoundersReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the RootFounders contract.
type RootFoundersReceivedIterator struct {
	Event *RootFoundersReceived // Event containing the contract specifics and raw log

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
func (it *RootFoundersReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootFoundersReceived)
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
		it.Event = new(RootFoundersReceived)
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
func (it *RootFoundersReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootFoundersReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootFoundersReceived represents a Received event raised by the RootFounders contract.
type RootFoundersReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) FilterReceived(opts *bind.FilterOpts) (*RootFoundersReceivedIterator, error) {

	logs, sub, err := _RootFounders.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &RootFoundersReceivedIterator{contract: _RootFounders.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *RootFoundersReceived) (event.Subscription, error) {

	logs, sub, err := _RootFounders.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootFoundersReceived)
				if err := _RootFounders.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_RootFounders *RootFoundersFilterer) ParseReceived(log types.Log) (*RootFoundersReceived, error) {
	event := new(RootFoundersReceived)
	if err := _RootFounders.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
