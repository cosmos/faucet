package main

import (
	"encoding/json"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var amountFaucet string
var amountSteak string
var key string
var node string
var chain string
var pass string
var faucet string

type claim_struct struct {
	Address string
}

func main() {
	amountFaucet = "10faucetToken"
	amountSteak = "1steak"

	key = os.Getenv("KEY")
	if key == "" {
		key = "default"
	}

	node = os.Getenv("NODE")
	if node == "" {
		node = "http://localhost:46657"
	}

	chain = os.Getenv("CHAIN")
	if chain == "" {
		chain = "gaia-6001"
	}

	pass = os.Getenv("PASS")
	if pass == "" {
		pass = "1234567890"
	}

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <reCaptcha private key>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {
		recaptcha.Init(os.Args[2])

		http.Handle("/", http.FileServer(http.Dir("./frontend/dist/")))
		http.HandleFunc("/claim", getCoinsHandler)

		if err := http.ListenAndServe("127.0.0,1:9001", nil); err != nil {
			log.Fatal("failed to start server", err)
		}
	}
}

func executeCmd(command string, writes ...string) {
	cmd, wc, _ := goExecute(command)

	for _, write := range writes {
		wc.Write([]byte(write + "\n"))
	}
	cmd.Wait()
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

func processRequest(request *http.Request) (result bool) {
	recaptchaResponse, responseFound := request.Form["g-recaptcha-response"]
	if responseFound {
		result, err := recaptcha.Confirm("127.0.0.1", recaptchaResponse[0])
		if err != nil {
			log.Println("recaptcha server error", err)
		}
		return result
	}
	return false
}

func getCoinsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var claim claim_struct
	err := decoder.Decode(&claim)

	if err != nil {
		panic(err)
	}

	var captcha = processRequest(r)
	addr := claim.Address

	if captcha {
		sendFaucet := fmt.Sprintf(
			"gaiacli send --amount=%v --to=%v --name=%v --chain-id=%v",
			amountFaucet, addr, key, chain)
		fmt.Println(sendFaucet)
		executeCmd(sendFaucet, pass)

		time.Sleep(2 * time.Second)

		sendSteak := fmt.Sprintf(
			"gaiacli send --amount=%v --to=%v --name=%v --chain-id=%v",
			amountSteak, addr, key, chain)
		fmt.Println(sendSteak)
		executeCmd(sendSteak, pass)
	}

	return
}
