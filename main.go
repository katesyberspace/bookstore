package main

import (
	"bufio"
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/bookstore"
	"os"
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
		} else {
			fmt.Print("please enter a valid option")
		}
	}
	printExitScreen()

}