package main

import(
	"fmt"
	"net"
	"time"
	"os/exec"
	"bytes"
	"strings"
)

func log(logs chan string) {
	for true {
		log := <- logs
		fmt.Println(log)
	}
}

func clientMatch() {
	fmt.Println("Client Match")
	path, err := exec.LookPath("SimpleConnect")
	if err != nil {
		fmt.Println("SimpleConnect is not found")
	}
	arg := "-r " + SERVER_ADDR
	fmt.Println("args:", arg)
	cmd := exec.Command(path)
	var out bytes.Buffer
	cmd.Stdout = & out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	localAddr := out.String()
	localPort := ":" + strings.Split(localAddr, ":")[1]
	localPort = strings.TrimSpace(localPort)
	fmt.Println("local port", localPort)
	time.Sleep(time.Microsecond)
	for {
		_ , err = net.Listen(PROTOCOL, localAddr)
		if err == nil {
			break
		}
		fmt.Println(err.Error())
		time.Sleep(2 * time.Second)
	}
}

func clientMatch_old() {
	fmt.Println("Client Match")
	readyChan := make(chan bool)
	addrChan := make(chan string)
	lnChan := make(chan string)
	logs := make(chan string)
	go dialRemote(readyChan, addrChan, logs)
	go connector(readyChan, addrChan, lnChan, logs)
	go listenLocal(lnChan, logs)
	log(logs)
}

func dialRemote(readyChan chan bool, addrChan chan string, logs chan string) {
	logs <- "dial remote"
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err)
	}
	localAddr := conn.LocalAddr().String()
	conn.Close()
	time.Sleep(time.Millisecond)
	logs<- "connected to " + SERVER_ADDR
	logs<- "local: " + localAddr
	readyChan <- true
	addrChan <- localAddr
}

func connector(
	readyChan chan bool,
	addrChan,lnChan chan string,
	logs chan string) {
	logs <- "connector"
	ready := <- readyChan
	if ready {
		logs <- "ready to listen port"
		addr :=  <- addrChan
		logs <- "preparing to connect " + addr
		lnChan <- addr 
	}
	logs <- "connector done"
	return
}

func listenLocal(lnChan chan string, logs chan string) {
	addr := <- lnChan
	logs <- "start listening" + addr
	loop := true
	for loop {
		_ , err := net.Listen(PROTOCOL, addr)
		if err== nil {
			break
		}
		logs <- err.Error()
		logs <- "reconnect in 2 second"
		time.Sleep(time.Second * 2)
	}
}
