package repositories

import (
	"errors"

	"github.com/mariusfa/golf/middleware"
)

type UserRepository struct {
	users []middleware.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: []middleware.User{}}
}

func (r *UserRepository) AddUser(user middleware.User) {
	r.users = append(r.users, user)
}

func (r *UserRepository) FindById(id string) (*middleware.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
