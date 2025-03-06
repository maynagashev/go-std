// Package main получает стандартный вывод команды echo и перенаправляет в команду cat.
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmdout := exec.Command("echo", "Hello, world!")
	stdout, err := cmdout.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmdin := exec.Command("cat")
	cmdin.Stdout = os.Stdout
	cmdin.Stdin = stdout

	if err = cmdout.Start(); err != nil {
		log.Fatal(err)
	}
	if err = cmdin.Start(); err != nil {
		log.Fatal(err)
	}
	cmdin.Wait()
}
