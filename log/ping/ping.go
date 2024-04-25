package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

var sugar zap.SugaredLogger

func main() {
	// создаём предустановленный регистратор zap
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic(err)
	}
	defer logger.Sync()

	// делаем регистратор SugaredLogger
	sugar = *logger.Sugar()

	http.Handle("/ping", pingHandler())

	addr := "127.0.0.1:8080"
	// записываем в лог, что сервер запускается
	sugar.Infow(
		"Starting server",
		"addr", addr,
	)
	if err := http.ListenAndServe(addr, nil); err != nil {
		// записываем в лог ошибку, если сервер не запустился
		sugar.Fatalw(err.Error(), "event", "start server")
	}
}

// хендлер для /ping
func pingHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "pong\n")
	}
	return http.HandlerFunc(fn)
}
