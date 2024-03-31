package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

var version = "0.0.1"

// NetAddress example: --addr=example.com:60
type NetAddress struct {
	Host string
	Port int
}

func main() {
	addr := new(NetAddress)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Version: %s\n", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// проверка реализации интерфейса
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)

	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}
func (n *NetAddress) Set(value string) error {
	host, port, err := net.SplitHostPort(value)
	if err != nil {
		return err
	}
	n.Host = host
	n.Port, err = strconv.Atoi(port)
	if err != nil {
		return err
	}
	return nil
}
