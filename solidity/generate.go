package solidity

//go:generate go run contract/main.go
//go:generate abigen --sol contract_g16.sol --pkg solidity --out solidity_groth16.go
//go:generate abigen --sol contract_plonk.sol --pkg solidity --out solidity_plonk.go
