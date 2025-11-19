// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract reverse {
    function reverse_str(string calldata str) public pure returns (string memory)  {
        bytes memory b = bytes(str);
        uint len = b.length;
        for (uint i = 0; i < len / 2; i++) {
            bytes1 temp = b[i];
            b[i] = b[len - 1 - i];
            b[len - 1 - i] = temp;
        }

        return string(b);
    }
}
