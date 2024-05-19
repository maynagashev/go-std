package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	b, _ := RandBytes(10)
	fmt.Printf("%#v", b)
}

func RandBytes(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ``, err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
