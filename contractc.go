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
