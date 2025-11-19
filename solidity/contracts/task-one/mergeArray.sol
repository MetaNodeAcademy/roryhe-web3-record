// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract mergeArray {
    constructor(){

    }

    function merge(uint[] memory array1, uint[] memory array2) public pure returns (uint[] memory){
        uint alen = array1.length;
        uint blen = array2.length;
        uint total = array1.length + array2.length;
        uint[] memory array = new uint[](total);

        //执行合并逻辑
        uint c = 0;
        for (uint i = 0; i < alen; i++) {
            array[c++] = array1[i];
        }
        for (uint j = 0; j < blen; j++) {
            array[c++] = array2[j];
        }

        //保持有序
        sort(array, 0, int(array.length - 1));

        return array;
    }

    function sort(uint[] memory array, int left, int right) private pure {
        if (left >= right) return;

        int i = left;
        int j = right;
        uint pivot = array[uint(left + (right - left) / 2)];

        while (i <= j) {
            while (array[uint(i)] < pivot) i++;
            while (pivot < array[uint(j)]) j--;

            if (i <= j) {
                uint tmp = array[uint(i)];
                array[uint(i)] = array[uint(j)];
                array[uint(j)] = tmp;
                i++;
                j--;
            }
        }

        if (left < j) sort(array, left, j);
        if (i < right) sort(array, i, right);
    }
}
