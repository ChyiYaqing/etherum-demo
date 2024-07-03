package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

/**
Go获取以太坊账户余额
*/

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	accountAddr := os.Getenv("ACCOUNT_ADDRESS")
	url := "https://eth-sepolia.g.alchemy.com/v2/" + apiKey
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Could not connect to Infura with ethclient: %s", err)
	}

	// 查询最新余额, nil
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(accountAddr), nil)
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("balance in Wei: %s", balance) // Wei
	bf := big.NewFloat(0).SetInt(balance)
	bf.Quo(bf, big.NewFloat(1e18))
	log.Printf("balance in Ether (converted by big.Float): %s", bf.String())

	bd := decimal.RequireFromString(balance.String())
	bd = bd.Div(decimal.NewFromFloat(1e18))
	log.Printf("balance in Ether (converted decimal.Decimal): %s", bd.String())
}
