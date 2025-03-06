// Package main получает стандартный вывод команды echo и перенаправляет в команду cat.
package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	cmdout := exec.Command("echo", "Hello, world!")
	stdout, err := cmdout.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmdin := exec.Command("cat")
	// указываем текущую консоль для стандартного вывода
	cmdin.Stdout = os.Stdout
	stdin, err := cmdin.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	if err = cmdout.Start(); err != nil {
		log.Fatal(err)
	}
	if err = cmdin.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		// перенаправляем потоки данных
		if _, err = io.Copy(stdin, stdout); err != nil {
			log.Fatal(err)
		}
		wg.Done()
		// закрываем, чтобы завершился процесс cat
		stdin.Close()
	}()
	wg.Wait()
	cmdout.Wait()
	cmdin.Wait()
}
