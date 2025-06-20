package grpcserver

import (
	v1 "book_service/gen/v1"
	"book_service/internal/infrastructure/repository"
	"context"
)

type GrpcServer struct {
	v1.UnimplementedBookServiceServer
	repo *repository.PostgresRepository
}

func NewGrpcServer(repo *repository.PostgresRepository) *GrpcServer {
	return &GrpcServer{
		repo: repo,
	}
}

func (s *GrpcServer) AddBook(ctx context.Context, req *v1.BookInfo) (*v1.State, error) {
	// Логика добавления книги
	// Обращение на слой репозитория и
	return &v1.State{}, nil
}

func (s *GrpcServer) Get(ctx context.Context, req *v1.ID) (*v1.BookInfo, error) {
	// Логика получения книги по ID
	// Переход на слой репозитория
	return &v1.BookInfo{}, nil
}

func (s *GrpcServer) ListBook(ctx context.Context, req *v1.Empty) (*v1.BookInfoList), error) {
	// Логика добавления книги
	// Переход на слой репозитория
	return &v1.BookInfoList{}, nil
}
