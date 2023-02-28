package infra

import (
	"database/sql"

	"github.com/ArtuoS/booker-api/internal/entity"
	"github.com/google/uuid"
)

type BookRepository struct {
	Db *sql.DB
}

func (b *BookRepository) GetBooks() ([]entity.Book, error) {
	rows, err := b.Db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Edition, &book.PublicationYear)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) GetBook(id uuid.UUID) (*entity.Book, error) {
	rows, err := b.Db.Query("SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	var book entity.Book
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Name, &book.Edition, &book.PublicationYear)
		if err != nil {
			return nil, err
		}
	}

	return &book, nil
}

func (b *BookRepository) CreateBook(book entity.Book) error {
	query, err := b.Db.Prepare("INSERT INTO books (id, name, edition, publication_year) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = query.Exec(book.Id, book.Name, book.Edition, book.PublicationYear)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) UpdateBook(book *entity.Book) error {
	query, err := b.Db.Prepare("UPDATE books SET name = ?, edition = ?, publication_year = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = query.Exec(book.Name, book.Edition, book.PublicationYear, book.Id)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepository) DeleteBook(id uuid.UUID) error {
	query, err := b.Db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = query.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
