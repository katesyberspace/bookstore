package bookstore

import (
	"github.com/katesyberspace/bookstore/internal/app/books"
	"github.com/katesyberspace/bookstore/internal/app/users"
)

type Bookstore struct {
	Users []*users.User
}


func(b *Bookstore) ListBooksForSale() (booksForSale []books.Book) {

	for i := range b.Users {
		booksForSale = append(booksForSale, b.Users[i].BooksForSale...)
	}

	return
}
