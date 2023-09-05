# gnark-tests

**UNUSED / archived**, superseded by https://github.com/ConsenSys/gnark-solidity-checker and CI tests in gnark directly. 


This repo contains tests (interop or integration) that may drag some extra dependencies, for the following projects:

* [`gnark`: a framework to execute (and verify) algorithms in zero-knowledge](https://github.com/consensys/gnark) 
* [`gnark-crypto`](https://github.com/consensys/gnark-crypto)

## Solidity verifier (groth16 and plonk)

```bash
cd solidity
go generate
go test
```
or
```bash
make
```

Note that since the verifying key of the contract is included in the `solidity/contract.sol`, changes to gnark version or circuit should result in running `go generate`  to regenerate keys and solidity contracts.

It needs `solc` and `abigen` (1.10.17-stable).
