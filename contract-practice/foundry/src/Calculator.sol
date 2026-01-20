// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Calculator {
    uint256 public lastResult;

    function add(uint256 a, uint256 b) external returns (uint256) {
        uint256 result = a + b;
        lastResult = result;
        return result;
    }

    function sub(uint256 a, uint256 b) external returns (uint256) {
        require(a >= b, "underflow");
        uint256 result = a - b;
        lastResult = result;
        return result;
    }
}
