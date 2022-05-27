package user

import (
	"go-clean/entities"
	"go-clean/repository/user"

	"github.com/go-playground/validator/v10"
)

type serviceUser struct {
	repo       user.UserRepository
	validation *validator.Validate
}

func New(rp user.UserRepository, v *validator.Validate) *serviceUser {
	return &serviceUser{
		repo:       rp,
		validation: v,
	}
}

func (su serviceUser) Add(newData entities.RequestUser) (entities.User, error) {
	res, err := su.repo.Insert(newData.Transform())
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}
func (su *serviceUser) Update(updatedData entities.RequestUser) (entities.User, error) {
	res, err := su.repo.Update(updatedData.Transform())
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}
func (su *serviceUser) Deactive(userID int) error {
	err := su.repo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
func (su *serviceUser) ListUsers() ([]entities.User, error) {
	res, err := su.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (su *serviceUser) MyProfile(userID int) (entities.User, error) {
	res, err := su.repo.GetByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}
