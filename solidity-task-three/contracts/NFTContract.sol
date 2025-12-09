// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";


contract NFTContract is ERC721, Ownable {
    uint256 private _nextTokenId;

    constructor() ERC721("NFTContract", "NFTC") Ownable(msg.sender) {
        _nextTokenId = 1;
    }

    function mint(address to) external onlyOwner returns (uint256) {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, _nextTokenId);
        return tokenId;
    }

    // 查询当前最新 tokenId
    function currentTokenId() external view returns (uint256) {
        return _nextTokenId - 1;
    }

    function _baseURI() internal pure override returns (string memory) {
        return "https://peach-quiet-canidae-313.mypinata.cloud/ipfs/bafkreibo55v3qcvfla35nsqnmiflvh7vjyyyg2x45pvtdxvzm65fovznom";
    }
}
