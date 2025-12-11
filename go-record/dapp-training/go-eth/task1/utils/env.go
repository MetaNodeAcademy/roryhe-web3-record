package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	RpcUrl     string
	PrivateKey string
}

func Load() *Config {
	cwd, _ := os.Getwd()
	fmt.Printf("当前工作目录: %s\n", cwd)
	err := godotenv.Load(filepath.Join(cwd, "/task1/utils/.env"))
	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}

	return &Config{
		RpcUrl:     os.Getenv("RPC_URL"),
		PrivateKey: os.Getenv("PRIVATE_KEY"),
	}
}
