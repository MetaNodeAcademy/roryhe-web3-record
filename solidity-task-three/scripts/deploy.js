const {ethers, upgrades} = require("hardhat");

async function main() {
    const Auction = await ethers.getContractFactory("AuctionContract");

    console.log("Deploying AuctionContract...");

    const auction = await upgrades.deployProxy(Auction, [], {
        initializer: "initialize",
        kind: "uups",
    });

    await auction.waitForDeployment();

    console.log("AuctionContract proxy deployed to:", await auction.getAddress());

    // const admin = await upgrades.admin.getInstance();
    // console.log("ProxyAdmin:", await admin.getAddress());
}

main().catch((e) => {
    console.error(e);
    process.exit(1);
});
