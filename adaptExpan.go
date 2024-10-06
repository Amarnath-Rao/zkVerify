package main

import (
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"

	"github.com/PolyhedraZK/ExpanderCompilerCollection"
	"github.com/PolyhedraZK/ExpanderCompilerCollection/test"
)

type VerifyHeaderCircuit struct {
	BlockNumber frontend.Variable
	BlockHash   frontend.Variable
	ParentHash  frontend.Variable
}

func (circuit *VerifyHeaderCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(circuit.BlockHash, circuit.ParentHash)
	return nil
}

func main() {
	// Define the circuit and assignment
	assignment := &VerifyHeaderCircuit{
		BlockNumber: 1,
		BlockHash:   0xabc,
		ParentHash:  0xabc,
	}

	// Compile the circuit with Expander
	circuit, _ := ExpanderCompilerCollection.Compile(ecc.BN254.ScalarField(), &VerifyHeaderCircuit{})
	c := circuit.GetLayeredCircuit()
	os.WriteFile("circuit.txt", c.Serialize(), 0o644)

	// Solve the input to generate witness
	inputSolver := circuit.GetInputSolver()
	witness, _ := inputSolver.SolveInputAuto(assignment)
	os.WriteFile("witness.txt", witness.Serialize(), 0o644)

	// Verify the circuit using Expander
	if !test.CheckCircuit(c, witness) {
		panic("verification failed")
	}
}
