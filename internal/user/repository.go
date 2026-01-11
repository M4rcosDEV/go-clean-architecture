package user

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id int) (*User, error)
}

// {
// 		FindByID(id int) (*User, error)
// 	Create() (*User, error)
// 	Update(user *User) (*User, error)
// 	Delete(id int) error
// }