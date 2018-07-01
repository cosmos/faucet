## Makefile

# Dependencies: yarn, dep

KEY ?= default
NODE ?= http://127.0.0.1:26657
CHAIN ?= gaia-6002
PASS ?= 1234567890

# Compile the faucet for deployment
build:
	dep ensure
	KEY=$(KEY) NODE=$(NODE) CHAIN=$(CHAIN) PASS=$(PASS) go build faucet.go
	cd faucet && yarn && yarn build

.PHONY: build

