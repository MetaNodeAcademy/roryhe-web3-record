const {ethers, upgrades} = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deployer:", deployer.address);

    // 1) Deploy reward token (MetaNodeToken)
    const Token = await ethers.getContractFactory("MetaNodeToken");
    const token = await Token.connect(deployer).deploy();
    await token.waitForDeployment();
    console.log("MetaNodeToken deployed:", token.target);

    // 2) Deploy MetaNodeStake (UUPS) and initialize
    const Stake = await ethers.getContractFactory("MetaNodeStake");
    console.log("Deploying MetaNodeStake (UUPS proxy) ...");
    const stakeProxy = await upgrades.deployProxy(Stake, [
        token.target,            // _MetaNode (reward token)
        (await ethers.provider.getBlockNumber()), // startBlock
        (await ethers.provider.getBlockNumber()) + 100000, // endBlock
        ethers.parseEther("10")  // MetaNodePerBlock (10 tokens per block)
    ], {kind: "uups"});
    await stakeProxy.waitForDeployment();
    console.log("MetaNodeStake proxy deployed:", stakeProxy.target);

    // 3) Transfer reward tokens to Stake contract
    const fundingAmount = ethers.parseEther("50000"); // fund 50k tokens for rewards
    const txFund = await token.connect(deployer).transfer(stakeProxy.target, fundingAmount);
    await txFund.wait();
    console.log(`Transferred ${fundingAmount} tokens to stake contract`);

    // 4) Add ETH pool (pid=0). According to your contract, first pool must be ETH address(0)
    // Parameters: stTokenAddress, poolWeight, minDepositAmount, unstakeLockedBlocks, withUpdate
    // For ETH pool stTokenAddress = ethers.ZeroAddress
    const txAddEthPool = await stakeProxy.connect(deployer).addPool(
        ethers.ZeroAddress,
        100,       // poolWeight
        0,         // minDepositAmount
        10,        // unstakeLockedBlocks (in blocks)
        false      // withUpdate
    );
    await txAddEthPool.wait();
    console.log("Added ETH pool at pid 0");

    // 5) Optional: deploy Mock ERC20 to use as stake token and add another pool
    const MockToken = await ethers.getContractFactory("MetaNodeToken"); // reuse same contract for mock
    const mock = await MockToken.connect(deployer).deploy();
    await mock.waitForDeployment();
    console.log("Mock staking token deployed:", mock.target);

    const txAddErcPool = await stakeProxy.connect(deployer).addPool(
        mock.target,
        100,
        0,
        10,
        false
    );
    await txAddErcPool.wait();
    console.log("Added ERC20 pool at pid 1 with token:", mock.target);

    console.log("Deployment finished.");
    console.log("MetaNodeToken:", token.target);
    console.log("StakeProxy:", stakeProxy.target);
    console.log("Mock staking token:", mock.target);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
