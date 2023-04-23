package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"ssv_operator_metadata/config"
	"ssv_operator_metadata/controllers"
	"ssv_operator_metadata/db"
	"ssv_operator_metadata/db/models"
	_ "ssv_operator_metadata/docs"
	"ssv_operator_metadata/eth"
	"ssv_operator_metadata/ipfs"
	"ssv_operator_metadata/routers"

	"github.com/gin-gonic/gin"
	log "github.com/ipfs/go-log/v2"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli/v2"
)

var Version string = "1.0.0"

// @title   SSV Operator API
// @version 1.0

// @contact.name  FileDrive
// @contact.url   https://filedrive.io/
// @contact.email filedriveteam@outlook.com

// @BasePath /api/v1
func main() {

	local := []*cli.Command{
		serverCmd,
		syncCmd,
		syncIpfsCmd,
	}

	app := &cli.App{
		Name:        "SSV Operators Server",
		Description: "SSV Operators Server",
		Commands:    local,
		Version:     Version,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
		os.Exit(1)
	}
}

var serverCmd = &cli.Command{
	Name:        "start",
	Usage:       "start the server",
	Description: "start the server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:   "config",
			Usage:  "config file name",
			Hidden: false,
			Value:  "app.toml",
		},
		&cli.StringFlag{
			Name:  "loglevel",
			Value: "info",
			Usage: "set log level",
		},
	},
	Action: func(c *cli.Context) error {
		log.SetLogLevel("*", c.String("loglevel"))

		cfg, err := config.LoadConfig(c.String("config"))
		if err != nil {
			panic(fmt.Errorf("load config failed:%v", err))
		}

		ginMode := cfg.App.Env

		serverAddr := fmt.Sprintf("%v:%d", cfg.Server.Host, cfg.Server.Port)
		fmt.Println(serverAddr)

		// todo : add net type control
		ctl := &controllers.Control{
			IpfsCli: ipfs.NewIpfsClient(cfg.Ipfs.Url),
			// Mainnet: &controllers.NetControl{
			// 	DM:   dm,
			// 	ECli: eCli,
			// },
			Testnet: controllers.NewNetControl(
				db.NewDBManager(cfg.Database.TestNet.Url),
				eth.NewEthClient(false, cfg.Eth.TestNet.Url)),
		}

		gin.SetMode(ginMode)
		router := routers.InitRouter(ctl)
		url := ginSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", serverAddr))
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		router.Run(serverAddr)

		return nil
	},
}

var syncCmd = &cli.Command{
	Name:        "sync",
	Usage:       "sync ssv info on chain",
	Description: "sync ssv info on chain",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:   "config",
			Usage:  "config file name",
			Hidden: false,
			Value:  "app.toml",
		},
		&cli.BoolFlag{
			Name:  "all",
			Usage: "scan all event log or only subscribe mode",
			Value: false,
		},
		&cli.StringFlag{
			Name:  "loglevel",
			Value: "info",
			Usage: "set log level",
		},
	},
	Action: func(c *cli.Context) error {
		log.SetLogLevel("*", c.String("loglevel"))
		cfg, err := config.LoadConfig(c.String("config"))
		if err != nil {
			panic(fmt.Errorf("load config failed:%v", err))
		}
		isAllScan := c.Bool("all")
		// todo : add net type control
		dbUrl := cfg.Database.TestNet.Url
		ethUrl := cfg.Eth.TestNet.Url

		dbm := db.NewDBManager(dbUrl)
		// init eth client
		eCli := eth.NewEthClient(false, ethUrl)

		infoChanBuf := make(chan interface{}, 10)
		go func() {
			for e := range infoChanBuf {
				// update operator base info
				if i, ok := e.(models.OperatorBaseInfo); ok {
					if isOk, err := dbm.UpdateOperatorBaseInfo(&i); err != nil {
						fmt.Printf("[%v] UpdateOperatorBaseInfo failed:%v\n", i.Id, err)
					} else {
						fmt.Println(i.Id, isOk)
					}
				}
			}
		}()
		defer func() {
			close(infoChanBuf)
		}()
		eCli.SyncSSVContractInfo(c.Context, isAllScan, infoChanBuf)

		<-c.Context.Done()

		return nil
	},
}

var syncIpfsCmd = &cli.Command{
	Name:        "sync-ipfs",
	Usage:       "sync ipfs ",
	Description: "sync ipfs ",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:   "config",
			Usage:  "config file name",
			Hidden: false,
			Value:  "app.toml",
		},
		&cli.StringFlag{
			Name:  "loglevel",
			Value: "info",
			Usage: "set log level",
		},
		&cli.StringFlag{
			Name:  "pprof",
			Value: "",
			Usage: "pprof url: if nil not open ,else open",
		},
	},
	Action: func(c *cli.Context) error {
		log.SetLogLevel("*", c.String("loglevel"))
		cfg, err := config.LoadConfig(c.String("config"))
		if err != nil {
			panic(fmt.Errorf("load config failed:%v", err))
		}
		// todo : add net type control
		dbUrl := cfg.Database.TestNet.Url
		dbm := db.NewDBManager(dbUrl)

		ipfsCtl := ipfs.NewControl(cfg.Ipfs.Url, cfg.Ipfs.NftStorageToken, "")

		ticker := time.NewTicker(time.Millisecond)
		once := sync.Once{}

		go func() {
			if _pprofUrl := c.String("pprof"); len(_pprofUrl) > 0 {
				fmt.Println("open pprof:", _pprofUrl)
				fmt.Println(http.ListenAndServe(_pprofUrl, nil))
			}
		}()

		for {
			select {
			case <-c.Done():
				fmt.Println("exit ....")
				return nil
			case <-ticker.C:
				fmt.Println("ticker sync ipfs ....")
				once.Do(func() {
					fmt.Println("do once")
					ticker.Reset(time.Minute * 1)
				})
				cids, err := dbm.GetNotSyncCids()
				if err != nil {
					fmt.Println("get no sync cid error:", err)
					continue
				}
				for _, cidRecord := range *cids {
					fmt.Println(cidRecord.Cid)
					if _cid, err := ipfsCtl.Sync(cidRecord.Cid); err == nil {
						fmt.Println("store as", _cid)
						newCidRecord := cidRecord
						newCidRecord.IsSync = true
						newCidRecord.UpdatedAt = time.Now().UnixMilli()
						success, err := dbm.UpdateCidRecord(&newCidRecord)
						if err != nil {
							fmt.Printf("update cid_record on %v failed: %v\n", newCidRecord.Cid, err)
							continue
						}
						fmt.Printf("update cid_record on %v to %v success: %v\n", newCidRecord.Cid, _cid, success)
					} else {
						fmt.Println("Sync: error:", err)
					}
				}
				fmt.Println("------------------end")
			}
		}
	},
}
