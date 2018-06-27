# Cosmos Testnet Faucet

This faucet app allows anyone to easily request 10 faucetToken and 1 steak.

## Get reCAPTCHA Key

Go to the [Google reCAPTCHA Admin](https://www.google.com/recaptcha/admin) and create a new reCAPTCHA site. For the version of captcha, choose `reCAPTCHA v2`.

In the file `./frontend/src/views/Faucet.vue` on line 60, change the `sitekey` to your new reCAPTCHA client side integration site key.

```
sitekey: "6LdqyV0UAAAAAEqgBxvSsDpL2aeTEgkz_VTz1Vi1"
```

## Prepare the Faucet

You need to have Golang and Yarn/Node.js installed on your system.

```
go get $GOPATH/src/github.com/cosmos/faucet
cd $GOPATH/src/github.com/cosmos/faucet
dep ensure

cd frontend
yarn && yarn build
cd ..
```

## Run the Faucet

This will run the faucet on port 8080. It's highly recommended that you run a reverse proxy with rate limiting in front of this app.

```
go run faucet.go RECAPTCHA_SERVER_SIDE_SECRET
```
