package main

import(
	"fmt"
	"net"
)

func log(logs chan string) {
	for true {
		log := <- logs
		fmt.Println(log)
	}
}

func clientMatch() {
	fmt.Println("Client Match")
	readyChan := make(chan bool)
	addrChan := make(chan string)
	lnChan := make(chan string)
	logs := make(chan string)
	go dialRemote(readyChan, addrChan, logs)
	go connector(readyChan, addrChan, lnChan, logs)
}

func dialRemote(readyChan chan bool, addrChan chan string, logs chan string) {
	logs <- "dial remote"
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err)
	}
	localAddr := conn.LocalAddr().String()
	conn.Close()
	logs<- "connected to " + SERVER_ADDR
	logs<- "local: " + localAddr
	addrChan <- localAddr
	readyChan <- true
}

func connector(
	readyChan chan bool,
	addrChan,lnChan chan string,
	logs chan string) {
	ready := <- readyChan
	if ready {
		fmt.Println("ready to listen port")
		logs <- "ready to listen port"
		addr :=  <- addrChan
		logs <- "preparing to connect " + addr
		lnChan <- addr 
	}
	return
}

func listenLocal(lnChan chan string, logs chan string) {
	addr := <- lnChan
	logs <- "start listening" + addr
	_ , err := net.Listen(PROTOCOL, addr)
	if err!= nil {
		panic(err)
	}
}
