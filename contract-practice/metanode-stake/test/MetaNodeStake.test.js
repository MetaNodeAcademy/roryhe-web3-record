import {ethers} from "hardhat";
import {expect} from "chai";

describe("MetaNodeStake", function () {
    let stake;
    let token;
    let owner;
    let alice;
    let bob;

    const META_PER_BLOCK = ethers.parseEther("10");

    beforeEach(async () => {
        [owner, alice, bob] = await ethers.getSigners();

        // Deploy MetaNodeToken
        const Token = await ethers.getContractFactory("MetaNodeToken");
        token = await Token.deploy();
        await token.waitForDeployment();

        // Deploy MetaNodeStake (UUPS)
        const Stake = await ethers.getContractFactory("MetaNodeStake");
        stake = await Stake.deploy();
        await stake.waitForDeployment();

        const currentBlock = await ethers.provider.getBlockNumber();

        await stake.initialize(
            token.target,
            currentBlock,
            currentBlock + 1000,
            META_PER_BLOCK
        );

        // 给 Stake 合约打钱用于发奖励
        await token.transfer(
            stake.target,
            ethers.parseEther("1000000")
        );
    });

    //初始化与角色
    it("should initialize correctly", async () => {
        expect(await stake.MetaNode()).to.equal(token.target);
        expect(await stake.MetaNodePerBlock()).to.equal(META_PER_BLOCK);

        expect(
            await stake.hasRole(await stake.ADMIN_ROLE(), owner.address)
        ).to.equal(true);
    });

    //添加 ETH 池（pid = 0）
    it("should add ETH pool as first pool", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            10,
            false
        );

        const pool = await stake.pool(0);
        expect(pool.stTokenAddress).to.equal(ethers.ZeroAddress);
        expect(pool.poolWeight).to.equal(100);
    });

    //ETH 质押
    it("should allow ETH deposit and update staking balance", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            10,
            false
        );

        await stake.connect(alice).depositETH({
            value: ethers.parseEther("1"),
        });

        const user = await stake.user(0, alice.address);
        expect(user.stAmount).to.equal(
            ethers.parseEther("1")
        );
    });


    //ERC20 池质押流程
    it("should allow ERC20 staking", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            10,
            false
        );

        await stake.addPool(
            token.target,
            100,
            0,
            10,
            false
        );

        await token.transfer(alice.address, ethers.parseEther("100"));
        await token.connect(alice).approve(
            stake.target,
            ethers.parseEther("100")
        );

        await stake.connect(alice).deposit(
            1,
            ethers.parseEther("10")
        );

        const user = await stake.user(1, alice.address);
        expect(user.stAmount).to.equal(
            ethers.parseEther("10")
        );
    });

    //pendingMetaNode

    it("should accumulate pending rewards correctly", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            1,
            false
        );

        await stake.connect(alice).depositETH({
            value: ethers.parseEther("1"),
        });

        // 推进区块
        await ethers.provider.send("evm_mine", []);
        await ethers.provider.send("evm_mine", []);

        const pending = await stake.pendingMetaNode(
            0,
            alice.address
        );

        expect(pending).to.be.gt(0);
    });

    //claim
    it("should allow claim rewards", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            1,
            false
        );

        await stake.connect(alice).depositETH({
            value: ethers.parseEther("1"),
        });

        await ethers.provider.send("evm_mine", []);
        await ethers.provider.send("evm_mine", []);

        const before = await token.balanceOf(alice.address);

        await stake.connect(alice).claim(0);

        const after = await token.balanceOf(alice.address);
        expect(after).to.be.gt(before);
    });

    //解锁逻辑
    it("should lock unstake and withdraw after unlock", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            2,
            false
        );

        await stake.connect(alice).depositETH({
            value: ethers.parseEther("1"),
        });

        await stake.connect(alice).unstake(
            0,
            ethers.parseEther("0.5")
        );

        // 未解锁不能提现
        await expect(
            stake.connect(alice).withdraw(0)
        ).to.not.changeEtherBalance(alice, 0);

        // 推进区块
        await ethers.provider.send("evm_mine", []);
        await ethers.provider.send("evm_mine", []);

        await expect(
            stake.connect(alice).withdraw(0)
        ).to.changeEtherBalance(
            alice,
            ethers.parseEther("0.5")
        );
    });

    //暂停逻辑
    it("should block withdraw when paused", async () => {
        await stake.addPool(
            ethers.ZeroAddress,
            100,
            0,
            1,
            false
        );

        await stake.pauseWithdraw();

        await expect(
            stake.connect(alice).withdraw(0)
        ).to.be.revertedWith("WITHDRAW_PAUSED");
    });


})
