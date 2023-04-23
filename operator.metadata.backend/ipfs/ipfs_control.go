package ipfs

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/web3-storage/go-w3s-client"
)

type IpfsControl struct {
	ctx     context.Context
	ipfsCli *shell.Shell
	nftsCli *NftStorageCli
	w3sCli  w3s.Client
}

func NewControl(ipfsUrl, nftsToken, w3sToken string) *IpfsControl {
	return &IpfsControl{
		ctx:     context.Background(),
		ipfsCli: NewIpfsClient(ipfsUrl),
		nftsCli: NewNftStorageClient("", nftsToken),
		// w3sCli:  InitWeb3Client(w3sToken),
	}
}

func (c *IpfsControl) Sync(cid string) (string, error) {
	// check status
	resStatus, err := c.nftsCli.Status(c.ctx, cid)
	if err == nil {
		fmt.Println("status ok : ", resStatus.Value.Cid)
		return resStatus.Value.Cid, nil
	} else {
		// fmt.Printf("[status:%v]todo upload ...\n", err)
	}

	// do upload
	rd, err := c.ipfsCli.Cat(cid)
	if err != nil {
		return "", fmt.Errorf("get local cid %v failedL:%v", cid, err)
	}
	defer rd.Close()
	// do upload
	resUpload, err := c.nftsCli.Upload(c.ctx, rd)
	if err != nil {
		return "", fmt.Errorf("nft storage upload failed:%v", err)
	}

	return resUpload.Value.Cid, nil
}
