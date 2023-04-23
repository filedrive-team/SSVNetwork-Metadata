package ipfs_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"ssv_operator_metadata/ipfs"

	"github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/web3-storage/go-w3s-client"
)

var web3Token = "your_web3_storage_token"

func TestAddFile(t *testing.T) {
	c := ipfs.InitWeb3Client(web3Token)

	fileName := "/Users/john/Downloads/it-1.jpg"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("file open error:", err)
	}

	cid, err := c.Put(context.Background(), f, w3s.WithDirname("it-1.jpg"))
	if err != nil {
		log.Fatal("web3 put error:", err)
	}
	fmt.Println(cid.String())
}

func TestIpfs(t *testing.T) {
	fileName := "/Users/john/workspace/ssv/source/operators-meta-backend/upload/1667214871648_animation_500_l1661cp6.gif"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("file open error:", err)
	}

	sh := shell.NewShell("localhost:5001")
	_cid, err := sh.Add(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	//
	c, _ := cid.Decode(_cid)
	// cid.NewCidV0(c.Hash())

	fmt.Println("added:", _cid, cid.NewCidV1(c.Type(), c.Hash()))
}
func TestIpfsKubeRpc(t *testing.T) {
	kuboCli := ipfs.NewKobeClient("/ip4/127.0.0.1/tcp/5001")
	type pinRefKeyList struct {
		Keys map[string]struct {
			Type string
		}
	}
	var out pinRefKeyList
	err := kuboCli.Request("pin/ls").Exec(context.Background(), &out)
	if err != nil {
		panic(err)
	}
	for k, v := range out.Keys {
		fmt.Println(k, v)
	}

}

func TestPinata(t *testing.T) {
	url := "https://api.pinata.cloud/data/pinList?status=pinned&pinSizeMin=100"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer your_token")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestNftStorage(t *testing.T) {
	cli := ipfs.NewNftStorageClient("https://api.nft.storage", os.Getenv("NFT_STORAGE_API_KEYS"))
	// test status
	res, err := cli.Status(context.TODO(), "bafybeiae6x4gjyd3s3zqltrhtnxxkz6ahlewxzplsn6nkioia2zj52vaji")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", *res.Value)
	}
	return

	// test upload
	fileName := "/Users/john/Downloads/it-5.1.png"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("file open error:", err)
	}
	res1, err := cli.Upload(context.TODO(), f)
	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}
	// fmt.Println(*res)
	fmt.Printf("%+v\n", *res1.Value)
}

func TestIpfsControl(t *testing.T) {
	ctl := ipfs.NewControl("localhost:5001", os.Getenv("NFT_STORAGE_API_KEYS"), "")

	_cid, err := ctl.Sync("QmeuwQsGEo4vqiqYkJxZJmQMMLdaMMKdMnTWsDvQa6F5zJ")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("nft storage cid:", _cid)
	time.Sleep(time.Second)
}
