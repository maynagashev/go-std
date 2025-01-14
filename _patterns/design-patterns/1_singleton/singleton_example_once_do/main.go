package main

import (
	"fmt"
	"os"
	"sync"
)

var dumpFile *os.File
var dumpOnce sync.Once

func Dump(data []byte) error {
	fmt.Printf("Attempting to dump data: %s\n", string(data))

	dumpOnce.Do(func() {
		fmt.Println("Initializing dump file...")
		// Создаем файл в текущей директории
		var err error
		dumpFile, err = os.Create("output.dump")
		if err != nil {
			fmt.Printf("Error creating dump file: %v\n", err)
			panic(err)
		}
		fmt.Println("Dump file created successfully")
	})

	n, err := dumpFile.Write(data)
	fmt.Printf("Wrote %d bytes to dump file\n", n)
	return err
}

func main() {
	defer dumpFile.Close()
	Dump([]byte("Hello, World!"))
}
