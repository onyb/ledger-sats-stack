module github.com/ledgerhq/satstack

go 1.15

require (
	github.com/btcsuite/btcd v0.21.0-beta.0.20201114000516-e9c7a5ac6401
	github.com/btcsuite/btcutil v1.0.2
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/magefile/mage v1.10.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.9
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)

replace github.com/btcsuite/btcd v0.21.0-beta.0.20201114000516-e9c7a5ac6401 => github.com/onyb/btcd v0.20.1-beta.0.20201116101952-848ee6a30375
