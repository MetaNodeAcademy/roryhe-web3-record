const {expect} = require("chai");
const {ethers, upgrades} = require("hardhat");

describe("AuctionContract - Integration Test", function () {

    let auction, erc20, nft, owner, alice, bob;

    beforeEach(async () => {
        [owner, alice, bob] = await ethers.getSigners();

        /** ---------------------------
         *  Deploy Mock ERC20
         * --------------------------- */
        const ERC20 = await ethers.getContractFactory("MockERC20");
        erc20 = await ERC20.deploy();
        await erc20.waitForDeployment();

        /** ---------------------------
         *  Deploy Mock NFT
         * --------------------------- */
        const NFT = await ethers.getContractFactory("MockNFT");
        nft = await NFT.deploy();
        await nft.waitForDeployment();
        await nft.mint(alice.address);

        /** ---------------------------
         *  Deploy AuctionContract Proxy (UUPS)
         * --------------------------- */
        const Auction = await ethers.getContractFactory("AuctionContract");
        auction = await upgrades.deployProxy(Auction, [], {initializer: "initialize"});
        await auction.waitForDeployment();

        /** ---------------------------
         *  Deploy Chainlink Mock Feed
         * --------------------------- */
        const Feed = await ethers.getContractFactory("MockV3Aggregator");
        const ethFeed = await Feed.deploy(8, "300000000000"); // ETH/USD = 3000
        await ethFeed.waitForDeployment();
        const erc20Feed = await Feed.deploy(8, "100000000"); // 1 ERC20 = $1.0
        await erc20Feed.waitForDeployment();


        await auction.setPriceFeeds(ethers.ZeroAddress, await ethFeed.getAddress());
        await auction.setPriceFeeds(await erc20.getAddress(), await erc20Feed.getAddress());
    });

    it("complete auction flow (ETH)", async () => {
        // Alice 上架 NFT
        await nft.connect(alice).approve(await auction.getAddress(), 1);
        await auction.connect(alice).createAuction(await nft.getAddress(), 1);

        // Bob 出价 1 ETH
        await auction.connect(bob).bidAuction(
            1,                      // auctionId
            ethers.ZeroAddress,     // ETH
            0,                      // erc20Amount
            {value: ethers.parseEther("1")}
        );

        // 快进时间 1 天，拍卖结束
        await ethers.provider.send("evm_increaseTime", [86400]);
        await ethers.provider.send("evm_mine");

        const balanceBefore = await ethers.provider.getBalance(alice.address);

        // Owner 结算拍卖
        await auction.connect(owner).settleAuction(1);

        const balanceAfter = await ethers.provider.getBalance(alice.address);

        expect(balanceAfter).to.be.gt(balanceBefore);
    });

    it("complete auction flow (ERC20)", async () => {
        // Alice 上架 NFT
        await nft.connect(alice).approve(await auction.getAddress(), 1);
        await auction.connect(alice).createAuction(await nft.getAddress(), 1);

        // Mint ERC20 给 Bob
        await erc20.mint(bob.address, ethers.parseUnits("1000", 18));

        // Bob approve 给拍卖合约
        await erc20.connect(bob).approve(await auction.getAddress(), ethers.parseUnits("1000", 18));

        // Bob 出价 100 ERC20
        await auction.connect(bob).bidAuction(
            1,
            await erc20.getAddress(),              // ERC20 token
            ethers.parseUnits("100", 18)          // 出价数量
        );

        // 快进时间 1 天
        await ethers.provider.send("evm_increaseTime", [86400]);
        await ethers.provider.send("evm_mine");

        // 结算拍卖
        const sellerBalanceBefore = await erc20.balanceOf(alice.address);
        await auction.connect(owner).settleAuction(1);
        const sellerBalanceAfter = await erc20.balanceOf(alice.address);

        expect(sellerBalanceAfter).to.be.gt(sellerBalanceBefore);
    });

});
