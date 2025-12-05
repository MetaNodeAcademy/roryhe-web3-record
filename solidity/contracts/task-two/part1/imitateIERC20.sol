// SPDX-License-Identifier: MIT

/**
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
合约包含以下标准 ERC20 功能：,
balanceOf：查询账户余额。,
transfer：转账。,
approve 和 transferFrom：授权和代扣转账。,
使用 event 记录转账和授权操作。,
提供 mint 函数，允许合约所有者增发代币。,
提示：
使用 mapping 存储账户余额和授权信息。,
使用 event 定义 Transfer 和 Approval 事件。,
部署到sepolia 测试网，导入到自己的钱包
 */
pragma solidity ^0.8.20;

interface imitateIERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approval(address indexed owner, address indexed spender, uint256 value);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address toAccount, uint256 value) external returns (bool);

    function approve(address spenderAccount, uint256 value) external returns (bool);

    function transferFrom(address fromAccount, address toAccount, uint256 value) external returns (bool);
}
