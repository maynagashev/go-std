package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var srv = http.Server{Addr: ":8089"}
	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})
	// канал для перенаправления прерываний
	// поскольку нужно отловить всего одно прерывание,
	// ёмкости 1 для канала будет достаточно
	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	// запускаем горутину обработки пойманных прерываний
	go func() {
		// читаем из канала прерываний
		// поскольку нужно прочитать только одно прерывание,
		// можно обойтись без цикла
		<-sigint
		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := srv.Shutdown(context.Background()); err != nil {
			// ошибки закрытия Listener
			log.Printf("HTTP server Shutdown: %v", err)
		}
		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// ошибки старта или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	// ждём завершения процедуры graceful shutdown
	<-idleConnsClosed
	// получили оповещение о завершении
	// здесь можно освобождать ресурсы перед выходом,
	// например закрыть соединение с базой данных,
	// закрыть открытые файлы
	fmt.Println("Server Shutdown gracefully")
}
