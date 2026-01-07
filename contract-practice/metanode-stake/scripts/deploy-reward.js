const {ethers} = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deployer:", deployer.address);

    const RewardToken = await ethers.getContractFactory("RewardToken");

    //1 亿枚 reward token
    const initialSupply = ethers.parseEther("100000000");

    const rewardToken = await RewardToken.deploy(
        "MetaNode Reward Token",
        "MNODE-R",
        initialSupply
    );

    await rewardToken.waitForDeployment();

    console.log("RewardToken deployed to:", rewardToken.target);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
