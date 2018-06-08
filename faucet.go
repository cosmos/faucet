package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var amount string
var key string
var node string
var chain string
var pass string
var faucet string
var sequence string

type claim_struct struct {
	Address string
}

func main() {
	amount = os.Getenv("AMOUNT")
	if amount == "" {
		amount = "1steak"
	}

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
		chain = "gaia-tardigrade"
	}

	pass = os.Getenv("PASS")
	if pass == "" {
		pass = "1234567890"
	}

	sequence = os.Getenv("SEQUENCE")
	if sequence == "" {
		sequence = "0"
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
	//sequence := executeGetSequence(faucet)

	cmd := fmt.Sprintf("gaiacli send --amount=%v --to=%v --name=%v --chain-id=%v --sequence=%v", amount, addr, key, chain, sequence)
	fmt.Println(cmd)

	executeCmd(cmd, pass)

	i, _ := strconv.Atoi(sequence)
	i++
	sequence = strconv.Itoa(i)

	return
}

/*
func executeGetSequence(addr string) (sequence int64) {
	command := fmt.Sprintf("gaiacli account %v --node=%v --chain-id=%v", addr, node, chain)
	fmt.Println(command)
	cmd := getCmd(command)
	bz, _ := cmd.CombinedOutput()
	out := strings.Trim(string(bz), "\n")
	time.Sleep(time.Second)

	var res map[string]json.RawMessage
	json.Unmarshal([]byte(out), &res)
	fmt.Println(res)

	var value map[string]json.RawMessage
	json.Unmarshal([]byte(res["value"]), &value)
	fmt.Println(value)

	json.Unmarshal([]byte(value["sequence"]), &sequence)
	fmt.Println(sequence)

	return sequence
}
*/
