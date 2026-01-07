import {expect} from "chai";
import {ethers} from "hardhat";

describe("MetaNodeToken", function () {
    it("should mint initial supply to deployer", async () => {
        const [owner] = await ethers.getSigners();

        const Token = await ethers.getContractFactory("MetaNodeToken");
        const token = await Token.deploy();
        await token.waitForDeployment();

        const totalSupply = await token.totalSupply();
        const ownerBalance = await token.balanceOf(owner.address);

        expect(ownerBalance).to.equal(totalSupply);
        expect(totalSupply).to.equal(
            ethers.parseEther("10000000")
        );
    });

    it("should have correct name and symbol", async () => {
        const Token = await ethers.getContractFactory("MetaNodeToken");
        const token = await Token.deploy();

        expect(await token.name()).to.equal("MetaNodeToken");
        expect(await token.symbol()).to.equal("MetaNode");
        expect(await token.decimals()).to.equal(18);
    });
});
