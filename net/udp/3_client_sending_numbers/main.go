package main

import (
	"fmt"
	"net"
)

func UDPClient(addr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 1000; i++ {
		_, err := fmt.Fprintf(conn, "%d", i)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	conn.Close()
}

func main() {
	b := make([]byte, 1024)
	addr := &net.UDPAddr{
		Port: 52001,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	go UDPClient(addr)
	var count int
	for {
		_, remoteaddr, err := conn.ReadFromUDP(b)
		if err != nil {
			fmt.Printf("error:  %v\n", err)
			continue
		}
		count++
		if count%1 == 0 {
			fmt.Printf("read from %v, i = %s, count = %d\n", remoteaddr, b, count)
		}
	}
}
