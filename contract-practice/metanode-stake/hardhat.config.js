require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("dotenv").config();

const {RPC_URL, PRIVATE_KEY} = process.env;

module.exports = {
    solidity: "0.8.22",
    settings: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
    networks: {
        sepolia: {
            url: RPC_URL,
            accounts: [PRIVATE_KEY]
        },
    }
};
