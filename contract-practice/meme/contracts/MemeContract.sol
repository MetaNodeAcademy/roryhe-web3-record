// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {InterfaceMeme} from "./InterfaceMeme.sol";
import {InterfaceMemeFactory} from "./InterfaceMemeFactory.sol";


contract MemeContract is ERC20, Ownable, ReentrancyGuard {
    uint256 public constant TOTAL_SUPPLY = 1_000_000_000 * 1e18;

    InterfaceMeme public uniswapRouter;
    address public uniswapPair;

    address public taxWallet; // 税收地址（项目金库）

    //税费
    uint256 public buyTax = 3;      // 买入税 3%
    uint256 public sellTax = 5;     // 卖出税 5%
    uint256 public transferTax = 1; // 普通转账税 1%

    mapping(address => bool) public isTaxExempt;

    //交易限制
    uint256 public maxTxAmount;      // 单笔最大交易
    uint256 public maxWalletAmount; // 单地址最大持仓

    mapping(address => uint256) public dailyTxCount;
    mapping(address => uint256) public lastTxDay;
    uint256 public maxDailyTx = 20;

    constructor(address _router, address _taxWallet) ERC20("MemeContract", "MEME_CONTRACT") {
        _mint(msg.sender, TOTAL_SUPPLY);

        taxWallet = _taxWallet;

        uniswapRouter = InterfaceMeme(_router);
        uniswapPair = InterfaceMemeFactory(uniswapRouter.factory())
            .createPair(address(this), uniswapRouter.WETH());

        maxTxAmount = TOTAL_SUPPLY / 100;      // 1%
        maxWalletAmount = TOTAL_SUPPLY / 50;   // 2%

        isTaxExempt[msg.sender] = true;
        isTaxExempt[address(this)] = true;
        isTaxExempt[taxWallet] = true;
    }

    //转帐逻辑
    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal override {

        require(amount <= maxTxAmount || isTaxExempt[sender], "Exceeds max tx");
        if (recipient != uniswapPair) {
            require(balanceOf(recipient) + amount <= maxWalletAmount, "Exceeds max wallet");
        }

        _checkDailyLimit(sender);

        uint256 taxAmount = 0;

        if (!isTaxExempt[sender] && !isTaxExempt[recipient]) {
            if (sender == uniswapPair) {
                // Buy
                taxAmount = (amount * buyTax) / 100;
            } else if (recipient == uniswapPair) {
                // Sell
                taxAmount = (amount * sellTax) / 100;
            } else {
                // Normal transfer
                taxAmount = (amount * transferTax) / 100;
            }
        }

        if (taxAmount > 0) {
            super._transfer(sender, taxWallet, taxAmount);
        }

        super._transfer(sender, recipient, amount - taxAmount);
    }

    //交易限制-每日
    function _checkDailyLimit(address user) internal {
        uint256 currentDay = block.timestamp / 1 days;

        if (lastTxDay[user] < currentDay) {
            lastTxDay[user] = currentDay;
            dailyTxCount[user] = 0;
        }

        dailyTxCount[user] += 1;
        require(dailyTxCount[user] <= maxDailyTx, "Daily tx limit exceeded");
    }

    //流动性池实现
    function addLiquidity(uint256 tokenAmount) external payable onlyOwner nonReentrant {
        _approve(address(this), address(uniswapRouter), tokenAmount);

        uniswapRouter.addLiquidityETH{value: msg.value}(
            address(this),
            tokenAmount,
            0,
            0,
            owner(),
            block.timestamp
        );
    }

    function removeLiquidity(uint256 liquidity) external onlyOwner nonReentrant {
        _approve(uniswapPair, address(uniswapRouter), liquidity);

        uniswapRouter.removeLiquidityETH(
            address(this),
            liquidity,
            0,
            0,
            owner(),
            block.timestamp
        );
    }

    function setTaxes(uint256 _buy, uint256 _sell, uint256 _transfer) external onlyOwner {
        require(_buy <= 10 && _sell <= 10, "Tax too high");
        buyTax = _buy;
        sellTax = _sell;
        transferTax = _transfer;
    }

    function setLimits(uint256 _maxTx, uint256 _maxWallet, uint256 _dailyTx) external onlyOwner {
        maxTxAmount = _maxTx;
        maxWalletAmount = _maxWallet;
        maxDailyTx = _dailyTx;
    }

    function setTaxExempt(address user, bool exempt) external onlyOwner {
        isTaxExempt[user] = exempt;
    }

    receive() external payable {}
}
