package main

import (
	"book_service/config"
	grpcserver "book_service/internal/infrastructure/grpc_server"
	"book_service/internal/infrastructure/repository"
	"log"
)

func main() {
	// Инициализация конфига
	cfg,err := config.LoadConfig()
	if err != nil{
		log.Fatal("Error load config: ",err)
	}
	
	// Инициализируем сервер на порту()
	
	// Проверка подключение к БД
	
	// Инициализация хранилища
	repo := repository.NewPostgresRepository()
	
	// Слой инфаструктуры
	gserver := grpcserver.NewGrpcServer(repo)

	// Инзиициалируем grpc сервер

	// Регистрируем для grpc сервера, те методы которые которые должен реализовать
	// т.е мы прикрепляем для grpc сервера слой инфаструктуры (т.е логику вызова gRPC-методов)

	// Запускаем и слушаем grpc сервер

}
