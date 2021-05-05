package mock

import (
	"github.com/sfchong/gorest"
)

type User = gorest.User
type UserFilter = gorest.UserFilter
type UserUpdate = gorest.UserUpdate

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetById(id int) (*User, error) {
	user := User{
		Id:       id,
		FullName: "SF",
		Email:    "sf@gmail.com",
		Age:      26,
	}

	return &user, nil
}

func (s *UserService) Get(filter UserFilter) ([]User, error) {
	users := make([]User, 0)
	users = append(users, User{
		Id:       1,
		FullName: filter.FullName,
		Email:    filter.Email,
		Age:      22,
	})
	users = append(users, User{
		Id:       2,
		FullName: filter.FullName,
		Email:    filter.Email,
		Age:      18,
	})

	return users, nil
}

func (s *UserService) Create(user User) error {
	return nil
}

func (s *UserService) Update(id int, userUpdate UserUpdate) error {
	return nil
}

func (s *UserService) DeleteById(id int) error {
	return nil
}
