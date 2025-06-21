package grpcserver

import (
	v1 "book_service/gen/v1"
	"book_service/internal/domain"
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

func (s *GrpcServer) Add(ctx context.Context, req *v1.BookInfo) (*v1.State, error) {
	select {
	case <-ctx.Done():
		return &v1.State{
			Status: "Request timeout",
		}, nil
	default:
		user := &domain.Book{
			Title:  req.Title,
			Author: req.Author,
			Year:   int(req.Year),
		}

		// Возвращаем состояние
		status := s.repo.Add(user)
		return &v1.State{
			Status: status,
		}, nil
	}
}

func (s *GrpcServer) Get(ctx context.Context, req *v1.ID) (*v1.BookInfo, error) {
	select {
	case <-ctx.Done():
		return &v1.BookInfo{}, nil
	default:
		userGetting := s.repo.Get(int(req.Id))
		return &v1.BookInfo{
			Title:  userGetting.Title,
			Author: userGetting.Author,
			Year:   int32(userGetting.Year),
		}, nil
	}
}

func (s *GrpcServer) ListBook(ctx context.Context, req *v1.Empty) (*v1.BookInfoList, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		users := s.repo.ListBook()
		var bookInfos []*v1.BookInfo
		for _, user := range users {
			bookInfos = append(bookInfos, &v1.BookInfo{
				Title:  user.Title,
				Author: user.Author,
				Year:   int32(user.Year),
			})
		}
		return &v1.BookInfoList{
			Books: bookInfos,
		}, nil
	}

}
