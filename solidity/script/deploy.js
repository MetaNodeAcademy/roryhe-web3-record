const hre = require("hardhat");

async function main() {
    // 获取合约工厂
    const Token = await hre.ethers.getContractFactory("BeggingContract");

    // 部署合约（构造函数无参数）
    const token = await Token.deploy();
    await token.waitForDeployment();

    console.log("合约部署成功！");
    console.log("合约地址:", await token.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
