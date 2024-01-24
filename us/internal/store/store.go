package store

type Store interface {
	GetItems() ([]Item, error)
	AddItem(item Item) error
	GetItemByID(id string) (*Item, error)
	GetItemByName(Name string) (*Item, error)
	UpdateItem(id string, updatedItem Item) error
	DeleteItem(id string) error

	/*GetUsers() ([]User, error)
	GetUserByID(userID int) (*User, error)
	AddUser(user User) error
	AddItemToUserCart(userID int, item Item) error
	DeleteUser(userID int) error
	AddUserRole(userID int, role string) error*/
}
