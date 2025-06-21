package repository

import (
	"book_service/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Add добавляет новую книгу в базу данных и возвращает ID созданной записи
func (r *PostgresRepository) Add(book *domain.Book) string {
	query := `INSERT INTO Learn_Library (Title, Author, Year) VALUES ($1, $2, $3,$4);`

	var id int
	err := r.db.QueryRowContext(
		context.Background(),
		query,
		book.Title,
		book.Author,
		book.Year,
	).Scan(&id)

	if err != nil {
		return fmt.Sprintf("Error exucute query add user: %v", err)
	}

	return "Succesful!"
}

// Get возвращает книгу по ID
func (r *PostgresRepository) Get(id int) *domain.Book {
	query := `SELECT id, Title, Author, Year FROM Learn_Library WHERE id = $1;`

	var book domain.Book
	err := r.db.QueryRowContext(
		context.Background(),
		query,
		id,
	).Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Book{}
		}
		return &domain.Book{}
	}

	return &book
}

// ListBook возвращает список всех книг
func (r *PostgresRepository) ListBook() []domain.Book {
	query := `SELECT id, Title, Author, Year FROM Learn_Library ORDER BY id;`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return []domain.Book{}
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			return []domain.Book{}
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return []domain.Book{}
	}

	return books
}
