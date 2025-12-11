package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/task1/utils"
)

type QueryResult struct {
	blockHash         string
	timestamp         uint64
	transactionVolume int
}

func queryBlockchain(client *ethclient.Client, blockNumber int64) (*QueryResult, error) {
	blockInfo, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}
	result := QueryResult{
		timestamp:         blockInfo.Time(),
		transactionVolume: len(blockInfo.Transactions()),
		blockHash:         blockInfo.Hash().Hex(),
	}

	return &result, nil
}

func sendTransaction(client *ethclient.Client, key string) (string, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println("Error loading private key", err.Error())
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	wei := big.NewInt(100)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	toAddress := common.HexToAddress("0x870464ac244b125922777f82b2e7bc054295d282")

	tx := types.NewTransaction(nonce, toAddress, wei, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func main() {
	cfg := utils.Load()
	fmt.Println("cfg:", cfg)

	client, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer client.Close()

	blockchain, err := queryBlockchain(client, 9816115)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("blockchain-hash", blockchain.blockHash)
	fmt.Println("blockchain-time", blockchain.timestamp)
	fmt.Println("blockchain-transactions", blockchain.transactionVolume)

	transactionHex, err := sendTransaction(client, cfg.PrivateKey)
	if err != nil {
		return
	}

	// 执行后交易hash：0x8845d66bcee0719826b4d5fdb3d3df46ec6ca934246aed8e2dc7b6cb93a05b2d
	fmt.Println("transactionHex", transactionHex)
}
