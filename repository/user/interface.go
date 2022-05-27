package user

import "go-clean/entities"

type UserRepository interface {
	Insert(newUser entities.User) (entities.User, error)
	Update(updateData entities.User) (entities.User, error)
	Delete(id int) error
	GetAll() ([]entities.User, error)
	GetByID(id int) (entities.User, error)
}
