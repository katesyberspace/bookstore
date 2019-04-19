package users

import "github.com/katesyberspace/bookstore/internal/app/books"

type User struct {
	Name string
	BooksForSale []books.Book
	BooksOwned []books.Book
	DigitalWallet int
}

