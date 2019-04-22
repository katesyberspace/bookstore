package bookstore

type Book struct {
	Title string
	Owner *User
	ForSale bool
	SalePrice int
}
