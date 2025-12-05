// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import {imitateIERC20} from "./imitateIERC20.sol";

contract imitateERC20 is Context, imitateIERC20, Ownable {
    mapping(address => uint256) private _balances;//每个地址的余额
    mapping(address => mapping(address => uint256)) private _allowances; // 授权额度 owner => spender => amount

    uint256 private _totalSupply; //代币总数量

    constructor() Ownable(msg.sender){
        _transferOwnership(msg.sender);
    }

    function balanceOf(address account) public view virtual returns (uint256){
        return _balances[account];
    }

    function transfer(address toAccount, uint256 value) public virtual returns (bool){
        address owner = _msgSender();
        _transfer(owner, toAccount, value);
        return true;
    }

    function transferFrom(address fromAccount, address toAccount, uint256 value) public virtual returns (bool){
        address spender = _msgSender();
        uint256 currentAllowValue = _allowances[fromAccount][spender];
        if (currentAllowValue <= 0) {
            revert("approve value is zero");
        }

        if (currentAllowValue < value) {
            revert("approve value is limit");
        }
        _approve(fromAccount, spender, currentAllowValue - value);
        _transfer(fromAccount, toAccount, value);
        return true;
    }

    function approve(address spenderAccount, uint256 value) public virtual returns (bool){
        address owner = _msgSender();
        _approve(owner, spenderAccount, value);
        return true;
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function _approve(address owner, address spender, uint256 value) internal virtual {
        require(owner != address(0), "approve from the zero address");
        require(spender != address(0), "approve to the zero address");
        _allowances[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    function _transfer(address from, address to, uint256 value) internal virtual {
        require(from != address(0), "transfer from the zero address");
        require(to != address(0), " transfer to the zero address");
        _update(from, to, value);
    }

    function _mint(address account, uint256 value) internal {
        require(account != address(0), "mind is zero address");
        _update(address(0), account, value);
    }

    function _update(address from, address to, uint256 value) internal virtual {
        if (from == address(0)) {
            _totalSupply += value;
        } else {
            uint256 fromBalance = _balances[from];
            require(fromBalance >= value, "_update amount exceeds balance");//from地址的余额不足

            unchecked{
                _balances[from] = fromBalance - value;
            }
        }

        if (to == address(0)) {
            unchecked{
                _totalSupply -= value;
            }
        } else {
            unchecked {
                _balances[to] += value;
            }
        }

        emit Transfer(from, to, value);
    }
}
