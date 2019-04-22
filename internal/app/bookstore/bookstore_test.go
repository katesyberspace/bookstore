package bookstore_test

import (
	"bufio"
	"fmt"
	"github.com/katesyberspace/bookstore/internal/app/bookstore"
	"strconv"
	"strings"
)

func printWelcomeScreen() {
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("*                  Welcome to Assembly Bookstore                *")
	fmt.Println("*                                                               *")
	fmt.Println("*             Type exit to leave the store at any time          *")
	fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Press 1 to see books for sale")
	fmt.Println("Press 2 to see users' digital wallet balances")
	fmt.Println("Press 3 to sell a book")

}

func sellBook(b *bookstore.Bookstore, reader *bufio.Reader) {
	var sellerName, bookTitle string
	var salePrice int
	runSellBook := true
	for runSellBook {
		fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
		fmt.Println("*                         Sell A Book                           *")
		fmt.Println("*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*")
		fmt.Println("> Enter seller's name")
		fmt.Print("->")
		sellerName, _ := reader.ReadString('\n')
		sellerName = strings.Replace(sellerName, "\n", "", -1)

		fmt.Println("> Enter book title")
		fmt.Print("->")
		bookTitle, _ := reader.ReadString('\n')
		bookTitle = strings.Replace(bookTitle, "\n", "", -1)

		fmt.Println("> Enter sale price")
		fmt.Print("->")
		salePriceString, _ := reader.ReadString('\n')
		salePriceString = strings.Replace(salePriceString, "\n", "", -1)
		salePrice, _ = strconv.Atoi(salePriceString)

		runSellBook = false
	}

	b.SellBook(sellerName, bookTitle, salePrice)
}

