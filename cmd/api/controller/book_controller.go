package controller

import (
	"github.com/ArtuoS/booker-api/internal/infra"
)

type BookController struct {
	bookRepository infra.BookRepository
}

func NewBookController(bookService infra.BookRepository) BookController {
	return BookController{
		bookRepository: bookService,
	}
}

/*
func (b *BookController) GetBooks(http.ResponseWriter, *http.Request) ([]entity.Book, error) {
	books, err := b.bookRepository.GetBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookController) GetBook(http.ResponseWriter, *http.Request) (*entity.Book, error) {
	book, err := b.bookRepository.GetBook(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookController) CreateBook(http.ResponseWriter, *http.Request) error {
	err := b.bookRepository.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookController) UpdateBook(http.ResponseWriter, *http.Request) error {
	err := b.bookRepository.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookController) DeleteBook(http.ResponseWriter, *http.Request) error {
	err := b.bookRepository.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
*/
