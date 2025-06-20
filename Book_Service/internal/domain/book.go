package domain

type Book struct {
	ID          string 
	Title       string 
	Author      string 
	Year        int    
}

type BookRepository interface {
	Add(book Book) string
	Get(id int) Book 
	GetAll() []Book
}

func NewBook(id string, title string, author string, year int) Book {
	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
	}
}

