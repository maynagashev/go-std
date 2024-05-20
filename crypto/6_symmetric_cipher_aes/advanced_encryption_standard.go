package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
)

func generateRandom(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

/*
Зашифруем и расшифруем текст с помощью алгоритма AES (Advanced Encryption Standard).
Это блочный алгоритм, размер блока — 16 байт.

Для работы алгоритма нужно сгенерировать ключ из 16, 24 или 32 байт.
В зависимости от размера ключа будет выбрано шифрование AES-128, AES-192 или AES-256.
Чем длиннее ключ, тем более криптостойким получится шифр.
*/
func main() {
	src := []byte("Слепой банкир") // данные, которые хотим зашифровать
	fmt.Printf("original: %s\n", src)

	// константа aes.BlockSize определяет размер блока, она равна 16 байтам
	key, err := generateRandom(aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// получаем cipher.Block
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	dst := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(dst, src)
	fmt.Printf("encrypted: %x\n", dst)

	src2 := make([]byte, aes.BlockSize) // расшифровываем
	aesblock.Decrypt(src2, dst)
	fmt.Printf("decrypted: %s\n", src2)
}
