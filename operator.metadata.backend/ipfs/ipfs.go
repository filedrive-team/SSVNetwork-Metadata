package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	log "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"
)

var logging = log.Logger("ipfs")

func NewIpfsClient(url string) *shell.Shell {
	return shell.NewShell(url)
}

func NewKobeClient(rpcMultiAddr string) *httpapi.HttpApi {
	_addr, err := ma.NewMultiaddr(rpcMultiAddr)
	if err != nil {
		panic(err)
	}
	kuboCli, err := httpapi.NewApi(_addr)
	if err != nil {
		panic(err)
	}
	return kuboCli
}
