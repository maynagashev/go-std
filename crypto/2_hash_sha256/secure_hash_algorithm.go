package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
		"Но после хеширования уже не прочитаешь.")

	// создаём новый hash.Hash, вычисляющий контрольную сумму SHA-256
	h := sha256.New()
	// передаём байты для хеширования
	h.Write(src)
	// вычисляем хеш
	dst := h.Sum(nil)

	fmt.Printf("%x, len=%d", dst, len(dst))
}
