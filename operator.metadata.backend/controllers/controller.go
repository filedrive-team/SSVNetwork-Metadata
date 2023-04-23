package controllers

import (
	"ssv_operator_metadata/db"
	"ssv_operator_metadata/eth"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/ipfs/go-log/v2"
)

var logging = log.Logger("controllers")

type NetControl struct {
	*db.DBManager
	*eth.Client
}
type Control struct {
	Mainnet *NetControl
	Testnet *NetControl

	IpfsCli *shell.Shell
}

func NewNetControl(dm *db.DBManager, eCli *eth.Client) *NetControl {
	return &NetControl{
		DBManager: dm,
		Client:    eCli,
	}
}
