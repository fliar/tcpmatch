package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"bufio"
)

const TITLE = "TCP Make Match"
const SERVER = "Server Mode"
const CLIENT = "Client Mode"
const PRIMARY_PORT = ":8013"
const NAT_PORT = ":8014"
const PROTOCOL = "tcp"
const SERVER_IP = "192.168.202.78"
const SERVER_ADDR = SERVER_IP + PRIMARY_PORT
const DELIM = '\x03'
func main() {
	modeArg := flag.String("m", "c", "running mode")
	flag.Parse()
	mode := strings.ToLower(*modeArg)
	fmt.Println(TITLE)
	if mode == "s" || mode == "server" {
		fmt.Println("running mode: ", mode)
		serverMain()
		return
	}
	if mode == "c" || mode == "client" {
		clientMain()
		return
	}
	fmt.Println("unknown mode")
}

func serverMain() {
	ln, err := net.Listen(PROTOCOL, PRIMARY_PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listenning ", PRIMARY_PORT)
	tln, err := net.Listen(PROTOCOL, NAT_PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listenning ", NAT_PORT)
	for i := 0; i < 5; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handlePrimaryConnection(conn)
		conn, err = tln.Accept()
		if err != nil {
			panic(err)
		}
		
	}
}
func handleNATConnection(conn net.Conn) {
	fmt.Println("NAT Connection ", conn.RemoteAddr().String(), "Accepted")
}
func handlePrimaryConnection(conn net.Conn) {
	fmt.Println("Primary Connection ", conn.RemoteAddr().String(), "Accepted")
	writer := bufio.NewWriter(conn)
	writer.WriteString(NAT_PORT)
	writer.WriteByte(DELIM)
	writer.Flush()
}
func clientMain() {
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to ", conn.RemoteAddr().String())
	fmt.Println("I'm connecting on ", conn.LocalAddr().String())
	reader := bufio.NewReader(conn)
	portStr , _ := reader.ReadString(DELIM)
	fmt.Println("Getting:", portStr)
	//connect to NAT port
	connNAT, err := net.Dial(PROTOCOL, SERVER_IP + portStr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to ", connNAT.RemoteAddr().String())
	fmt.Println("I'm connecting on ", connNAT.LocalAddr().String())
}
