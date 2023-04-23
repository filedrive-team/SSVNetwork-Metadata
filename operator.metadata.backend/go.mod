module ssv_operator_metadata

go 1.16

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/ethereum/go-ethereum v1.11.2
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/gin-contrib/gzip v0.0.6
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/ipfs/go-cid v0.3.2
	github.com/ipfs/go-ipfs-api v0.3.0
	github.com/ipfs/go-ipfs-http-client v0.4.0
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/ipld/go-ipld-prime v0.19.0 // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.3 // indirect
	github.com/libp2p/go-libp2p v0.24.2 // indirect
	github.com/libp2p/go-libp2p-core v0.17.0 // indirect
	github.com/libp2p/go-libp2p-mplex v0.5.0 // indirect
	github.com/libp2p/go-libp2p-record v0.2.0 // indirect
	github.com/libp2p/go-libp2p-transport-upgrader v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/multiformats/go-multiaddr v0.8.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/polydawn/refmt v0.89.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969 // indirect
	github.com/storyicon/sigverify v1.1.0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.3
	github.com/swaggo/swag v1.8.6
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/urfave/cli/v2 v2.16.3
	github.com/web3-storage/go-w3s-client v0.0.6
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20230105000112-eab7a2c85304 // indirect
	golang.org/x/tools v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/postgres v1.4.5
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755
	gorm.io/plugin/dbresolver v1.3.0
)

replace github.com/prysmaticlabs/prysm => github.com/prysmaticlabs/prysm v1.4.2-0.20211101172615-63308239d94f

replace github.com/libp2p/go-libp2p-core => github.com/libp2p/go-libp2p-core v0.9.0

replace github.com/libp2p/go-libp2p-transport-upgrader => github.com/libp2p/go-libp2p-transport-upgrader v0.4.3

replace github.com/ipld/go-car => github.com/ipld/go-car v0.3.3

replace github.com/ipfs/go-path => github.com/ipfs/go-path v0.2.1

replace github.com/ipfs/go-ipfs-blockstore => github.com/ipfs/go-ipfs-blockstore v1.1.2

// v3
replace github.com/decred/dcrd/dcrec/secp256k1/v4 => github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0

replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.25

replace github.com/gin-contrib/gzip => github.com/gin-contrib/gzip v0.0.6

replace github.com/google/go-cmp => github.com/google/go-cmp v0.5.8

replace github.com/ipfs/go-cid => github.com/ipfs/go-cid v0.1.0

replace github.com/libp2p/go-buffer-pool => github.com/libp2p/go-buffer-pool v0.1.0

replace github.com/libp2p/go-flow-metrics => github.com/libp2p/go-flow-metrics v0.1.0

replace github.com/libp2p/go-openssl v0.1.0 => github.com/libp2p/go-openssl v0.1.0

replace github.com/multiformats/go-base32 => github.com/multiformats/go-base32 v0.1.0

replace github.com/multiformats/go-multiaddr => github.com/multiformats/go-multiaddr v0.7.0

replace github.com/multiformats/go-multibase => github.com/multiformats/go-multibase v0.1.1

replace github.com/multiformats/go-multicodec => github.com/multiformats/go-multicodec v0.6.0

replace github.com/multiformats/go-multihash => github.com/multiformats/go-multihash v0.2.1

replace lukechampine.com/blake3 => lukechampine.com/blake3 v1.1.7

replace google.golang.org/protobuf => google.golang.org/protobuf v1.28.1

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c

replace github.com/libp2p/go-libp2p-testing => github.com/libp2p/go-libp2p-testing v0.9.2
