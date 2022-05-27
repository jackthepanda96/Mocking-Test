package user

import "go-clean/entities"

type UserService interface {
	Add(newData entities.RequestUser) (entities.User, error)
	Update(updatedData entities.RequestUser) (entities.User, error)
	Deactive(userID int) error
	ListUsers() ([]entities.User, error)
	MyProfile(userID int) (entities.User, error)
}
