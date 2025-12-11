// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


contract Simple {
    uint256 public nextTokenId;
    constructor() {
        nextTokenId = 1;
    }

    function add() external {
        nextTokenId++;
    }

    function getCurrentTokenId() external view returns (uint256){
        return nextTokenId;
    }

}
