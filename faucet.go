package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const tpl = `
<!doctype html>

<html lang="en">
<head>
	<meta charset="utf-8">

	<title>Cosmos Hub Testnet Faucet</title>
	<meta name="description" content="A faucet to get some coins.">
	<meta name="author" content="Adrian Brink">

	<link href="https://fonts.googleapis.com/css?family=Slabo+27px" rel="stylesheet">
	<link rel="stylesheet" href="https://unpkg.com/blaze/scss/dist/components.buttons.min.css">
	<link rel="stylesheet" href="https://unpkg.com/blaze/scss/dist/components.inputs.min.css">

	<style>
		body {
			font-family: 'Slabo 27px', serif;
			font-size: 21px;
		}
	</style>

	<!--[if lt IE 9]>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
	<![endif]-->
</head>

<body>
	<div style="width:40%; margin:0 auto;">
		{{.}}
	</div>
</body>
</html>`

var pageTemplate = template.Must(template.New("page").Parse(tpl))

const amount = "10steak"
const key = "faucet"
const node = "http://localhost:46657"
const chain = "test-chain-Ra9zi4"
const pass = "1234567890"
const faucet = "BB5ED43DA0FA8EB0CA447DA966C9A33BC4187E63"

func main() {
	fmt.Println("faucet")

	http.HandleFunc("/", faucetHandler)
	http.HandleFunc("/claim", getCoinsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func executeCmd(command string, writes ...string) {
	cmd, wc, _ := goExecute(command)

	for _, write := range writes {
		wc.Write([]byte(write + "\n"))
	}
	cmd.Wait()
}

func executeGetSequence(addr string) (sequence int64) {
	command := fmt.Sprintf("gaiacli account %v --node=%v --chain-id=%v", addr, node, chain)
	cmd := getCmd(command)
	bz, _ := cmd.CombinedOutput()
	out := strings.Trim(string(bz), "\n")
	time.Sleep(time.Second)

	var res map[string]json.RawMessage
	json.Unmarshal([]byte(out), &res)

	var value map[string]json.RawMessage
	json.Unmarshal([]byte(res["value"]), &value)

	json.Unmarshal([]byte(value["sequence"]), &sequence)

	return sequence
}

func goExecute(command string) (cmd *exec.Cmd, pipeIn io.WriteCloser, pipeOut io.ReadCloser) {
	cmd = getCmd(command)
	pipeIn, _ = cmd.StdinPipe()
	pipeOut, _ = cmd.StdoutPipe()
	go cmd.Start()
	time.Sleep(time.Second)
	return cmd, pipeIn, pipeOut
}

func getCmd(command string) *exec.Cmd {
	// split command into command and args
	split := strings.Split(command, " ")

	var cmd *exec.Cmd
	if len(split) == 1 {
		cmd = exec.Command(split[0])
	} else {
		cmd = exec.Command(split[0], split[1:]...)
	}

	return cmd
}

func faucetHandler(w http.ResponseWriter, r *http.Request) {
	data := `
	<h1>Cosmos Hub Testnet Faucet</h1>
	<form action="/claim" method="POST">
	Address <input type="text" name="address" class="c-field" required>
	<br>
	<input type="submit" value="Claim" class="c-button c-button--info">
	</form>
	`

	err := pageTemplate.Execute(w, template.HTML(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func getCoinsHandler(w http.ResponseWriter, r *http.Request) {
	addr := r.FormValue("address")
	sequence := executeGetSequence(faucet)

	cmd := fmt.Sprintf("gaiacli send --amount=%v --to=%v --name=%v --node=%v --chain-id=%v --sequence=%v", amount, addr, key, node, chain, sequence)
	executeCmd(cmd, pass)

	faucetHandler(w, r)
	return
}
