# 操作指南

### 部署
```shell
npm install @openzeppelin/contracts
```
部署时传入：
```shell
router = UniswapV2Router02地址
taxWallet = 项目金库地址
```
### 使用说明 

#### 代币交易

用户直接在 Uniswap 交易

合约自动扣税

#### 添加流动性
```solidity
addLiquidity(tokenAmount) + msg.value(ETH)
```
#### 移除流动性
```solidity
removeLiquidity(lpAmount)
```
