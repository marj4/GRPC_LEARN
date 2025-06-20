package repository

import (
	"book_service/internal/domain"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Add(book *domain.Book) string {
	// Логика добавления книги в базу данных
	return ""
}

func (r *PostgresRepository) Get(id int) *domain.Book {
	// Логика добавления книги в базу данных
	return &domain.Book{}
}

func (r *PostgresRepository) ListBook() []domain.Book {
	// Логика добавления книги в базу данных
	return []domain.Book{}
}
