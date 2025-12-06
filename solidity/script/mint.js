const hre = require("hardhat");

async function main() {
    const nft = await hre.ethers.getContractAt(
        "TestNFT",
        "0x7d811b9DA7f302958a8700003197ef59E520d283"
    );

    // 铸造NFT给钱包
    const tx = await nft.mint("0x400401589f17086e11ccfb337b088cbfb3d6748e");
    await tx.wait();

    console.log("NFT 铸造成功，交易哈希:", tx.hash);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
