package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/task1/constracts"
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
	privateKey, fromAddress, err := ecdsaTool(key)
	if err != nil {
		return "", err
	}

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

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	fmt.Println("执行成功")

	return signedTx.Hash().Hex(), nil
}

func ecdsaTool(key string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println("Error loading private key", err.Error())
		return nil, common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, common.Address{}, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, fromAddress, nil
}

func useAbigenTool(client *ethclient.Client, key string) (int, error) {
	privateKey, fromAddress, err := ecdsaTool(key)
	if err != nil {
		return 0, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//solAddress, tx, solInstance, err := constracts.DeployConstracts(auth, client)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return 0, err
	//}
	//
	//fmt.Println(solAddress.Hex())
	//fmt.Println(tx.Hash().Hex())
	//_ = solInstance

	//部署完成，需要加载调用合约并执行内部方法
	//simpleContracts, err := constracts.NewConstracts(solAddress, client)
	simpleContracts, err := constracts.NewConstracts(common.HexToAddress("0x18CeF2d6Eb6B332Ddd1804E7e9F9D5673c089CEF"), client)
	if err != nil {
		fmt.Println("create contracts instance fail", err.Error())
		return 0, err
	}

	callOpts := &bind.CallOpts{Context: context.Background()}

	addTx, err := simpleContracts.Add(auth)
	if err != nil {
		return 0, err
	}

	fmt.Println("call simple.sol func add success", addTx.Hash().Hex())

	tokenId, err := simpleContracts.NextTokenId(callOpts)
	if err != nil {
		return 0, err
	}

	fmt.Println("NextTokenId:", tokenId)

	return int(tokenId.Int64()), nil
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

	//transactionHex, err := sendTransaction(client, cfg.PrivateKey)
	//if err != nil {
	//	return
	//}

	// 执行后交易hash：0x8845d66bcee0719826b4d5fdb3d3df46ec6ca934246aed8e2dc7b6cb93a05b2d
	//fmt.Println("transactionHex", transactionHex)

	result, err := useAbigenTool(client, cfg.PrivateKey)
	if err != nil {
		return
	}
	fmt.Println("result", result)
}

/**
 * output：
blockchain-hash 0xe7df8ebad086554c63488121110ea462c87d3a40750ec616c79d2c8f77293dbb
blockchain-time 1765443708
blockchain-transactions 112
call simple.sol func add success 0x2fed2cb859014b906aaba9f10657ec846373399b3879cffe235f6e988402ceff
NextTokenId: 1
result 1
*/
