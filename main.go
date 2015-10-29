package main
import(
	"fmt"
	"net"
	"flag"
	"strings"
)
const TITLE 		= "TCP Make Match"
const SERVER 		= "Server Mode"
const CLIENT		= "Client Mode"
const PORT			= ":8013"
const PROTOCOL		= "tcp"
const SERVER_ADDR	= "192.168.202.78" + PORT

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
	ln, err := net.Listen(PROTOCOL, PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listenning ", PORT)
	for i := 0; i < 5; i++ {
		conn , err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Connection ", conn.RemoteAddr().String() ,"Accepted")
	}
}
func clientMain() {
	conn, err := net.Dial(PROTOCOL, SERVER_ADDR)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to ", conn.RemoteAddr().String())
	conn.Close()
}
