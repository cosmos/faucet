package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var key string
var node string
var chain string
var pass string
var faucet string

type claim_struct struct {
	Address string
}

func main() {
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

	http.Handle("/", http.FileServer(http.Dir("./frontend/dist/")))
	http.HandleFunc("/claim", getCoinsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// only serve to localhost
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, nil))
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

func getCoinsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var claim claim_struct
	err := decoder.Decode(&claim)

	if err != nil {
		panic(err)
	}

	addr := claim.Address

	cmd := fmt.Sprintf("gaiacli send --amount=10faucetToken --to=%v --name=%v --chain-id=%v", addr, key, chain)
	fmt.Println(cmd)
	executeCmd(cmd, pass)

	time.Sleep(2 * time.Second)

	cmdTwo := fmt.Sprintf("gaiacli send --amount=1steak --to=%v --name=%v --chain-id=%v", addr, key, chain)
	fmt.Println(cmdTwo)
	executeCmd(cmdTwo, pass)

	return
}
