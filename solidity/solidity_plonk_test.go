package solidity

import (
	"math/big"
	"os"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	plonk_bn254 "github.com/consensys/gnark/backend/plonk/bn254"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/examples/cubic"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type ExportSolidityTestSuitePlonk struct {
	suite.Suite

	// backend
	backend *backends.SimulatedBackend

	// verifier contract
	verifierContract *PlonkVerifier

	// plonk gnark objects
	vk      plonk.VerifyingKey
	pk      plonk.ProvingKey
	circuit cubic.Circuit
	scs     constraint.ConstraintSystem
}

func TestRunExportSolidityTestSuitePlonk(t *testing.T) {
	suite.Run(t, new(ExportSolidityTestSuitePlonk))
}

func (t *ExportSolidityTestSuitePlonk) SetupTest() {

	const gasLimit uint64 = 4712388

	// setup simulated backend
	key, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	t.NoError(err, "init keyed transactor")

	genesis := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(1000000000000000000)}, // 1 Eth
	}
	t.backend = backends.NewSimulatedBackend(genesis, gasLimit)

	// deploy verifier contract

	_, _, v, err := DeployPlonkVerifier(auth, t.backend)
	t.NoError(err, "deploy verifier contract failed")
	t.verifierContract = v
	t.backend.Commit()

	t.scs, err = frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &t.circuit)
	t.NoError(err, "compiling SCS failed")

	// read proving and verifying keys
	t.pk = plonk.NewProvingKey(ecc.BN254)
	{
		f, _ := os.Open("cubic.plonk.pk")
		_, err = t.pk.ReadFrom(f)
		f.Close()
		t.NoError(err, "reading proving key failed")
	}
	t.vk = plonk.NewVerifyingKey(ecc.BN254)
	{
		f, _ := os.Open("cubic.plonk.vk")
		_, err = t.vk.ReadFrom(f)
		f.Close()
		t.NoError(err, "reading verifying key failed")
	}

}

func (t *ExportSolidityTestSuitePlonk) TestVerifyProof() {

	// create a valid proof
	var assignment cubic.Circuit
	assignment.X = 3
	assignment.Y = 35

	// witness creation
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	t.NoError(err, "witness creation failed")

	// prove
	proof, err := plonk.Prove(t.scs, t.pk, witness)
	t.NoError(err, "proving failed")

	// ensure gnark (Go) code verifies it
	publicWitness, _ := witness.Public()
	err = plonk.Verify(proof, t.vk, publicWitness)
	t.NoError(err, "verifying failed")

	var publicInputs [1]*big.Int

	p := proof.(*plonk_bn254.Proof)
	serializedProof := p.MarshalSolidity()

	// public witness
	publicInputs[0] = new(big.Int).SetUint64(35)
	// call the contract
	res, err := t.verifierContract.Verify(&bind.CallOpts{}, serializedProof[:], publicInputs[:])
	if t.NoError(err, "calling verifier on chain gave error") {
		t.True(res, "calling verifier on chain didn't succeed")
	}

	// (wrong) public witness
	publicInputs[0] = new(big.Int).SetUint64(42)

	// call the contract should fail
	res, err = t.verifierContract.Verify(&bind.CallOpts{}, serializedProof[:], publicInputs[:])
	if t.NoError(err, "calling verifier on chain gave error") {
		t.False(res, "calling verifier on chain succeed, and shouldn't have")
	}
}
