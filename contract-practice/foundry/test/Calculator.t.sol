// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/Calculator.sol";

contract CalculatorTest is Test {
    Calculator calc;

    function setUp() public {
        calc = new Calculator();
    }

    function testAddGas() public {
        calc.add(10, 20);
    }

    function testSubGas() public {
        calc.sub(20, 10);
    }
}
