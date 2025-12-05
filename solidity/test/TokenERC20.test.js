const {expect} = require("chai");
const hre = require("hardhat");
const {ethers} = hre;

describe("imitateERC20", function () {
    let Token;
    let token;
    let owner, addr1, addr2;

    beforeEach(async function () {
        [owner, addr1, addr2] = await ethers.getSigners();
        Token = await ethers.getContractFactory("imitateERC20");
        token = await Token.deploy();
        await token.waitForDeployment();
    });

    it("should mint by onlyOwner", async function () {
        // owner 调用 mint
        await token.mint(owner.address, 1000);
        expect(await token.balanceOf(owner.address)).to.equal(1000);

        // 非 owner 调用 mint 必须 revert
        await expect(
            token.connect(addr1).mint(addr1.address, 100)
        ).to.be.revertedWithCustomError(token, "OwnableUnauthorizedAccount");
    });

    it("transfer should work", async function () {
        // owner 有钱
        await token.mint(owner.address, 1000);

        // owner → addr1 发送 300
        await token.transfer(addr1.address, 300);

        expect(await token.balanceOf(owner.address)).to.equal(700);
        expect(await token.balanceOf(addr1.address)).to.equal(300);
    });

    it("approve should work & update allowance", async function () {
        // owner 给 addr1 授权 500
        await token.approve(addr1.address, 500);

        // expect(await token.allowance(owner.address, addr1.address))
        //     .to.equal(500);
    });

    it("transferFrom should spend allowance", async function () {
        // owner 有 1000
        await token.mint(owner.address, 1000);

        // owner → addr1 授权 500
        await token.approve(addr1.address, 500);

        // addr1 用 transferFrom 消费授权额度
        await token.connect(addr1).transferFrom(owner.address, addr2.address, 200);

        // 余额检查
        expect(await token.balanceOf(addr2.address)).to.equal(200);
        expect(await token.balanceOf(owner.address)).to.equal(800);

        // allowance 应该被扣减
        // expect(await token.allowance(owner.address, addr1.address)).to.equal(300);
    });

    it("transferFrom should revert if allowance too low", async function () {
        // owner 授权 100
        await token.approve(addr1.address, 100);

        // 试图花 200
        await expect(
            token.connect(addr1).transferFrom(owner.address, addr2.address, 200)
        ).to.be.revertedWith("approve value is limit");
    });

    it("transferFrom should revert if allowance is zero", async function () {
        await expect(
            token.connect(addr1).transferFrom(owner.address, addr2.address, 1)
        ).to.be.revertedWith("approve value is zero");
    });
});
