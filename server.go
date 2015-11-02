package main
import(
	"fmt"
	"net"
	"time"
)

func listen() {
	ln, err := net.Listen(PROTOCOL, PRIMARY_PORT)
	if err != nil {
		panic(err.Error())
	}
	conns := make(map[string] string)
	for i := 0; i < 5; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err.Error())
		}
		name := fmt.Sprint("Client", i+1)
		fmt.Println(name, "connected")
		conns[name] = conn.RemoteAddr().String()
		conn.Close()
		time.Sleep(time.Millisecond)
	}
}
