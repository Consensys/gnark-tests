# gnark-tests

This repo contains tests (interop or integration) that may drag some extra dependencies, for the following projects:

* [`gnark`: a framework to execute (and verify) algorithms in zero-knowledge](https://github.com/consensys/gnark) 
* [`gnark-crypto`](https://github.com/consensys/gnark-crypto)


Note that since the verifying key of the contract is included in the `solidity/contract.sol`, changes to gnark version or circuit should result in running `go run contract/main.go && abigen --sol contract.sol --pkg solidity --out solidity.go`  to regenerate keys and solidity  contract.