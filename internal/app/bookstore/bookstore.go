package bookstore

import (
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/books"
	"github.com/katesyberspace/bookstore/internal/app/users"
)

type Bookstore struct {
	Users []*users.User
}

func (b *Bookstore) SeedBookStoreData() {
	crazyTown := books.Book{
		Title: "CrazyTown",
		Author: "Job Bleuthe",
		SalePrice: 10,
	}

	Bob := users.User{
		Name: "Bob",
		BooksForSale: []books.Book{crazyTown},
		DigitalWallet: 100,
	}

	Kelly := users.User{
		Name: "Kelly",
		DigitalWallet: 100,
	}

	b.Users = append(b.Users, &Bob, &Kelly)
}


func(b *Bookstore) ListBooksForSale() (booksForSale []books.Book) {

	for i := range b.Users {
		booksForSale = append(booksForSale, b.Users[i].BooksForSale...)
	}

	return
}

func(b *Bookstore) ListDigitalWallets() (digitalWallets map[string]int) {
	digitalWallets = make(map[string]int)
	for i := range b.Users {
		digitalWallets[b.Users[i].Name] = b.Users[i].DigitalWallet
	}

	for user, balance := range digitalWallets {
		fmt.Sprintf("User: %s Balance: %d\n", user, balance)
	}
	return
}

func(b *Bookstore) CreateTransaction() {



}