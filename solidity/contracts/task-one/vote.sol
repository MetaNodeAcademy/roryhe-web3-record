// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract vote {

    mapping(string => uint)private voteMap;

    string[] private candidateList;//所有出现过的候选人,因为map只能通过key获取，所以需要额外存储记录候选人有哪些

    function voteFunc(string calldata candidate) external {
        if (candidateList.length <= 0) {
            candidateList.push(candidate);
        } else {
            bool isExist = false;
            uint len = candidateList.length;
            for (uint i = 0; i < len; i++) {

                if (keccak256(bytes(candidateList[i])) == keccak256(bytes(candidate))) {
                    isExist = true;
                }
            }

            if (!isExist) {
                candidateList.push(candidate);
            }
        }

        voteMap[candidate] += 1;
    }

    function getVotes(string calldata candidate) external view returns (uint){
        return voteMap[candidate];
    }

    function resetVotes() external {
        uint len = candidateList.length;
        for (uint i = 0; i < len; i++) {
            voteMap[candidateList[i]] = 0;
        }
    }
}
