package bookstore

type Book struct {
	Title string
	Author string
	Owner *User
	ForSale bool
	SalePrice int
}

