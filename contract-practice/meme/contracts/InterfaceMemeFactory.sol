// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

contract InterfaceMemeFactory {
    function createPair(address tokenA, address tokenB) external returns (address pair);
}
