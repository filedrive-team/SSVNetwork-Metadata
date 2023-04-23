package config

import (
	"github.com/BurntSushi/toml"
)

type TomlApp struct {
	Env  string
	Swag bool
}

type TomlServer struct {
	Port         int
	Host         string
	WriteTimeout int
}
type TomlEth struct {
	Url               string
	SsvNetWorkAddress string
}
type EthCfg struct {
	MainNet TomlEth
	TestNet TomlEth
}
type TomlDatabase struct {
	Url string
}
type DatabaseCfg struct {
	MainNet TomlDatabase
	TestNet TomlDatabase
}
type TomlIpfs struct {
	Url             string
	NftStorageToken string
}
type Config struct {
	App    TomlApp
	Server TomlServer
	// Eth    TomlEth
	Eth      EthCfg
	Database DatabaseCfg
	Ipfs     TomlIpfs
}

func LoadConfig(configFileName string) (*Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(configFileName, &cfg)
	return &cfg, err
}
