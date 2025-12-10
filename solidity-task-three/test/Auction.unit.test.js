const {expect} = require("chai");
const {ethers, upgrades} = require("hardhat");

describe("AuctionContract - Unit Test", function () {

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
         *  Deploy Auction (UUPS Proxy)
         * --------------------------- */
        const Auction = await ethers.getContractFactory("AuctionContract");
        auction = await upgrades.deployProxy(Auction, [], {initializer: "initialize"});
        await auction.waitForDeployment();

        /** ---------------------------
         *  Deploy Mock Chainlink Feed
         * --------------------------- */
        const Feed = await ethers.getContractFactory("MockV3Aggregator");
        const ethFeed = await Feed.deploy(8, "300000000000"); // ETH/USD = 3000
        await ethFeed.waitForDeployment();

        await auction.setPriceFeeds(ethers.ZeroAddress, await ethFeed.getAddress());
    });

    it("should create auction", async () => {
        await nft.connect(alice).approve(await auction.getAddress(), 1);
        await auction.connect(alice).createAuction(await nft.getAddress(), 1);

        const a = await auction.auctions(1);
        expect(a.seller).to.equal(alice.address);
        expect(a.tokenId).to.equal(1n);
    });

    it("ETH bid should update highest bidder", async () => {
        await nft.connect(alice).approve(await auction.getAddress(), 1);
        await auction.connect(alice).createAuction(await nft.getAddress(), 1);

        // Bob 出价 1 ETH
        await auction.connect(bob).bidAuction(
            1,
            ethers.ZeroAddress,
            0,
            {value: ethers.parseEther("1")}
        );

        const a = await auction.auctions(1);
        expect(a.highestBidder).to.equal(bob.address);
        expect(a.highestBidAmount).to.equal(ethers.parseEther("1"));
    });

});
