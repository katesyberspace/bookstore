package main

import (
	"bufio"
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/bookstore"
	"log"
	"os"
	"strconv"
	"strings"
)

func printWelcomeScreen() {
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("*                  Welcome to Assembly Bookstore                *")
	fmt.Println("*                                                               *")
	fmt.Println("*             Type exit to leave the store at any time          *")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("Press 1 to see books for sale")
	fmt.Println("Press 2 to see users' digital wallet balances")
	fmt.Println("Press 3 to sell a book")
	fmt.Println("Press 4 to buy a book")

}

func printExitScreen() {
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("*                                                               *")
	fmt.Println("*          Thanks For Visiting the Assembly Bookstore           *")
	fmt.Println("*                                                               *")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
}

func printBooksForSale(b *bookstore.Bookstore) {
	booksForSale := b.ListBooksForSale()
	fmt.Printf("Books available for purchase:\n%+v\n", booksForSale)
}

func printUsersDigitalWallets(b *bookstore.Bookstore) {
	usersWalletAccounts := b.ListDigitalWallets()
	fmt.Println("Users digital wallets:")
	for user, balance := range usersWalletAccounts {
		fmt.Printf("User: %s Balance: %d\n", user, balance)
	}
}

func getUserInputFromReader(reader *bufio.Reader, prompt string) (userInput string){
	fmt.Println(">"+prompt)
	fmt.Print("->")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}


func sellBook(b *bookstore.Bookstore, reader *bufio.Reader) {
	var sellerName, bookTitle string
	var salePrice int

	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("*                         Sell A Book                           *")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	sellerName = getUserInputFromReader(reader, "Enter seller's name")
	fmt.Printf("sellerName: %s", sellerName)

	bookTitle = getUserInputFromReader(reader, "Enter book title")
	fmt.Printf("booktitle: %s\n", bookTitle)

	salePriceString := getUserInputFromReader(reader, "Enter sale price")
	salePrice, _ = strconv.Atoi(salePriceString)


	book, _ := b.SellBook(bookTitle, sellerName, salePrice)
	log.Printf(book.Title)
}

func buyBook(b *bookstore.Bookstore) {
	b.BuyBook("Kelly", "Coding For Cats", 20)
}

func main(){
	bookstore := bookstore.Bookstore{}
	bookstore.SeedBookStoreData()
	reader := bufio.NewReader(os.Stdin)
	printWelcomeScreen()

	stayInBookStore := true
	for stayInBookStore  {
		fmt.Print( "->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" || text == "Exit" {
			stayInBookStore = false
		} else if text == "1" {
			printBooksForSale(&bookstore)
		} else if text == "2" {
			printUsersDigitalWallets(&bookstore)
		} else if text == "3" {
			sellBook(&bookstore, reader)
			printBooksForSale(&bookstore)
		} else if text == "4" {
			buyBook(&bookstore)
			printBooksForSale(&bookstore)
		} else {
			fmt.Print("please enter a valid option")
		}
	}
	printExitScreen()

}