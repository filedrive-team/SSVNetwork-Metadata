package ipfs

import (
	"fmt"

	"github.com/web3-storage/go-w3s-client"
)

func InitWeb3Client(token string) w3s.Client {
	c, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		panic(fmt.Sprintf("web3 NewClient error:%v", err))
	}
	return c
}
