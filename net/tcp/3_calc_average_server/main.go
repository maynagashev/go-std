package main

import (
	"net"
)

const (
	Port   = ":52001" // порт сервера
	MaxLen = 1024     // максимальный размер слайса
)

// handleConn обрабатывает запросы и вычисляет среднее арифметическое.
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		in := make([]byte, MaxLen)
		n, err := c.Read(in)
		if err != nil {
			panic(err)
		}
		var sum int
		for i := 0; i < n; i++ {
			sum += int(in[i])
		}
		if _, err = c.Write([]byte{byte(sum / n)}); err != nil {
			panic(err)
		}
	}
}

// TCPServer запускает сервер и ожидает соединений.
func TCPServer(addr *net.TCPAddr) {
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(c)
	}
}

func main() {
	// Резолвим TCP адрес
	addr, err := net.ResolveTCPAddr("tcp", Port)
	if err != nil {
		panic(err)
	}

	// Выводим информацию о запуске сервера
	println("Запуск сервера на порту", Port)

	// Запускаем TCP сервер
	TCPServer(addr)
}
