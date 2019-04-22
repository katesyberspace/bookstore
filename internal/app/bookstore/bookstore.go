package bookstore

import (
	"errors"
	"fmt"
	"log"
)

type Bookstore struct {
	Users []*User
	Transactions []*Transaction
	BooksForSale map[string]interface{}
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
		Id: 1,
		Name: "Bob",
		DigitalWallet: 100,
		BooksOwned: []*Book{},
	}

	Kelly := User{
		Id: 2,
		Name: "Kelly",
		DigitalWallet: 100,
		BooksOwned: []*Book{},
	}

	crazyTown := Book{
		Title:     "CrazyTown",
		Owner: &Bob,
		ForSale: true,
		SalePrice: 10,
	}

	joke := Book{
		Title:     "joke",
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
				bookForSale["Owner"] = book.Owner.Name
				bookForSale["Price"] = book.SalePrice
				booksForSale[book.Title] = bookForSale
			}
		}
	}
	b.BooksForSale = booksForSale
	return b.BooksForSale
}

func(b *Bookstore) ListDigitalWallets() (digitalWallets map[string]int) {

	digitalWallets = make(map[string]int)
	for i := range b.Users {
		digitalWallets[b.Users[i].Name] = b.Users[i].DigitalWallet
	}

	return
}

func (b *Bookstore) SellBook(bookTitle string, sellerName string, salePrice int) (*Book, error) {
	//find seller
	var seller *User
	for i := range b.Users {
		log.Printf("Iteration: %d, User: %+v", i, b.Users[i].Name)
		if b.Users[i].Name == sellerName {
			seller = b.Users[i]
		} else {
			err := errors.New(fmt.Sprintf("No matching seller in bookstore"))
			log.Printf("error finding seller user: %v", err)
		}
	}
	log.Printf("seller: %+v", seller)
	//create the book
	fmt.Printf("booktitle in method: %s", bookTitle)
	book := Book{
		Title: bookTitle,
		Owner: seller,
		SalePrice: salePrice,
		ForSale: true,
	}
	log.Printf("book: %+v", book)

	//add book to owner's collection
	log.Printf("seller.BooksOwned: %+v", seller.BooksOwned)

	seller.BooksOwned = append(seller.BooksOwned, &book)
	//create the transaction
	txn := Transaction{
		seller: seller,
		book: &book,
	}
	//add to bookstore txns
	b.Transactions = append(b.Transactions, &txn)
	//update books for sale
	b.ListBooksForSale()
	return &book, nil
}

func (b *Bookstore) BuyBook(buyerName string, bookTitle string, offerPrice int) *Book {
	//find buyer
	var buyer *User
	for i := range b.Users {
		if b.Users[i].Name == buyerName {
			buyer = b.Users[i]
		}
	}

	if buyer.DigitalWallet < offerPrice {
		log.Printf("%s wallet balance %d less than offer %d", buyer.Name, buyer.DigitalWallet, offerPrice)
		return nil
	}

	//find book
	var book *Book
	for i := range b.Users {
		for bookIdx := range b.Users[i].BooksOwned {
			if b.Users[i].BooksOwned[bookIdx].Title == bookTitle {
				book = b.Users[i].BooksOwned[bookIdx]
			}
		}
	}
	if offerPrice < book.SalePrice {
		log.Printf("Offer price %d less than sale price %d", offerPrice, book.SalePrice)
		return nil
	}

	var txn *Transaction
	for i := range b.Transactions {
		if b.Transactions[i].book.Title == bookTitle {
			txn = b.Transactions[i]
		}
	}
	txn.buyer = buyer
	txn.purchasedAmount = offerPrice
	book.ForSale = false

	owner := book.Owner
	for i := range owner.BooksOwned {
		if owner.BooksOwned[i].Title == bookTitle {
			owner.BooksOwned = append(owner.BooksOwned[:i], owner.BooksOwned[i+1:]...)
		}
	}
	buyer.BooksOwned = append(buyer.BooksOwned, book)
	owner.DigitalWallet = owner.DigitalWallet + offerPrice
	buyer.DigitalWallet = buyer.DigitalWallet - offerPrice

	return book
}