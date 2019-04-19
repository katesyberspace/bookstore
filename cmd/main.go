package main

import (
	"bufio"
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/bookstore"
	"os"
	"strings"
)

func main(){
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("*                  Welcome to Assembly Bookstore                *")
	fmt.Println("*                                                               *")
	fmt.Println("*             Type exit to leave the store at any time          *")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)

	bookstore := bookstore.Bookstore{}
	bookstore.SeedBookStoreData()

	booksAvailableForPurchase := bookstore.ListBooksForSale()
	fmt.Printf("Books available for purchase:\n%+v\n", booksAvailableForPurchase)
	usersWalletAccounts := bookstore.ListDigitalWallets()
	fmt.Println("Users digital wallets:")
	for user, balance := range usersWalletAccounts {
		fmt.Printf("User: %s Balance: %d\n", user, balance)

	}

	stayInBookStore := true
	for stayInBookStore  {
		fmt.Print( "->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" || text == "Exit" {
			stayInBookStore = false
		} else {
			fmt.Print(text)
		}
	}

	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("              Thanks for visiting the bookstore!")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")


}