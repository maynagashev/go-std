package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// Создание клиента с отключенным сжатием
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	// Отправка GET запроса
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	// Отображение тела ответа
	fmt.Println("Ответ от сервера:")
	fmt.Println(string(body))
}
