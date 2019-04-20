package bookstore

import (
	"errors"
	"fmt"
)

type Bookstore struct {
	Users []*User
	Transactions []*Transaction
}

type Transaction struct {
	buyer *User
	seller *User
	book *Book
	purchasedAmount int
}

func(b *Bookstore) CreateTransaction(buyer, seller *User, book *Book, buyerOfferedAmount int) (transaction Transaction, err error) {

	if buyerOfferedAmount < book.SalePrice  {
		err = errors.New(fmt.Sprintf("error: offer price: %d lower than sale price: %d", buyerOfferedAmount, book.SalePrice))
		return transaction, err
	}
	if buyerOfferedAmount > buyer.DigitalWallet {
		err = errors.New(fmt.Sprintf("error: offer price: %d greater than buyer's wallet balance: %d", buyerOfferedAmount, buyer.DigitalWallet))
		return transaction, err
	}

	transaction.buyer = buyer
	transaction.seller = seller
	transaction.book = book
	transaction.purchasedAmount = buyerOfferedAmount

	b.Transactions = append(b.Transactions, &transaction)



	return transaction, nil
}

func (b *Bookstore) SeedBookStoreData() {

	Bob := User{
		Name: "Bob",
		DigitalWallet: 100,
	}

	Kelly := User{
		Name: "Kelly",
		DigitalWallet: 100,
	}

	crazyTown := Book{
		Title:     "CrazyTown",
		Author:    "Job Bleuthe",
		Owner: &Bob,
		ForSale: true,
		SalePrice: 10,
	}

	joke := Book{
		Title:     "joke",
		Author:    "Joke author",
		Owner:     &Kelly,
		ForSale:   true,
		SalePrice: 10,
	}
	Bob.BooksOwned = append(Bob.BooksOwned, &crazyTown)
	Kelly.BooksOwned = append(Kelly.BooksOwned, &joke)

	b.Users = append(b.Users, &Bob, &Kelly)
}

type BookForSale struct {
	Name string
	Author string
	Owner string
	Price int
}

func(b *Bookstore) ListBooksForSale() (booksForSale map[string]interface{}) {
	booksForSale = make(map[string]interface{})
	for userIdx := range b.Users {
		books := b.Users[userIdx].BooksOwned
		for _, book := range books {
			if book.ForSale {
				bookForSale := make(map[string]interface{})
				bookForSale["Author"] = book.Author
				bookForSale["Owner"] = book.Owner.Name
				bookForSale["Price"] = book.SalePrice
				booksForSale[book.Title] = bookForSale
			}
		}
	}

	return
}

func(b *Bookstore) ListDigitalWallets() (digitalWallets map[string]int) {

	digitalWallets = make(map[string]int)
	for i := range b.Users {
		digitalWallets[b.Users[i].Name] = b.Users[i].DigitalWallet
	}

	return
}