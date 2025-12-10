###项目结构

```
├── contracts/
│ ├── NFTContract.sol # NFT合约
│ ├── AuctionContract.sol # 主拍卖合约（UUPS 可升级）
│
├── docs/
│ ├── coverage # 测试报告
├── scripts/
│ ├── deploy.js # 一键部署 UUPS 代理
│ ├── upgrade.js # 升级逻辑合约
│
├── test/
│ ├── Auction.unit.test.js # 单元测试
│ ├── Auction.integration.test.js # 集成测试
│
├── hardhat.config.js
├── package.json
└── README.md
```

### 合约地址

https://sepolia.etherscan.io/address/0xd31E734f2E61bc686CFad694c551E847bA637F27#code### 部署

```shell
npx hardhat run scripts/deploy.js --network sepolia
```

### 升级合约

```shell
npx hardhat run scripts/upgrade.js --network sepolia
```
