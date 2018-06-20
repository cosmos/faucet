# Gaia Testnet Faucet

This faucet allows anyone to easily request 1 steak.

```
go get $GOPATH/src/github.com/cosmos/faucet
cd $GOPATH/src/github.com/cosmos/faucet
dep ensure

cd frontend
yarn && yarn build

cd ..
go run faucet.go
```
