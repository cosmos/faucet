# Cosmos Testnet Faucet

This faucet app allows anyone who passes a captcha to request tokens for a Cosmos account address. This app needs to be deployed on a Cosmos testnet full node, because it relies on using the `gaiacli` command to send tokens.

## Prerequisite: reCAPTCHA

If you don't have a reCAPTCHA site setup for the faucet, now is the time to get one. Go to the [Google reCAPTCHA Admin](https://www.google.com/recaptcha/admin) and create a new reCAPTCHA site. For the version of captcha, choose `reCAPTCHA v2`.

## Setup

You need to have Golang/Dep and Node.js/Yarn installed on your system.

```
go get $GOPATH/src/github.com/cosmos/faucet
cd $GOPATH/src/github.com/cosmos/faucet
dep ensure
cd frontend && yarn
```

Set the environment variables for the frontend and backend:

```
vi .env
vi frontend/.env
```

## Development

Edit the `./faucet.go` file to work on the backend. Run `go run faucet.go` to try it out.

Run `cd frontend && yarn serve` to preview the frontend with hot reload.

## Production

Build the backend first.

```
go build faucet.go
```

Then, build the front end.

```
cd frontend
yarn && yarn build
```

The following executable will run the faucet on port 8080. 

```
./faucet
```

**WARNING**: It's highly recommended to run a reverse proxy with rate limiting in front of this app. Included in this repo is an example `Caddyfile` that lets you run an TLS secured faucet that is rate limited to 1 claim per IP per day.
