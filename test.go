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

func IsCanonical(v Variable) bool {
	switch v.(type) {
	case expr.LinearExpression, *expr.LinearExpression, expr.Term, *expr.Term:
		return true
	}
	if reflect.TypeOf(v).String() == "expr.Expression" {
		return true
	}
	return false