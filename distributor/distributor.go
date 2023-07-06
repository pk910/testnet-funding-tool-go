// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distributor

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

// DistributorMetaData contains all meta data concerning the Distributor contract.
var DistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addrs\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addrs\",\"type\":\"bytes\"}],\"name\":\"distributeEqual\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addrs\",\"type\":\"bytes\"},{\"internalType\":\"uint32[]\",\"name\":\"values\",\"type\":\"uint32[]\"}],\"name\":\"distributeEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addrs\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"values\",\"type\":\"uint64[]\"}],\"name\":\"distributeGwei\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106a6806100206000396000f3fe60806040526004361061003f5760003560e01c806337b4a7b214610044578063985b7d6b14610059578063cb2223021461006c578063f85409631461007f575b600080fd5b6100576100523660046104b2565b610092565b005b6100576100673660046104b2565b6101c6565b61005761007a36600461051e565b6102a4565b61005761008d3660046104b2565b610361565b826000805b8263ffffffff168263ffffffff1610156101bd5761010f8763ffffffff8416886100c2866014610576565b63ffffffff16926100d59392919061059a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061041d92505050565b73ffffffffffffffffffffffffffffffffffffffff1685858363ffffffff1681811061013d5761013d6105c4565b905060200201602081019061015291906105da565b61016a9063ffffffff16670de0b6b3a7640000610607565b604051600081818185875af1925050503d80600081146101a6576040519150601f19603f3d011682016040523d82523d6000602084013e6101ab565b606091505b50505060149190910190600101610097565b50505050505050565b826000805b8263ffffffff168263ffffffff1610156101bd576101f68763ffffffff8416886100c2866014610576565b73ffffffffffffffffffffffffffffffffffffffff1685858363ffffffff16818110610224576102246105c4565b90506020020160208101906102399190610624565b6102519067ffffffffffffffff16633b9aca00610607565b604051600081818185875af1925050503d806000811461028d576040519150601f19603f3d011682016040523d82523d6000602084013e610292565b606091505b505050601491909101906001016101cb565b80600063ffffffff82166102b9346014610607565b6102c3919061064e565b905060005b8263ffffffff168163ffffffff16101561035a576102f38563ffffffff8316866100c2856014610576565b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461034a576040519150601f19603f3d011682016040523d82523d6000602084013e61034f565b606091505b5050506014016102c8565b5050505050565b826000805b8263ffffffff168263ffffffff1610156101bd576103918763ffffffff8416886100c2866014610576565b73ffffffffffffffffffffffffffffffffffffffff1685858363ffffffff168181106103bf576103bf6105c4565b9050602002013560405160006040518083038185875af1925050503d8060008114610406576040519150601f19603f3d011682016040523d82523d6000602084013e61040b565b606091505b50505060149190910190600101610366565b6014015190565b60008083601f84011261043657600080fd5b50813567ffffffffffffffff81111561044e57600080fd5b60208301915083602082850101111561046657600080fd5b9250929050565b60008083601f84011261047f57600080fd5b50813567ffffffffffffffff81111561049757600080fd5b6020830191508360208260051b850101111561046657600080fd5b600080600080604085870312156104c857600080fd5b843567ffffffffffffffff808211156104e057600080fd5b6104ec88838901610424565b9096509450602087013591508082111561050557600080fd5b506105128782880161046d565b95989497509550505050565b6000806020838503121561053157600080fd5b823567ffffffffffffffff81111561054857600080fd5b61055485828601610424565b90969095509350505050565b634e487b7160e01b600052601160045260246000fd5b63ffffffff81811683821601908082111561059357610593610560565b5092915050565b600080858511156105aa57600080fd5b838611156105b757600080fd5b5050820193919092039150565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156105ec57600080fd5b813563ffffffff8116811461060057600080fd5b9392505050565b808202811582820484141761061e5761061e610560565b92915050565b60006020828403121561063657600080fd5b813567ffffffffffffffff8116811461060057600080fd5b60008261066b57634e487b7160e01b600052601260045260246000fd5b50049056fea264697066735822122064a073572ba17daeef9b562533a4b9d5bffaa2106d3f69d0367214fedf1c2f9164736f6c63430008120033",
}

// DistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributorMetaData.ABI instead.
var DistributorABI = DistributorMetaData.ABI

// DistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributorMetaData.Bin instead.
var DistributorBin = DistributorMetaData.Bin

// DeployDistributor deploys a new Ethereum contract, binding an instance of Distributor to it.
func DeployDistributor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Distributor, error) {
	parsed, err := DistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// Distributor is an auto generated Go binding around an Ethereum contract.
type Distributor struct {
	DistributorCaller     // Read-only binding to the contract
	DistributorTransactor // Write-only binding to the contract
	DistributorFilterer   // Log filterer for contract events
}

// DistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributorSession struct {
	Contract     *Distributor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributorCallerSession struct {
	Contract *DistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributorTransactorSession struct {
	Contract     *DistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributorRaw struct {
	Contract *Distributor // Generic contract binding to access the raw methods on
}

// DistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributorCallerRaw struct {
	Contract *DistributorCaller // Generic read-only contract binding to access the raw methods on
}

// DistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributorTransactorRaw struct {
	Contract *DistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributor creates a new instance of Distributor, bound to a specific deployed contract.
func NewDistributor(address common.Address, backend bind.ContractBackend) (*Distributor, error) {
	contract, err := bindDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// NewDistributorCaller creates a new read-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorCaller(address common.Address, caller bind.ContractCaller) (*DistributorCaller, error) {
	contract, err := bindDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorCaller{contract: contract}, nil
}

// NewDistributorTransactor creates a new write-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributorTransactor, error) {
	contract, err := bindDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorTransactor{contract: contract}, nil
}

// NewDistributorFilterer creates a new log filterer instance of Distributor, bound to a specific deployed contract.
func NewDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributorFilterer, error) {
	contract, err := bindDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributorFilterer{contract: contract}, nil
}

// bindDistributor binds a generic wrapper to an already deployed contract.
func bindDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.DistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transact(opts, method, params...)
}

// Distribute is a paid mutator transaction binding the contract method 0xf8540963.
//
// Solidity: function distribute(bytes addrs, uint256[] values) payable returns()
func (_Distributor *DistributorTransactor) Distribute(opts *bind.TransactOpts, addrs []byte, values []*big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distribute", addrs, values)
}

// Distribute is a paid mutator transaction binding the contract method 0xf8540963.
//
// Solidity: function distribute(bytes addrs, uint256[] values) payable returns()
func (_Distributor *DistributorSession) Distribute(addrs []byte, values []*big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Distribute(&_Distributor.TransactOpts, addrs, values)
}

// Distribute is a paid mutator transaction binding the contract method 0xf8540963.
//
// Solidity: function distribute(bytes addrs, uint256[] values) payable returns()
func (_Distributor *DistributorTransactorSession) Distribute(addrs []byte, values []*big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Distribute(&_Distributor.TransactOpts, addrs, values)
}

// DistributeEqual is a paid mutator transaction binding the contract method 0xcb222302.
//
// Solidity: function distributeEqual(bytes addrs) payable returns()
func (_Distributor *DistributorTransactor) DistributeEqual(opts *bind.TransactOpts, addrs []byte) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeEqual", addrs)
}

// DistributeEqual is a paid mutator transaction binding the contract method 0xcb222302.
//
// Solidity: function distributeEqual(bytes addrs) payable returns()
func (_Distributor *DistributorSession) DistributeEqual(addrs []byte) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeEqual(&_Distributor.TransactOpts, addrs)
}

// DistributeEqual is a paid mutator transaction binding the contract method 0xcb222302.
//
// Solidity: function distributeEqual(bytes addrs) payable returns()
func (_Distributor *DistributorTransactorSession) DistributeEqual(addrs []byte) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeEqual(&_Distributor.TransactOpts, addrs)
}

// DistributeEther is a paid mutator transaction binding the contract method 0x37b4a7b2.
//
// Solidity: function distributeEther(bytes addrs, uint32[] values) payable returns()
func (_Distributor *DistributorTransactor) DistributeEther(opts *bind.TransactOpts, addrs []byte, values []uint32) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeEther", addrs, values)
}

// DistributeEther is a paid mutator transaction binding the contract method 0x37b4a7b2.
//
// Solidity: function distributeEther(bytes addrs, uint32[] values) payable returns()
func (_Distributor *DistributorSession) DistributeEther(addrs []byte, values []uint32) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeEther(&_Distributor.TransactOpts, addrs, values)
}

// DistributeEther is a paid mutator transaction binding the contract method 0x37b4a7b2.
//
// Solidity: function distributeEther(bytes addrs, uint32[] values) payable returns()
func (_Distributor *DistributorTransactorSession) DistributeEther(addrs []byte, values []uint32) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeEther(&_Distributor.TransactOpts, addrs, values)
}

// DistributeGwei is a paid mutator transaction binding the contract method 0x985b7d6b.
//
// Solidity: function distributeGwei(bytes addrs, uint64[] values) payable returns()
func (_Distributor *DistributorTransactor) DistributeGwei(opts *bind.TransactOpts, addrs []byte, values []uint64) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distributeGwei", addrs, values)
}

// DistributeGwei is a paid mutator transaction binding the contract method 0x985b7d6b.
//
// Solidity: function distributeGwei(bytes addrs, uint64[] values) payable returns()
func (_Distributor *DistributorSession) DistributeGwei(addrs []byte, values []uint64) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeGwei(&_Distributor.TransactOpts, addrs, values)
}

// DistributeGwei is a paid mutator transaction binding the contract method 0x985b7d6b.
//
// Solidity: function distributeGwei(bytes addrs, uint64[] values) payable returns()
func (_Distributor *DistributorTransactorSession) DistributeGwei(addrs []byte, values []uint64) (*types.Transaction, error) {
	return _Distributor.Contract.DistributeGwei(&_Distributor.TransactOpts, addrs, values)
}
