package solidity

//go:generate go run contract/main.go
//go:generate solc --evm-version paris --combined-json abi,bin contract_plonk.sol -o abi --overwrite
//go:generate abigen --combined-json abi/combined.json --pkg solidity --out solidity_plonk.go
//go:generate solc --evm-version paris --combined-json abi,bin contract_g16.sol -o abi --overwrite
//go:generate abigen --combined-json abi/combined.json --pkg solidity --out solidity_groth16.go
