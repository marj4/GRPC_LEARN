package main

import (
	grpcserver "book_service/internal/infrastructure/grpc_server"
	"book_service/internal/infrastructure/repository"
)

func main() {
	// Инициализируем сервер на порту()

	// Инициализация хранилища
	repo := repository.NewPostgresRepository()

	// Слой инфаструктуры
	gserver := grpcserver.NewGrpcServer(repo)

	// Инзиициалируем grpc сервер

	// Регистрируем для grpc сервера, те методы которые которые должен реализовать
	// т.е мы прикрепляем для grpc сервера слой инфаструктуры (т.е логику вызова gRPC-методов)

	// Запускаем и слушаем grpc сервер

}
