// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract TestNFT is ERC721, Ownable {
    uint256 public nextTokenId = 1;
    constructor() ERC721("TestNFT", "TNFT") Ownable(msg.sender) {}

    function mint(address to) external onlyOwner returns (uint256) {
        uint256 tokenId = nextTokenId++;
        _safeMint(to, tokenId);
        return tokenId;
    }

    function _baseURI() internal pure override returns (string memory) {
        return "https://peach-quiet-canidae-313.mypinata.cloud/ipfs/bafkreibo55v3qcvfla35nsqnmiflvh7vjyyyg2x45pvtdxvzm65fovznom";
    }
}
