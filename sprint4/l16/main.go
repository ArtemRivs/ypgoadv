package main // объявление пакета main

import (
	"context" // подключение пакета context
	"log" // подключение пакета log
	"net/http" // подключение пакета net/http
	"os" // подключение пакета os
	"os/signal" // подключение пакета os/signal
)

func main() { // объявление функции main

	// создание объекта http.Server
	var srv http.Server

	// создание канала для сигнала завершения работы сервера
	idleConnsClosed := make(chan struct{})

	// запуск анонимной горутины
	go func() {
		// создание канала для сигналов прерывания
		sigint := make(chan os.Signal, 1)
		// регистрация сигналов прерывания
		signal.Notify(sigint, os.Interrupt)
		// ожидание сигнала прерывания
		<-sigint

		// остановка сервера и обработка ошибки
		if err := srv.Shutdown(context.Background()); err != nil {
			// вывод ошибки в лог
			log.Printf("HTTP server Shutdown: %v", err)
		}

		// закрытие канала для сигнала завершения работы сервера
		close(idleConnsClosed)
	}()

	// запуск сервера и обработка ошибки
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// вывод ошибки в лог и завершение работы программы
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	// ожидание завершения работы сервера
	<-idleConnsClosed
}
