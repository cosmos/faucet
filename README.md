# Cosmos Testnet Faucet

This faucet app allows anyone who passes a captcha to request tokens for a Cosmos account address. This app needs to be deployed on a Cosmos testnet full node, because it relies on using the `gaiacli` command to send tokens.

## Prerequisites

### reCAPTCHA

If you don't have a reCAPTCHA site setup for the faucet, now is the time to get one. Go to the [Google reCAPTCHA Admin](https://www.google.com/recaptcha/admin) and create a new reCAPTCHA site. For the version of captcha, choose `reCAPTCHA v2`.

### Checkout Code

For the backend, need to have Go and the `dep` dependency tool installed on your system. For the frontend, you also need to have node.js and the `yarn` dependency tool installed on your system. 

```
go get git@github.com:cosmos/faucet
```

## Backend Setup

### Production

First, set the environment variables for the backend:

```
cd $GOPATH/src/github.com/cosmos/faucet/backend
vi .env
```

Then build the backend.

```
dep ensure
go build faucet.go
```

The following executable will run the faucet on port 8080. 

```
./faucet
```

**WARNING**: It's highly recommended to run a reverse proxy with rate limiting in front of this app. Included in this repo is an example `Caddyfile` that lets you run an TLS secured faucet that is rate limited to 1 claim per IP per day.

### Development

Run `go run faucet,.go` in the `backend` directory to serve the backend.

## Frontend Setup

### Production

First, set the environment variables for the frontend:

```
cd $GOPATH/src/github.com/cosmos/faucet/frontend
vi .env
```

Then build the frontend.

```
yarn
yarn build
```

Lastly, serve the `./frontend/dist` directory with the web server of your choice.

### Development

Run `yarn serve` in the `frontend` directory to serve the frontend with hot reload.
