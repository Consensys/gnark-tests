package neogo

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/nspcc-dev/neo-go/pkg/neotest"
	"github.com/nspcc-dev/neo-go/pkg/neotest/chain"
	"github.com/nspcc-dev/neo-go/pkg/util/slice"
	"github.com/stretchr/testify/require"
)

// CubicCircuit defines a simple circuit
// x**3 + x + 5 == y
// that checks that the prover knows the solution of the provided expression.
// The circuit must declare its public and secret inputs as frontend.Variable.
// At compile time, frontend.Compile(...) recursively parses the struct fields
// that contains frontend.Variable to build the frontend.constraintSystem.
// By default, a frontend.Variable has the gnark:",secret" visibility.
type CubicCircuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x,secret"` // Secret input
	Y frontend.Variable `gnark:"y,public"` // Public input
}

// A gnark circuit must implement the frontend.Circuit interface
// (https://docs.gnark.consensys.net/HowTo/write/circuit_structure).
var _ = frontend.Circuit(&CubicCircuit{})

// Define declares the circuit constraints
// x**3 + x + 5 == y.
func (circuit *CubicCircuit) Define(api frontend.API) error {
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)

	// Can be used for the circuit debugging.
	api.Println("X^3", x3)

	api.AssertIsEqual(circuit.Y, api.Add(x3, circuit.X, 5))
	return nil
}

// TestExportNeoGo_EndToEnd shows how to generate proof for pre-defined cubic circuit,
// how to generate Go verification contract that can be compiled by NeoGo and deployed
// to the chain and how to verify proofs via verification contract invocation.
func TestExportNeoGo_EndToEnd(t *testing.T) {
	var (
		circuit    CubicCircuit
		assignment = CubicCircuit{X: 3, Y: 35}
	)

	// Compile our circuit into a R1CS (a constraint system).
	ccs, err := frontend.Compile(ecc.BLS12_381.ScalarField(), r1cs.NewBuilder, &circuit)
	require.NoError(t, err)

	// One time setup (groth16 zkSNARK).
	pk, vk, err := groth16.Setup(ccs)
	require.NoError(t, err)

	// Intermediate step: witness definition.
	witness, err := frontend.NewWitness(&assignment, ecc.BLS12_381.ScalarField())
	require.NoError(t, err)
	publicWitness, err := witness.Public()
	require.NoError(t, err)

	// Proof creation (groth16).
	proof, err := groth16.Prove(ccs, pk, witness)
	require.NoError(t, err)

	// Ensure that gnark can successfully verify the proof (just in case).
	err = groth16.Verify(proof, vk, publicWitness)
	require.NoError(t, err)

	// Now, when we're sure that the proof is valid, we can create and deploy verification
	// contract to the Neo testing chain.

	// Get the proof bytes (points are in the compressed form, as Verification contract accepts it).
	proofSizeCompressed := int64(bls12381.SizeOfG1AffineCompressed + bls12381.SizeOfG2AffineCompressed + bls12381.SizeOfG1AffineCompressed)
	var buf bytes.Buffer
	n, err := proof.WriteTo(&buf)
	require.NoError(t, err)
	require.Equal(t, proofSizeCompressed, n)
	proofBytes := slice.Copy(buf.Bytes())

	aBytes := proofBytes[:bls12381.SizeOfG1AffineCompressed]
	bBytes := proofBytes[bls12381.SizeOfG1AffineCompressed : bls12381.SizeOfG1AffineCompressed+bls12381.SizeOfG2AffineCompressed]
	cBytes := proofBytes[bls12381.SizeOfG1AffineCompressed+bls12381.SizeOfG2AffineCompressed:]

	publicWitnessBytes, err := publicWitness.MarshalBinary()
	require.NoError(t, err)
	numPublicWitness := binary.BigEndian.Uint32(publicWitnessBytes[:4])
	numSecretWitness := binary.BigEndian.Uint32(publicWitnessBytes[4:8])
	numVectorElements := binary.BigEndian.Uint32(publicWitnessBytes[8:12])

	// Ensure that serialization format is as expected (just in case).
	require.Equal(t, uint32(0), numSecretWitness)
	require.Equal(t, numPublicWitness+numSecretWitness, numVectorElements)

	// Create public witness input.
	input := make([]any, numVectorElements)
	offset := 12
	for i := range input { // firstly - public witnesses, after that - private ones (but they are missing from publicWitness anyway).
		start := offset + i*fr.Bytes
		end := start + fr.Bytes
		slice.Reverse(publicWitnessBytes[start:end]) // gsnark stores witnesses in the BE form, but native CryptoLib accepts LE-encoded fields elements (not a canonical form).
		input[i] = publicWitnessBytes[start:end]
	}

	// Generate verification contract.
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "verify.go")
	f, err := os.Create(srcPath)
	require.NoError(t, err)
	vk.ExportNeoGo(f)
	f.Close()

	// Create contract configuration file.
	cfgPath := filepath.Join(tmpDir, "verify.yml")
	f, err = os.Create(cfgPath)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(cfgPath, []byte(verifyCfg), os.ModePerm))

	// Create go.mod and go.sum for the verification contract.
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(verifyGomod), os.ModePerm))
	require.NoError(t, os.WriteFile(filepath.Join(tmpDir, "go.sum"), []byte(verifyGosum), os.ModePerm))

	// Create testing chain and deploy contract onto it.
	bc, committee := chain.NewSingle(t)
	e := neotest.NewExecutor(t, bc, committee, committee)

	// Compile verification contract and deploy the contract onto chain.
	c := neotest.CompileFile(t, e.Validator.ScriptHash(), srcPath, cfgPath)
	e.DeployContract(t, c, nil)

	// Verify proof via verification contract call.
	validatorInvoker := e.ValidatorInvoker(c.Hash)
	validatorInvoker.Invoke(t, true, "verifyProof", aBytes, bBytes, cBytes, input)
}

const (
	// verifyCfg is a contract configuration file required to compile smart
	// contract.
	verifyCfg = `name: "Groth16 prooving contract example"
sourceurl: https://github.com/nspcc-dev/neo-go/
supportedstandards: []`

	// verifyGomod is a standard go.mod file containing module name, go version
	// and dependency packages version needed for smart contract compilation.
	verifyGomod = `module verify

go 1.19

require github.com/nspcc-dev/neo-go/pkg/interop v0.0.0-20230606150208-a2daad6ba614
`

	// verifyGosum is a standard go.sum file needed for contract compilation.
	verifyGosum = `github.com/nspcc-dev/neo-go/pkg/interop v0.0.0-20230606150208-a2daad6ba614 h1:MiDBj73HNgPUbJRpXWLXrsGvX4rkYVDrdSmfOwivGR8=
github.com/nspcc-dev/neo-go/pkg/interop v0.0.0-20230606150208-a2daad6ba614/go.mod h1:ZUuXOkdtHZgaC13za/zMgXfQFncZ0jLzfQTe+OsDOtg=
`
)
