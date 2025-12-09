/**
 *  nft拍卖合约部署
 */

const {ethers, upgrades} = require('hardhat');

async function main() {
    const NFTContract = await ethers.getContractFactory('AuctionContract');
    const NFTUpgrade = await upgrades.deployProxy(
        NFTContract,
        {initializer: "initialize", kind: "uups"}
    );

    await NFTUpgrade.waitForDeployment();

    console.log('代理合约地址：', await NFTUpgrade.getAddress());

    main().catch((err) => {
        console.error(err);
        process.exit(1);
    });
}
