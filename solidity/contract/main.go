package main

import (
	"log"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/examples/cubic"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test"
)

func main() {
	err := generateGroth16()
	if err != nil {
		log.Fatal("groth16 error:", err)
	}

	err = generatePlonk()
	if err != nil {
		log.Fatal("plonk error:", err)
	}
}

func generateGroth16() error {
	var circuit cubic.Circuit

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		return err
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return err
	}
	{
		f, err := os.Create("cubic.g16.vk")
		if err != nil {
			return err
		}
		_, err = vk.WriteRawTo(f)
		if err != nil {
			return err
		}
	}
	{
		f, err := os.Create("cubic.g16.pk")
		if err != nil {
			return err
		}
		_, err = pk.WriteRawTo(f)
		if err != nil {
			return err
		}
	}

	{
		f, err := os.Create("contract_g16.sol")
		if err != nil {
			return err
		}
		err = vk.ExportSolidity(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func generatePlonk() error {
	var circuit cubic.Circuit

	scs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &circuit)
	if err != nil {
		return err
	}

	srs, err := test.NewKZGSRS(scs)
	if err != nil {
		return err
	}

	pk, vk, err := plonk.Setup(scs, srs)
	if err != nil {
		return err
	}
	{
		f, err := os.Create("cubic.plonk.vk")
		if err != nil {
			return err
		}
		_, err = vk.WriteTo(f)
		if err != nil {
			return err
		}
	}
	{
		f, err := os.Create("cubic.plonk.pk")
		if err != nil {
			return err
		}
		_, err = pk.WriteTo(f)
		if err != nil {
			return err
		}
	}

	{
		f, err := os.Create("contract_plonk.sol")
		if err != nil {
			return err
		}
		err = vk.ExportSolidity(f)
		if err != nil {
			return err
		}
	}
	return nil
}
