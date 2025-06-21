package main

import (
	"book_service/config"
	v1 "book_service/gen/v1"
	grpcserver "book_service/internal/infrastructure/grpc_server"
	"book_service/internal/infrastructure/repository"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Инициализация конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error load config: ", err)
	}

	// Формирование строки подключения
	connect_string := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Name + "?sslmode=" + cfg.SslMode
	// Создание БД
	DB, err := repository.InitDB(connect_string)
	if err != nil {
		log.Fatal("Error create DB ", err)
	}
	defer DB.Close()

	// Проверка подключение к БД
	if err := repository.Ping(DB); err != nil {
		log.Fatal("Error connect to DB: ", err)
	}

	// Инициализация хранилища
	repo := repository.NewPostgresRepository(DB)

	// Слой инфаструктуры
	gserver := grpcserver.NewGrpcServer(repo)

	// Инициализируем сервер на порту
	listener, err := net.Listen("tcp", "0.0.0.0:8193")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
	defer listener.Close()

	// Инзиициалируем grpc сервер
	grpcserver := grpc.NewServer()

	// Регистрируем для grpc сервера, те методы которые которые должен реализовать
	// т.е мы прикрепляем для grpc сервера слой инфаструктуры (т.е логику вызова gRPC-методов)
	v1.RegisterBookServiceServer(grpcserver, gserver)

	log.Printf("server listening at %v", listener.Addr())

	// Запускаем и слушаем grpc сервер
	if err := grpcserver.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// Запускаем и слушаем grpc сервер

}
