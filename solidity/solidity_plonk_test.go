package solidity

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/kzg"
	"github.com/consensys/gnark/backend/plonk"
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
	verifierContract *KeyedPlonkVerifier

	// plonk gnark objects
	srs     kzg.SRS
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
	_, _, v, err := DeployKeyedPlonkVerifier(auth, t.backend)
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

	t.srs = kzg.NewSRS(ecc.BN254)
	{
		f, _ := os.Open("kzg.plonk.srs")
		_, err = t.srs.ReadFrom(f)
		f.Close()
		t.NoError(err, "reading kzg srs failed")
	}

	t.vk.InitKZG(t.srs)
	t.pk.InitKZG(t.srs)

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

	formattedProof, err := NewSolidityProof(proof)
	t.NoError(err, "formatting plonk proof")

	serializedProof := formattedProof.toArray()

	// public witness
	publicInputs[0] = new(big.Int).SetUint64(35)
	// printTx(input[:], serializedProof)
	// call the contract
	res, err := t.verifierContract.VerifySerializedProof(nil, publicInputs[:], serializedProof[:])
	if t.NoError(err, "calling verifier on chain gave error") {
		t.True(res, "calling verifier on chain didn't succeed")
	}

	// (wrong) public witness
	publicInputs[0] = new(big.Int).SetUint64(42)

	// call the contract should fail
	res, err = t.verifierContract.VerifySerializedProof(nil, publicInputs[:], serializedProof[:])
	if t.NoError(err, "calling verifier on chain gave error") {
		t.False(res, "calling verifier on chain succeed, and shouldn't have")
	}
}

// print inputs to debug through remix IDE or JS.
func printTx(input, proof []*big.Int) {
	fmt.Println("-----------")
	fmt.Print("[", "0x"+input[0].Text(16), "],[")
	for i := 0; i < len(proof); i++ {
		fmt.Print("0x" + proof[i].Text(16))
		if i == len(proof)-1 {
			fmt.Println("]")
		} else {
			fmt.Print(",")
		}
	}
	fmt.Println("-----------")
}

// derived from https://github.com/lightning-li/plonk_verifier_example/blob/main/proof_generation/src/main.go#L57
type solidityProof struct {
	WireCommitments               [3][2]*big.Int
	GrandProductCommitment        [2]*big.Int
	QuotientPolyCommitments       [3][2]*big.Int
	WireValuesAtZeta              [3]*big.Int
	GrandProductAtZetaOmega       *big.Int
	QuotientPolynomialAtZeta      *big.Int
	LinearizationPolynomialAtZeta *big.Int
	PermutationPolynomialsAtZeta  [2]*big.Int
	OpeningAtZetaProof            [2]*big.Int
	OpeningAtZetaOmegaProof       [2]*big.Int
}

func (p *solidityProof) toArray() []*big.Int {
	var r []*big.Int
	for i := 0; i < 3; i++ {
		r = append(r, p.WireCommitments[i][:]...)
	}
	r = append(r, p.GrandProductCommitment[:]...)
	for i := 0; i < 3; i++ {
		r = append(r, p.QuotientPolyCommitments[i][:]...)
	}
	r = append(r, p.WireValuesAtZeta[:]...)
	r = append(r, p.GrandProductAtZetaOmega)
	r = append(r, p.QuotientPolynomialAtZeta)
	r = append(r, p.LinearizationPolynomialAtZeta)
	r = append(r, p.PermutationPolynomialsAtZeta[:]...)
	r = append(r, p.OpeningAtZetaProof[:]...)
	r = append(r, p.OpeningAtZetaOmegaProof[:]...)
	return r
}

func NewSolidityProof(oProof plonk.Proof) (proof *solidityProof, err error) {
	proof = new(solidityProof)
	const fpSize = 32
	var buf bytes.Buffer
	_, err = oProof.WriteRawTo(&buf)

	if err != nil {
		return nil, err
	}
	proofBytes := buf.Bytes()
	index := 0
	for i := 0; i < 3; i++ {
		proof.WireCommitments[i][0] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
		index++
		proof.WireCommitments[i][1] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
		index++
	}

	proof.GrandProductCommitment[0] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
	index++
	proof.GrandProductCommitment[1] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
	index++

	for i := 0; i < 3; i++ {
		proof.QuotientPolyCommitments[i][0] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
		index++
		proof.QuotientPolyCommitments[i][1] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
		index++
	}

	proof.OpeningAtZetaProof[0] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
	index++
	proof.OpeningAtZetaProof[1] = new(big.Int).SetBytes(proofBytes[fpSize*index : fpSize*(index+1)])
	index++

	// plonk proof write len(ClaimedValues) which is 4 bytes
	offset := 4
	proof.QuotientPolynomialAtZeta = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
	index++

	proof.LinearizationPolynomialAtZeta = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
	index++

	for i := 0; i < 3; i++ {
		proof.WireValuesAtZeta[i] = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
		index++
	}

	for i := 0; i < 2; i++ {
		proof.PermutationPolynomialsAtZeta[i] = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
		index++
	}

	proof.OpeningAtZetaOmegaProof[0] = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
	index++
	proof.OpeningAtZetaOmegaProof[1] = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
	index++

	proof.GrandProductAtZetaOmega = new(big.Int).SetBytes(proofBytes[offset+fpSize*index : offset+fpSize*(index+1)])
	return proof, nil
}
