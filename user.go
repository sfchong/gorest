package gorest

type User struct {
	Id       int
	FullName string
	Email    string
	Age      int
}

type UserService interface {
	GetById(id int) (*User, error)
	Get(filter UserFilter) ([]User, error)
	Create(user User) error
	Update(id int, userUpdate UserUpdate) error
	DeleteById(id int) error
}

type UserFilter struct {
	FullName string
	Email    string
}

type UserUpdate struct {
	FullName string
	Email    string
	Age      int
}
