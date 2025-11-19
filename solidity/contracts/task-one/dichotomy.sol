// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract dichotomy {
    constructor(){

    }

    function binarySearch(uint[] memory array, uint target) public view returns (int){
        uint left = 0;
        uint right = array.length;

        while (left < right) {
            uint mid = left + (right - left) / 2;
            if (array[mid] == target) {
                return int(mid);
            } else if (array[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return -1;
    }
}
