// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "hardhat/console.sol";

contract roman {
    mapping(bytes1 => uint) private romanMap;
    mapping(uint => string) private integerMap;
    constructor(){
        romanMap["I"] = 1;
        romanMap["V"] = 5;
        romanMap["X"] = 10;
        romanMap["L"] = 50;
        romanMap["C"] = 100;
        romanMap["D"] = 500;
        romanMap["M"] = 1000;

        integerMap[1] = "I";
        integerMap[5] = "V";
        integerMap[10] = "X";
        integerMap[50] = "L";
        integerMap[100] = "C";
        integerMap[500] = "D";
        integerMap[1000] = "M";
    }


    function roman_to_integer(string calldata r) external view returns (uint){
        uint total = 0;
        uint prev = 0;
        bytes memory rb = bytes(r);

        for (uint i = rb.length; i > 0; i--) {
            bytes1 key = rb[i - 1];
            uint value = romanMap[key];

            if (value < prev) {
                total -= value;
            } else {
                total += value;
            }

            prev = value;
        }

        console.log("total =", total);

        return total;
    }

    function integer_to_roman(uint num) external view returns (string memory){
        string memory result;
        uint thousands = num / 1000;
        uint hundreds = (num % 1000) / 100;
        uint tens = (num % 100) / 10;
        uint ones = num % 10;

        for (uint i = 0; i < thousands; i++) {
            result = string.concat(result, integerMap[1000]);
        }

        result = string.concat(
            result,
            digitToRoman(hundreds, 100),
            digitToRoman(tens, 10),
            digitToRoman(ones, 1)
        );

        return result;
    }

    function digitToRoman(uint digit, uint u) internal view returns (string memory) {

        if (digit == 0) return "";

        if (digit <= 3) {
            return repeat(integerMap[u], digit);
        }
        else if (digit == 4) {
            return string.concat(integerMap[u], integerMap[u * 5]);
        }
        else if (digit <= 8) {
            return string.concat(
                integerMap[u * 5],
                repeat(integerMap[u], digit - 5)
            );
        }
        else if (digit == 9) {
            return string.concat(integerMap[u], integerMap[u * 10]);
        }

        revert("invalid digit");
    }

    function repeat(string memory s, uint count) internal pure returns (string memory) {
        string memory out = "";
        for (uint i = 0; i < count; i++) {
            out = string.concat(out, s);
        }
        return out;
    }
}
