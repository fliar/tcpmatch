package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"bytes"
	"io"
)

const TITLE = "TCP Make Match"
const SERVER = "Server Mode"
const CLIENT = "Client Mode"
const PRIMARY_PORT = ":8013"
const NAT_PORT = ":8014"
const PROTOCOL = "tcp"
const SERVER_ADDR = "192.168.202.78" + PRIMARY_PORT

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
	conn.Write(([]byte)(NAT_PORT))
}
func clientMain() {
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to ", conn.RemoteAddr().String())
	fmt.Println("I'm connecting on ", conn.LocalAddr().String())
	var buffer bytes.Buffer
	io.Copy(&buffer, conn)
	fmt.Println("Getting:", buffer.String())
	conn.Close()
}
