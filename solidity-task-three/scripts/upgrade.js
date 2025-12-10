const {ethers, upgrades} = require("hardhat");

const PROXY_ADDRESS = "";//需要代理的合约地址

async function main() {
    const AuctionV2 = await ethers.getContractFactory("AuctionContract");

    console.log("Upgrading AuctionContract...");

    const upgraded = await upgrades.upgradeProxy(PROXY_ADDRESS, AuctionV2);
    await upgraded.waitForDeployment();

    console.log("AuctionContract upgraded. Address:", await upgraded.getAddress());
}

main().catch((e) => {
    console.error(e);
    process.exit(1);
});
