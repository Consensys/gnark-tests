// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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

// PairingMetaData contains all meta data concerning the Pairing contract.
var PairingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122030330cea3e1735c56cc98623a431006ee9fc9f0299fd9cd8295f740513946eee64736f6c63430008110033",
}

// PairingABI is the input ABI used to generate the binding from.
// Deprecated: Use PairingMetaData.ABI instead.
var PairingABI = PairingMetaData.ABI

// PairingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairingMetaData.Bin instead.
var PairingBin = PairingMetaData.Bin

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"43753b4d": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[1])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610ff4806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004361003e366004610e4e565b610057565b604051901515815260200160405180910390f35b6000610061610c9e565b60408051808201825287518152602080890151818301529083528151608081018352875151818401908152885183015160608301528152825180840184528883018051518252518301518184015281830152838201528151808301835286518152868201519181019190915290820152805151600080516020610f9f833981519152116101355760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d7100000000000000000060448201526064015b60405180910390fd5b805160200151600080516020610f9f833981519152116101975760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015260640161012c565b60208101515151600080516020610f9f833981519152116101fa5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015260640161012c565b602081810151015151600080516020610f9f8339815191521161025f5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015260640161012c565b602081810151510151600080516020610f9f833981519152116102c45760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015260640161012c565b6020818101518101510151600080516020610f9f8339815191521161032b5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015260640161012c565b604081015151600080516020610f9f8339815191521161038d5760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015260640161012c565b600080516020610f9f833981519152816040015160200151106103f25760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015260640161012c565b60005b6001811015610495577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000184826001811061043157610431610ef4565b6020020135106104835760405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015260640161012c565b8061048d81610f20565b9150506103f5565b5060006104a06105c6565b60408051808201909152600080825260208201529091506104bf610cef565b6104c7610d0d565b60408051808201825260008082526020808301919091527f1e62325568a43f02acf44dd3b4c021e839744c51eac52d822627ccc9a34291bc86527f1d1b0de1a6cab2189f1ccfbaf42805ed91e6476592978ba983aebfdcb96b3363868201527f1e0279e364797809e961d05a4edbce4347feb9c21cdbc4d8d28834b1bbac7d8f84527f25b747263e511dd1f3ebd29a1db05f2b875c385b5ecc132ce5404b5d4f918672908401528835918301919091526105838282858761083b565b6105b76105938760000151610872565b876020015187600001518860200151888a604001518c604001518c60600151610908565b9b9a5050505050505050505050565b6105ce610d2b565b6040805180820182527f23b8ebf14d72bd310341a1a08c2c4a27204543d8cc0ed4d45e0d6d721c3805a981527f15d0c42745f1d316ee28489e02877db5cbd9aefe22b0aeffdead5f676d45bde36020808301919091529083528151608080820184527f08f1643e6fbb6b00323bb87ead596be9530bdf1f8cbd8f8c48812bc8157bd0ee8285019081527f1c1810a659aae4bed0b715c41daff116c4ada7846dc91a5f0f14749687385434606080850191909152908352845180860186527f2b09bfc265010a2d147f6733918c5f6aa91b75e03e9b5b01806b5fd4c5eb6ef881527f259cc17d10edf863b2bbb3b3a8805ebeb3228846f6a3192955df38844c128303818601528385015285840192909252835180820185527f05998bde252994b348d1f9287be2672c7de1609b8dfb621b610b2c0d17e7bc7a8186019081527f18ee8eaa0acec7063f793664fd7bb32317196e71ab85216fa2faffc222e928ea828501528152845180860186527eb64aeadea89bf7f15d9047a885f70f9e78818f29f9d6706bd5654104cfce1c81527f27837e88becf28f174df308ee34b908df014bd144971238790c9398e6d284229818601528185015285850152835190810184527f1e84ed2c4ddeb367cd29a3775b980e8afc6a1426214188f1c77a322add0a38138185019081527f23ccc9ea64d4e97a589448b009629fc9269affc9001ed8b283bdf89969d6af6182840152815283518085019094527f23c4fcf8a53f1d33f9a961e3d0a030b717c011e8c5eafec3c63ab85c333fd72984527e092558aecca3774ec04ec196055a59547bc340d691d27df8d6a482de9afa9284840152918201929092529082015290565b6108458484610bdf565b805182526020808201518184015283516040840152830151606083015261086c8282610c41565b50505050565b6040805180820190915260008082526020820152815115801561089757506020820151155b156108b5575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020610f9f83398151915284602001516108e89190610f39565b61090090600080516020610f9f833981519152610f5b565b905292915050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015610b5c576000610987826006610f74565b905085826004811061099b5761099b610ef4565b602002015151836109ad836000610f8b565b815181106109bd576109bd610ef4565b6020026020010181815250508582600481106109db576109db610ef4565b602002015160200151838260016109f29190610f8b565b81518110610a0257610a02610ef4565b602002602001018181525050848260048110610a2057610a20610ef4565b6020020151515183610a33836002610f8b565b81518110610a4357610a43610ef4565b602002602001018181525050848260048110610a6157610a61610ef4565b6020020151516001602002015183610a7a836003610f8b565b81518110610a8a57610a8a610ef4565b602002602001018181525050848260048110610aa857610aa8610ef4565b602002015160200151600060028110610ac357610ac3610ef4565b602002015183610ad4836004610f8b565b81518110610ae457610ae4610ef4565b602002602001018181525050848260048110610b0257610b02610ef4565b602002015160200151600160028110610b1d57610b1d610ef4565b602002015183610b2e836005610f8b565b81518110610b3e57610b3e610ef4565b60209081029190910101525080610b5481610f20565b915050610971565b50610b65610d6f565b6000602082602086026020860160086107d05a03fa90508080610b8457fe5b5080610bca5760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015260640161012c565b505115159d9c50505050505050505050505050565b600060608260808560076107d05a03fa90508080610bf957fe5b5080610c3c5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015260640161012c565b505050565b600060608260c08560066107d05a03fa90508080610c5b57fe5b5080610c3c5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015260640161012c565b6040805160a081019091526000606082018181526080830191909152815260208101610cc8610d8d565b8152602001610cea604051806040016040528060008152602001600081525090565b905290565b60405180608001604052806004906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b6040805160c0810190915260006080820181815260a0830191909152815260208101610d55610d8d565b8152602001610d62610d8d565b8152602001610cea610d8d565b60405180602001604052806001906020820280368337509192915050565b6040518060400160405280610da0610da9565b8152602001610cea5b60405180604001604052806002906020820280368337509192915050565b6040805190810167ffffffffffffffff81118282101715610df857634e487b7160e01b600052604160045260246000fd5b60405290565b600082601f830112610e0f57600080fd5b610e17610dc7565b806040840185811115610e2957600080fd5b845b81811015610e43578035845260209384019301610e2b565b509095945050505050565b600080600080610120808688031215610e6657600080fd5b610e708787610dfe565b9450604087605f880112610e8357600080fd5b610e8b610dc7565b8060c089018a811115610e9d57600080fd5b838a015b81811015610ec257610eb38c82610dfe565b84526020909301928401610ea1565b50819750610ed08b82610dfe565b965050505050868187011115610ee557600080fd5b50929591945092610100019150565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610f3257610f32610f0a565b5060010190565b600082610f5657634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610f6e57610f6e610f0a565b92915050565b8082028115828204841417610f6e57610f6e610f0a565b80820180821115610f6e57610f6e610f0a56fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220f8cd667ce0fd9304fe03b182f43d9d63c23cc246d3d1b8fb7cbc5720b4ab60e664736f6c63430008110033",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// Deprecated: Use VerifierMetaData.Sigs instead.
// VerifierFuncSigs maps the 4-byte function signature to its string representation.
var VerifierFuncSigs = VerifierMetaData.Sigs

// VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierMetaData.Bin instead.
var VerifierBin = VerifierMetaData.Bin

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
