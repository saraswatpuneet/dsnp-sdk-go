package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	wallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// Language: go
// Path: core\config\config.go
// Core config provides for various settings and plugable modules.
type Config struct {
	Provider *ethclient.Client
	Signer   *wallet.Wallet
}
