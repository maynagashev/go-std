package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {

	var buf bytes.Buffer

	var mylog = log.New(&buf, `mylog: `, 0)
	mylog.Println("Hello, world!")
	mylog.Println("Goodbye")

	fmt.Print(&buf)
}
