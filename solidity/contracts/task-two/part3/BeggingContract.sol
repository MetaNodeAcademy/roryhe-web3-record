// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

contract BeggingContract {
    address public owner; //所有者
    mapping(address => uint256)private donations; //每个捐赠者的捐赠金额

    event Donation(address indexed donor, uint256 indexed amount); //捐赠事件

    uint256 public donationStartTime;
    uint256 public donationEndTime;

    address[10] topsAddress;
    uint256[10] topsAmount;

    constructor(){
        owner = msg.sender;
        donationStartTime = block.timestamp;
        donationEndTime = block.timestamp + 30 days;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, unicode"不属于所有者权限");
        _;
    }

    modifier onlyDuringDonationPeriod(){
        require(block.timestamp >= donationStartTime && block.timestamp <= donationEndTime, unicode"捐赠有效时间已过期");
        _;
    }

    function getTime() external view returns (uint256, uint256){
        return (donationStartTime, donationEndTime);
    }

    //捐赠
    function donate() external payable onlyDuringDonationPeriod {
        require(msg.value > 0, unicode"捐赠金额必须大于0");

        donations[msg.sender] += msg.value;
        emit Donation(msg.sender, msg.value);

        _updateTops(msg.sender);
    }

    //提现
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, unicode"可提现余额不足");
        // 将合约余额转给拥有者
        payable(owner).transfer(balance);
    }

    //查询捐赠金额
    function getDonationAmount(address donorAddress) external view returns (uint256){
        return donations[donorAddress];
    }

    //排行榜前十
    function topDonors() external view returns (address[10] memory addrs, uint256[10] memory amounts){
        return (topsAddress, topsAmount);
    }


    //更新排序
    function _updateTops(address donor) internal {
        uint256 amount = donations[donor];

        for (uint256 i = 0; i < 10; i++) {
            if (topsAddress[i] == donor) {
                topsAmount[i] = amount;
                _sortTops();
                return;
            }
        }

        for (uint256 i = 0; i < 10; i++) {
            if (amount > topsAmount[i]) {
                topsAddress[i] = donor;
                topsAmount[i] = amount;
                _sortTops(); // 排序保持降序
                break;
            }
        }
    }

    //排序
    function _sortTops() internal {
        for (uint256 i = 0; i < 10; i++) {
            for (uint256 j = i + 1; j < 10; j++) {
                if (topsAmount[j] > topsAmount[i]) {
                    uint256 tempAmount = topsAmount[i];
                    topsAmount[i] = topsAmount[j];
                    topsAmount[j] = tempAmount;

                    address tempAddress = topsAddress[i];
                    topsAddress[i] = topsAddress[j];
                    topsAddress[j] = tempAddress;
                }
            }
        }
    }

}
