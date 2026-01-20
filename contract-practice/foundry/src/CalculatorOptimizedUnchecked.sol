// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CalculatorOptimizedUnchecked {
    uint256 public lastResult;

    function add(uint256 a, uint256 b) external {
        unchecked {
            lastResult = a + b;
        }
    }

    function sub(uint256 a, uint256 b) external {
        unchecked {
            lastResult = a - b;
        }
    }
}
