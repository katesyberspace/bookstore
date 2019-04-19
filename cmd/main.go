package main

import (
	"bufio"
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/books"
	"github.com/katesyberspace/bookstore/internal/app/bookstore"
	"github.com/katesyberspace/bookstore/internal/app/users"
	"os"
	"strings"
)

func main(){
	fmt.Println("Welcome to Assembly Bookstore")
	fmt.Println("Type exit to leave the store at any time")

	reader := bufio.NewReader(os.Stdin)

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



	bookstore := bookstore.Bookstore{
		Users: []*users.User{&Bob, &Kelly},
	}

	booksAvailableForPurchase := bookstore.ListBooksForSale()
	fmt.Printf("Books available for purchase:\n%+v", booksAvailableForPurchase)

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