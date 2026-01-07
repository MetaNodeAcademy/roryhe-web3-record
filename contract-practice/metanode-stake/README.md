# MetaNodeStake 合约说明文档

## 1.BasicInfo

Contract Name: MetaNodeStake

Solidity Version: ^0.8.20

Network: Ethereum / EVM Compatible

Token Standard: ERC20 Compatible

Decimals: 18

## 接口定义及前端示例

### 1️⃣ depositETH —— ETH 质押

```solidity
function depositETH() external payable;
```

**语义说明**

- 向 `ETH_PID = 0` 的池子质押 ETH
    
- 仅当该池 `stTokenAddress == address(0)` 时有效
    

**前置条件**

- 池子存在
    
- `msg.value >= pool.minDepositAmount`
    
- 合约未暂停（`whenNotPaused`）
    

**前端示例**

```ts
await stakeContract.depositETH({
  value: ethers.parseEther("1")
});

```

### 2️⃣ deposit —— ERC20 质押

```solidity
function deposit(uint256 pid, uint256 amount) external;
```

**语义说明**

- 向指定池子质押 ERC20 代币
    
- `pid != ETH_PID`
    

**前置条件**

- 已执行 `ERC20.approve`
    
- `amount >= minDepositAmount`
    
- 合约未暂停
    

**前端调用顺序**

```ts
await erc20.approve(stakeAddress, amount);
await stakeContract.deposit(pid, amount);
```
### 3️⃣ unstake —— 发起解质押请求

```solidity
function unstake(uint256 pid, uint256 amount) external;
```

**语义说明**

- 不立即转账
    
- 生成一条 `UnstakeRequest`
    
- 等待 `unstakeLockedBlocks` 后才能 `withdraw`
    

**前置条件**

- `user.stAmount >= amount`
    
- 提现功能未暂停

```ts
await stakeContract.unstake(pid, amount);
```

### 4️⃣ withdraw —— 提取已解锁资产

```solidity
function withdraw(uint256 pid) external;
```

**语义说明**

- 扫描 `user.requests`
    
- 提取所有 `unlockBlocks <= block.number` 的请求
    
- ETH 或 ERC20 自动识别
    

**注意**

- 即使没有可提取金额，也会成功执行（amount = 0）

```ts
await stakeContract.withdraw(pid);
```

### 5️⃣ claim —— 领取 MetaNode 奖励

```solidity

function claim(uint256 pid) external;

```

**语义说明**

- 结算奖励
    
- 转账 MetaNode
    
- 更新 `finishedMetaNode`
    

**前置条件**

- claim 未暂停

```ts
await stakeContract.claim(pid);
```

### 6️⃣ poolLength —— 池子数量

```solidity
function poolLength() external view returns (uint256);
```

```ts
const length = await stakeContract.poolLength();
```

### 7️⃣ pendingMetaNode —— 当前区块的可领取奖励

```solidity
function pendingMetaNode(uint256 pid, address user)
  external
  view
  returns (uint256);

```

```ts
const reward = await stakeContract.pendingMetaNode(pid, user);
```

### 8️⃣ pendingMetaNode和pendingMetaNodeByBlockNumber —— 指定区块高度的奖励
```solidity
function pendingMetaNode(uint256 _pid, address _user)  
external  
view  
checkPid(_pid)  
returns (uint256)

function pendingMetaNodeByBlockNumber(
  uint256 pid,
  address user,
  uint256 blockNumber
) public view returns (uint256);
```

### 9️⃣ getMultiplier —— 奖励区间计算
```solidity
function getMultiplier(uint256 from, uint256 to)
  public
  view
  returns (uint256);

```

### 前端可读的状态变量

**Pool**

```
struct Pool {  
    address stTokenAddress;  
    uint256 poolWeight;  
    uint256 lastRewardBlock;  
    uint256 accMetaNodePerST;  
    uint256 stTokenAmount;  
    uint256 minDepositAmount;  
    uint256 unstakeLockedBlocks;  
}

Pool[] public pool;
```

```ts
const poolInfo = await stakeContract.pool(pid);
```

**User**

```solidity
struct User {  
    uint256 stAmount;  
    uint256 finishedMetaNode;  
    uint256 pendingMetaNode;  
    UnstakeRequest[] requests;  
}

mapping(uint256 => mapping(address => User)) public user;
```

```ts
const userInfo = await stakeContract.user(pid, userAddress);
// 返回 stAmount, finishedMetaNode, pendingMetaNode，requests
// 注意 requests只能逐index读取，无法一次性获取
```

**其他**

```solidity
startBlock()
endBlock()
MetaNodePerBlock()
totalPoolWeight()
withdrawPaused()
claimPaused()
MetaNode()
```

```ts
await stakeContract.startBlock() //uint256
await stakeContract.endBlock() //uint256
await stakeContract.MetaNodePerBlock() //uint256
await stakeContract.totalPoolWeight() //uint256
await stakeContract.withdrawPaused() //bool
await stakeContract.claimPaused() //bool
await stakeContract.MetaNode() //IERC20
```

## 事件
```solidity
1.Deposit
event Deposit(address indexed user, uint256 indexed poolId, uint256 amount);

2.RequestUnstake
event RequestUnstake(address indexed user, uint256 indexed poolId, uint256 amount);

3.Withdraw
event Withdraw(address indexed user, uint256 indexed poolId, uint256 amount, uint256 blockNumber);

4.Claim
event Claim(address indexed user, uint256 indexed poolId, uint256 MetaNodeReward);

```

```ts
//监听事件
stakeContract.on("Claim", (user, pid, amount) => {
  console.log("Claimed:", amount.toString());
});
```