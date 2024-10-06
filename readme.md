# zkVerify for BSC

## Project Overview

**zkVerify for BSC** is a zero-knowledge proof (ZKP) implementation designed to verify Binance Smart Chain (BSC) block headers. By utilizing the Expander Compiler Collection and the Gnark library, this project aims to ensure the correctness and integrity of BSC block headers in a decentralized and privacy-preserving manner. This solution aligns with the goals of the Polyhedra Challenge, enhancing blockchain verification processes with cutting-edge ZKP technology.

## Features

- **Zero-Knowledge Proof Implementation**: Leverages the Gnark library to prove the correctness of BSC block headers.
- **Expander Compiler Collection**: Compiles the ZKP circuits for efficient execution and verification.
- **Scalability and Security**: Ensures high-performance verification while maintaining strong security guarantees.
- **Open Source**: Source code is publicly available for community collaboration and improvement.

![image](https://github.com/user-attachments/assets/2880b831-2a5b-4734-8ded-9457fe058d96)
![image](https://github.com/user-attachments/assets/3a1aa7a2-5f3b-41a1-a6b0-e8e6799b9139)

## Theoretical Foundation

### Zero-Knowledge Proofs (ZKP)

Zero-knowledge proofs are cryptographic methods that allow one party (the prover) to prove to another party (the verifier) that they know a value without revealing the value itself. This is particularly useful in blockchain and privacy-preserving technologies, where proving the correctness of data without revealing sensitive information is essential.

### Binance Smart Chain (BSC) Block Header Verification

In the context of blockchain, a block header is a crucial component that contains metadata about a block, including its previous hash, timestamp, and Merkle root. Verifying the correctness of a block header ensures that the block is valid and that the blockchain's integrity is maintained.

### zkVerify for BSC

The **zkVerify for BSC** project applies zero-knowledge proofs to verify the correctness of BSC block headers. The process involves the following steps:

1. **Circuit Implementation in Gnark**: The verification logic for the BSC block header is implemented as a circuit using the [Gnark library](https://github.com/consensys/gnark). This circuit checks the validity of the block header according to the BSC consensus rules.

2. **Compilation with Expander Compiler Collection**: The implemented circuit is compiled using the [Expander Compiler Collection](https://github.com/PolyhedraZK/ExpanderCompilerCollection). This step transforms the circuit into a format that can be efficiently executed and verified.

3. **Zero-Knowledge Proof Generation**: After compiling the circuit, the Expander Compiler Collection generates a zero-knowledge proof that attests to the correctness of the block header. The proof is generated without revealing any sensitive information, ensuring privacy and security.

4. **Verification of the Proof**: The generated proof is then verified using the verifier module provided by the Expander framework. If the proof is valid, it confirms that the BSC block header is correct and follows the BSC consensus rules.

### Advantages

- **Scalability**: The use of ZKP allows for efficient verification of block headers without the need for heavy computation or large data transfers.
- **Privacy**: Sensitive information about the block header is never exposed, preserving the privacy of the blockchain's transactions.
- **Security**: ZKP provides a strong guarantee that the block header is correct, reducing the risk of attacks or fraud.

By leveraging these technologies, "zkVerify for BSC" enhances the security and integrity of BSC block header verification, making it a robust solution for blockchain applications.



## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Gnark](https://github.com/consensys/gnark)
- [Expander Compiler Collection](https://github.com/PolyhedraZK/ExpanderCompilerCollection)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/arkakxqi/zkVerify-for-BSC.git
   cd zkVerify-for-BSC
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Compile the Circuit**:
   - Modify the circuit to fit your specific use case within the BSC block header verification.
   - Compile using the Expander Compiler:
     ```bash
     go run main.go
     ```

4. **Run the Example**:
   ```bash
   go run zkVerify.go
   expander compile --circuit=my_circuit --output=my_circuit.compiled
   ```
![image](https://github.com/user-attachments/assets/1b802406-e9c6-4243-93ec-ba5702b00d6d)
-> BNB Chain：0xE014fe8c4d5C23EDB7AC4011F226e869ac7Ef5CC 
-> V2 BNB Chain: 0x8ddF05F9A5c488b4973897E278B58895bF87Cb24

```
function setOracle(uint16 dstChainId, address zkLightClient) external {

    uint256 TYPE_ORACLE = 6;

    ILayerZeroEndpoint(lzEndpointAddress).setConfig(

        ILayerZeroEndpoint(lzEndpointAddress).getSendVersion(address(this)),

        dstChainId,

        TYPE_ORACLE,

        abi.encode(zkLightClient)

    );

}
```

## Project Structure

- **/circuit**: Contains the ZKP circuit implementation.
- **/docs**: Documentation files, including technical details and architecture.
- **/examples**: Example usage of the zkVerify for BSC.
- **/test**: Test cases to ensure the correctness of the implementation.

## How It Works

1. **Circuit Definition**: The BSC block header verification logic is implemented in a Gnark circuit.
2. **Input Solver**: The circuit's input is solved using the Expander Compiler Collection's InputSolver.
3. **Proof Generation**: The solved input is used to generate a proof, which is then verified to ensure the correctness of the block header.

## Contributing

We welcome contributions from the community. If you'd like to contribute, please follow the standard [GitHub Fork & Pull Request workflow](https://guides.github.com/activities/forking/).

1. Fork the Project.
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`).
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the Branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.



# What is ExpanderCompilerCollection

Expander is a proof generation backend for the Polyhedra Network. The ExpanderCompilerCollection is a component of the Expander proof system. It transforms circuits written in [gnark](https://github.com/Consensys/gnark) into an intermediate representation (IR) of a layered circuit. This IR can later be used by the [Expander prover](https://github.com/PolyhedraZK/Expander) to generate proofs.


## Code Overview

### 1. **Understand the `VerifyHeader` Function**

The `VerifyHeader` function in the provided code is part of the Ethereum consensus engine interface. Its purpose is to verify if a block header adheres to the consensus rules defined by a specific consensus algorithm. To use zero-knowledge proofs (ZKPs) to validate this, you need to:

1. **Understand the Function Logic:** The `VerifyHeader` function performs checks based on consensus rules. You'll need to translate these rules into a zero-knowledge proof that can be validated without revealing the actual block header details.

2. **Implement a Gnark Circuit:** Develop a Gnark circuit that represents the logic of the `VerifyHeader` function. This circuit will generate a proof to confirm that the header verification was done correctly.

### 2. **Implement Zero-Knowledge Proofs with Gnark**

**Setup Gnark:**
   - Install Gnark and its dependencies if you haven’t already.
   - Create a Gnark circuit that models the logic of `VerifyHeader`. 

**Example Circuit (Simplified):**

```go
package main

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark-crypto/ecc"
)

type VerifyHeaderCircuit struct {
	BlockNumber frontend.Variable
	BlockHash   frontend.Variable
	ParentHash  frontend.Variable
}

func (circuit *VerifyHeaderCircuit) Define(api frontend.API) error {
	// Example constraints, replace with actual verification logic
	api.AssertIsEqual(circuit.BlockHash, circuit.ParentHash)
	return nil
}

func main() {
	// Define assignment and compile the circuit
	assignment := &VerifyHeaderCircuit{
		BlockNumber: 1,
		BlockHash:   0xabc,
		ParentHash:  0xabc,
	}

	r1cs, _ := frontend.Compile(ecc.BN254.ScalarField(), frontend.NewBuilder, &VerifyHeaderCircuit{})
	witness, _ := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()
	pk, vk, _ := groth16.Setup(r1cs)
	proof, _ := groth16.Prove(r1cs, pk, witness)
	err := groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		panic(err)
	}
}
```

### 3. **Adapt Circuit for Expander**

Replace the Gnark-specific code with code for Expander, particularly for compiling and verifying the circuit.

#### **Modified Circuit for Expander**

```go
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
```

### 4. **Generate and Verify Proof**

Once you have `circuit.txt` and `witness.txt`, use Expander's tools to generate and verify the proof.

- **Generate Proof:** Use Expander’s command-line tools or API to generate the proof using your `circuit.txt` and `witness.txt`.
- **Verify Proof:** Verify the proof using Expander’s verification tools.

### 5. **Develop a Smart Contract for Verification**

Develop a Solidity smart contract that verifies the zero-knowledge proof generated by your Gnark circuit. The contract should accept the proof and the block header and check if the proof is valid.

```solidity
pragma solidity ^0.8.0;

interface IZkProofVerifier {
    function verifyProof(bytes calldata proof, bytes calldata input) external returns (bool);
}

contract VerifyBlockHeader {
    IZkProofVerifier public verifier;

    constructor(address _verifier) {
        verifier = IZkProofVerifier(_verifier);
    }

    function verifyHeader(bytes calldata proof, bytes calldata input) external returns (bool) {
        return verifier.verifyProof(proof, input);
    }
}
```

### 6. **Testing and Validation**

Test your Gnark circuit and Solidity contract to ensure they work together correctly. Validate that your proof accurately represents the correct execution of `VerifyHeader`.

**Example Test:**

```javascript
const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("VerifyBlockHeader", function () {
    it("should verify block header proof correctly", async function () {
        const [deployer] = await ethers.getSigners();
        const Verifier = await ethers.getContractFactory("VerifyBlockHeader");
        const verifier = await Verifier.deploy("0xVerifierAddress");

        const proof = "0x..."; // Your proof
        const input = "0x..."; // Your input data

        expect(await verifier.verifyHeader(proof, input)).to.be.true;
    });
});
```
