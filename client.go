package main

import(
	"fmt"
	"net"
)

func clientMatch() {
	fmt.Println("Client Match")
	readyChan := make(chan bool)
	addrChan := make(chan string)
	lnChan := make(chan string)
	go dialRemote(readyChan, addrChan)
	go connector(readyChan, addrChan, lnChan)
}

func dialRemote(readyChan chan bool, addrChan chan string) {
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err)
	}
	localAddr := conn.LocalAddr().String()
	conn.Close()
	fmt.Println("connected to ", SERVER_ADDR)
	fmt.Println("local: ", localAddr)
	addrChan <- localAddr
	readyChan <- true
}

func connector(readyChan chan bool, addrChan,lnChan chan string) {
	ready := <- readyChan
	if ready {
		fmt.Println("ready to listen port")
		addr :=  <- addrChan
		fmt.Println("preparing to connect ", addr)
		lnChan <- addr 
	}
	return
}

func listenLocal(lnChan chan string) {
	addr := <- lnChan
	fmt.Println("start listening", addr)
	_ , err := net.Listen(PROTOCOL, addr)
	if err!= nil {
		panic(err)
	}
}
