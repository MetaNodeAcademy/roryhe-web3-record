// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract AuctionContract is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    //拍卖结构数据
    struct Auction {
        address seller;//卖家
        address nftContract;//nft合约地址
        uint256 tokenId;//nftTokenId
        uint256 endTime;//拍卖结束时间

        address highestBidder;//最高出价者
        int256 highestBid;//最高出价金额USD
        uint256 highestBidAmount;//当前出价方式的最高出价金额（可能是ETH也可能是ERC20，取决于出价方式）
        address paymentToken;//出价方式（ETH/ERC20）；address(0) -> ETH，其他address->ERC20

        bool settled;//该竞拍是否已结束
    }

    mapping(uint256 => Auction) public auctions;
    uint256 public auctionCounter;

    uint256 public feeRate;        // 千分比费率，例如 200 = 2%
    address public feeRecipient;   // 手续费接收地址

    //预言机
    mapping(address => AggregatorV3Interface) public priceFeeds;

    event AuctionCreateEvent(uint256 actionId, address seller, address nftContract, uint256 tokenId, uint256 endTime, address paymentToken);

    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        feeRate = 200;//2%
        feeRecipient = msg.sender;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    //修改手续费
    function updateFee(uint256 rate) external onlyOwner {
        require(rate <= 1000 && rate >= 0, "fee rate needs to be within a reasonable range");
        feeRate = rate;
    }

    //上架NFT
    function createAuction(address nftContract, uint256 tokenId) external returns (uint256 auctionCounterId){
        //NFT锁定
        IERC721(nftContract).transferFrom(msg.sender, address(this), tokenId);

        auctionCounter++;
        auctions[auctionCounter] = Auction({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            endTime: block.timestamp + 1 days,
            highestBidder: address(0),
            highestBid: 0,
            highestBidAmount: 0,
            paymentToken: address(0),//默认ETH出价
            settled: false
        });

        emit AuctionCreateEvent(auctionCounter, msg.sender, nftContract, tokenId, auctions[auctionCounter].endTime, auctions[auctionCounter].paymentToken);

        return auctionCounter;
    }

    //出价
    function bidAuction(uint256 auctionId, address tokenAddress, uint256 erc20Amount) external payable {
        Auction storage auction = auctions[auctionId];
        uint256 bidAmount;
        int256 bidUSD;
        require(block.timestamp < auction.endTime && !auction.settled, "Auction ended");

        //非ETH出价
        if (tokenAddress != address(0)) {
            require(erc20Amount > 0, "erc20Amount required");
            IERC20 token = IERC20(tokenAddress);
            // 用户必须先 approve
            bidAmount = erc20Amount;
            require(token.transferFrom(msg.sender, address(this), bidAmount), "ERC20 transfer failed");
        } else {
            require(msg.value > 0, "bid required");
            bidAmount = msg.value;
        }

        bidUSD = getUSDValue(tokenAddress, bidAmount);
        require(bidUSD > auction.highestBid, "bid below to highestBid");

        //更新最高出价者：先返回旧最高出价，再记录最新的最高出价
        if (auction.highestBidder != address(0)) {
            if (auction.paymentToken == address(0)) {
//            ETH出价，返回给ETH
                (bool success,) = payable(auction.highestBidder).call{value: auction.highestBidAmount}("");
                require(success, "ETH bid return to bidder failed");
            } else {
//            ERC20出价返回给ERC20
                IERC20(auction.paymentToken).transfer(auction.highestBidder, auction.highestBidAmount);
            }
        }

        auction.highestBidder = msg.sender;
        auction.highestBid = bidUSD;
        auction.highestBidAmount = bidAmount;
        auction.paymentToken = tokenAddress;
    }

    //结束拍卖
    function settleAuction(uint256 auctionId) external {
        Auction storage auction = auctions[auctionId];
        require(block.timestamp >= auction.endTime, "Auction not ended");
        require(!auction.settled, "Auction already settled");

        auction.settled = true;

        uint256 fee = auction.highestBidAmount * feeRate / 10000;
        uint256 sellerReceive = auction.highestBidAmount - fee;

        if (auction.highestBidder != address(0)) {
            //NFT 转移给最高出价者，钱 转给卖家
            IERC721(auction.nftContract).transferFrom(address(this), auction.highestBidder, auction.tokenId);

            if (auction.paymentToken == address(0)) {

                (bool fs,) = payable(feeRecipient).call{value: fee}("");
                require(fs, "fee transfer failed");

                (bool success,) = payable(auction.seller).call{value: sellerReceive}("");
                require(success, "bid pay for seller failed");
            } else {

                IERC20 token = IERC20(auction.paymentToken);
                token.transfer(feeRecipient, fee);
                token.transfer(auction.seller, sellerReceive);}
        } else {
            //NFT 退回卖家
            IERC721(auction.nftContract).transferFrom(address(this), auction.seller, auction.tokenId);
        }
    }

    //获取USD转换
    function getUSDValue(address token, uint256 value) public view returns (int256 USD){
        AggregatorV3Interface feed = priceFeeds[token];
        require(address(feed) != address(0), "Feed not set");
        (, int256 price,,,) = feed.latestRoundData();
        return int256(value) * price / 1e8;
    }

    //设置预言机
    function setPriceFeeds(address token, address feedAddress) external onlyOwner {
        priceFeeds[token] = AggregatorV3Interface(feedAddress);
    }

    //动态手续费
}
