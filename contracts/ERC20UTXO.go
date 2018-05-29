// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ERC20UTXOABI is the input ABI used to generate the binding from.
const ERC20UTXOABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"utxos\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"createdBy\",\"type\":\"bytes32\"},{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getUtxo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"spend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"LogSpend\",\"type\":\"event\"}]"

// ERC20UTXOBin is the compiled bytecode used for deploying new contracts.
const ERC20UTXOBin = `0x6060604052341561000f57600080fd5b6106878061001e6000396000f3006060604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ecaea7381146100665780633cf18ab91461008a578063c972949b146100d6578063de86ff131461011a575b600080fd5b341561007157600080fd5b610088600160a060020a036004351660243561013f565b005b341561009557600080fd5b6100a060043561025a565b604051600160a060020a039094168452602084019290925260408084019190915260608301919091526080909101905180910390f35b34156100e157600080fd5b6100ec60043561028b565b604051600160a060020a03909316835260208301919091526040808301919091526060909101905180910390f35b341561012557600080fd5b610088600435602435600160a060020a03604435166102fc565b6000610149610634565b4333856040519283526c01000000000000000000000000600160a060020a039283168102602085015291160260348201526048016040518091039020915060806040519081016040908152600160a060020a03861682526020808301869052600082840181905260608401869052858152908190522090915081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201516002820155606082015160039091015550600180548401905581600160a060020a0385167fa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac633848560405190815260200160405180910390a350505050565b6000602081905290815260409020805460018201546002830154600390930154600160a060020a0390921692909184565b6000806000610298610634565b60008581526020819052604090819020906080905190810160409081528254600160a060020a031682526001830154602083015260028301549082015260039091015460608201529050805181602001518260400151935093509350509193909250565b610304610634565b600061030e610634565b6000610318610634565b60008881526020819052604090205433600160a060020a0390811691161461033f57600080fd5b6000888152602081905260409020600101548790101561035e57600080fd5b60008881526020819052604090819020906080905190810160409081528254600160a060020a0316825260018084015460208085019190915260028086015484860152600395860154606086015260008e815291829052928120805473ffffffffffffffffffffffffffffffffffffffff191681559182018190559181018290559092019190915594506103f286896105e6565b935060806040519081016040908152600160a060020a038816825260208083018a90528183018b9052606083018790526000878152908190522090935083908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201516002820155606082015160039091015550600160a060020a038087169033167fa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d186060880151878b60405192835260208301919091526040808301919091526060909101905180910390a384602001518710156105dc576104ed3360018a186105e6565b915060806040519081016040528033600160a060020a0316815260200188876020015103815260208082018b905260409182018590526000858152908190522090915081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201516002820155606082015160039091015550600160a060020a038087169033167fa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d186060880151858b8a602001510360405192835260208301919091526040808301919091526060909101905180910390a35b5050505050505050565b6000433384846040519384526c01000000000000000000000000600160a060020a03938416810260208601529190921602603483015260488201526068016040518091039020905092915050565b608060405190810160409081526000808352602083018190529082018190526060820152905600a165627a7a72305820cd6f95d63c272aba6c93304565d72086dd6deb20e5085af488929ff9e4cda0cb0029`

// DeployERC20UTXO deploys a new Ethereum contract, binding an instance of ERC20UTXO to it.
func DeployERC20UTXO(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20UTXO, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UTXOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20UTXOBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20UTXO{ERC20UTXOCaller: ERC20UTXOCaller{contract: contract}, ERC20UTXOTransactor: ERC20UTXOTransactor{contract: contract}, ERC20UTXOFilterer: ERC20UTXOFilterer{contract: contract}}, nil
}

// ERC20UTXO is an auto generated Go binding around an Ethereum contract.
type ERC20UTXO struct {
	ERC20UTXOCaller     // Read-only binding to the contract
	ERC20UTXOTransactor // Write-only binding to the contract
	ERC20UTXOFilterer   // Log filterer for contract events
}

// ERC20UTXOCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20UTXOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20UTXOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20UTXOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20UTXOSession struct {
	Contract     *ERC20UTXO        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20UTXOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20UTXOCallerSession struct {
	Contract *ERC20UTXOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20UTXOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20UTXOTransactorSession struct {
	Contract     *ERC20UTXOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20UTXORaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20UTXORaw struct {
	Contract *ERC20UTXO // Generic contract binding to access the raw methods on
}

// ERC20UTXOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20UTXOCallerRaw struct {
	Contract *ERC20UTXOCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20UTXOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20UTXOTransactorRaw struct {
	Contract *ERC20UTXOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20UTXO creates a new instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXO(address common.Address, backend bind.ContractBackend) (*ERC20UTXO, error) {
	contract, err := bindERC20UTXO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXO{ERC20UTXOCaller: ERC20UTXOCaller{contract: contract}, ERC20UTXOTransactor: ERC20UTXOTransactor{contract: contract}, ERC20UTXOFilterer: ERC20UTXOFilterer{contract: contract}}, nil
}

// NewERC20UTXOCaller creates a new read-only instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOCaller(address common.Address, caller bind.ContractCaller) (*ERC20UTXOCaller, error) {
	contract, err := bindERC20UTXO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOCaller{contract: contract}, nil
}

// NewERC20UTXOTransactor creates a new write-only instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20UTXOTransactor, error) {
	contract, err := bindERC20UTXO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOTransactor{contract: contract}, nil
}

// NewERC20UTXOFilterer creates a new log filterer instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20UTXOFilterer, error) {
	contract, err := bindERC20UTXO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOFilterer{contract: contract}, nil
}

// bindERC20UTXO binds a generic wrapper to an already deployed contract.
func bindERC20UTXO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UTXOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20UTXO *ERC20UTXORaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20UTXO.Contract.ERC20UTXOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20UTXO *ERC20UTXORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.ERC20UTXOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20UTXO *ERC20UTXORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.ERC20UTXOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20UTXO *ERC20UTXOCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20UTXO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20UTXO *ERC20UTXOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20UTXO *ERC20UTXOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.contract.Transact(opts, method, params...)
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOCaller) GetUtxo(opts *bind.CallOpts, _id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ERC20UTXO.contract.Call(opts, out, "getUtxo", _id)
	return *ret0, *ret1, *ret2, err
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOSession) GetUtxo(_id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	return _ERC20UTXO.Contract.GetUtxo(&_ERC20UTXO.CallOpts, _id)
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOCallerSession) GetUtxo(_id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	return _ERC20UTXO.Contract.GetUtxo(&_ERC20UTXO.CallOpts, _id)
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOCaller) Utxos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	ret := new(struct {
		Owner     common.Address
		Value     *big.Int
		CreatedBy [32]byte
		Id        [32]byte
	})
	out := ret
	err := _ERC20UTXO.contract.Call(opts, out, "utxos", arg0)
	return *ret, err
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOSession) Utxos(arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	return _ERC20UTXO.Contract.Utxos(&_ERC20UTXO.CallOpts, arg0)
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOCallerSession) Utxos(arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	return _ERC20UTXO.Contract.Utxos(&_ERC20UTXO.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOTransactor) Create(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.contract.Transact(opts, "create", _to, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOSession) Create(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Create(&_ERC20UTXO.TransactOpts, _to, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOTransactorSession) Create(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Create(&_ERC20UTXO.TransactOpts, _to, _value)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOTransactor) Spend(opts *bind.TransactOpts, _id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.contract.Transact(opts, "spend", _id, _amount, _to)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOSession) Spend(_id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Spend(&_ERC20UTXO.TransactOpts, _id, _amount, _to)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOTransactorSession) Spend(_id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Spend(&_ERC20UTXO.TransactOpts, _id, _amount, _to)
}

// ERC20UTXOLogCreateIterator is returned from FilterLogCreate and is used to iterate over the raw logs and unpacked data for LogCreate events raised by the ERC20UTXO contract.
type ERC20UTXOLogCreateIterator struct {
	Event *ERC20UTXOLogCreate // Event containing the contract specifics and raw log

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
func (it *ERC20UTXOLogCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UTXOLogCreate)
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
		it.Event = new(ERC20UTXOLogCreate)
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
func (it *ERC20UTXOLogCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UTXOLogCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UTXOLogCreate represents a LogCreate event raised by the ERC20UTXO contract.
type ERC20UTXOLogCreate struct {
	Owner common.Address
	Id    [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogCreate is a free log retrieval operation binding the contract event 0xa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac63384.
//
// Solidity: e LogCreate(owner indexed address, id indexed bytes32, value uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) FilterLogCreate(opts *bind.FilterOpts, owner []common.Address, id [][32]byte) (*ERC20UTXOLogCreateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC20UTXO.contract.FilterLogs(opts, "LogCreate", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOLogCreateIterator{contract: _ERC20UTXO.contract, event: "LogCreate", logs: logs, sub: sub}, nil
}

// WatchLogCreate is a free log subscription operation binding the contract event 0xa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac63384.
//
// Solidity: e LogCreate(owner indexed address, id indexed bytes32, value uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) WatchLogCreate(opts *bind.WatchOpts, sink chan<- *ERC20UTXOLogCreate, owner []common.Address, id [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC20UTXO.contract.WatchLogs(opts, "LogCreate", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UTXOLogCreate)
				if err := _ERC20UTXO.contract.UnpackLog(event, "LogCreate", log); err != nil {
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

// ERC20UTXOLogSpendIterator is returned from FilterLogSpend and is used to iterate over the raw logs and unpacked data for LogSpend events raised by the ERC20UTXO contract.
type ERC20UTXOLogSpendIterator struct {
	Event *ERC20UTXOLogSpend // Event containing the contract specifics and raw log

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
func (it *ERC20UTXOLogSpendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UTXOLogSpend)
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
		it.Event = new(ERC20UTXOLogSpend)
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
func (it *ERC20UTXOLogSpendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UTXOLogSpendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UTXOLogSpend represents a LogSpend event raised by the ERC20UTXO contract.
type ERC20UTXOLogSpend struct {
	From     common.Address
	To       common.Address
	OldId    [32]byte
	NewId    [32]byte
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSpend is a free log retrieval operation binding the contract event 0xa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d18.
//
// Solidity: e LogSpend(from indexed address, to indexed address, oldId bytes32, newId bytes32, newValue uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) FilterLogSpend(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20UTXOLogSpendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20UTXO.contract.FilterLogs(opts, "LogSpend", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOLogSpendIterator{contract: _ERC20UTXO.contract, event: "LogSpend", logs: logs, sub: sub}, nil
}

// WatchLogSpend is a free log subscription operation binding the contract event 0xa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d18.
//
// Solidity: e LogSpend(from indexed address, to indexed address, oldId bytes32, newId bytes32, newValue uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) WatchLogSpend(opts *bind.WatchOpts, sink chan<- *ERC20UTXOLogSpend, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20UTXO.contract.WatchLogs(opts, "LogSpend", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UTXOLogSpend)
				if err := _ERC20UTXO.contract.UnpackLog(event, "LogSpend", log); err != nil {
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
