// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";

// 原始版本
import "../src/Calculator.sol";

// 结构优化版
import "../src/CalculatorOptimizedStruct.sol";

// unchecked 优化版
import "../src/CalculatorOptimizedUnchecked.sol";

contract CalculatorGasTest is Test {
    Calculator original;
    CalculatorOptimizedStruct optimizedStruct;
    CalculatorOptimizedUnchecked optimizedUnchecked;

    function setUp() public {
        original = new Calculator();
        optimizedStruct = new CalculatorOptimizedStruct();
        optimizedUnchecked = new CalculatorOptimizedUnchecked();
    }

    //原始合约 Gas
    function testGas_Original_Add() public {
        original.add(10, 5);
    }

    function testGas_Original_Sub() public {
        original.sub(10, 5);
    }

    //结构优化版 Gas
    function testGas_Struct_Add() public {
        optimizedStruct.add(10, 5);
    }

    function testGas_Struct_Sub() public {
        optimizedStruct.sub(10, 5);
    }

    //unchecked 优化版 Gas
    function testGas_Unchecked_Add() public {
        optimizedUnchecked.add(10, 5);
    }

    function testGas_Unchecked_Sub() public {
        optimizedUnchecked.sub(10, 5);
    }
}
