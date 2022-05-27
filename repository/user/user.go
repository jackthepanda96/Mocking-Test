package user

import (
	"go-clean/entities"
	"log"

	"gorm.io/gorm"
)

type repoUser struct {
	db *gorm.DB
}

func New(conn *gorm.DB) *repoUser {
	return &repoUser{
		db: conn,
	}
}

func (ru *repoUser) Insert(newUser entities.User) (entities.User, error) {
	err := ru.db.Create(&newUser).Error
	if err != nil {
		log.Fatal(err)
		return entities.User{}, err
	}

	return newUser, nil
}

func (ru *repoUser) Update(updateData entities.User) (entities.User, error) {
	err := ru.db.Save(&updateData).Error
	if err != nil {
		return entities.User{}, err
	}
	return updateData, nil
}
func (ru *repoUser) Delete(id int) error {
	err := ru.db.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (ru *repoUser) GetAll() ([]entities.User, error) {
	var res []entities.User

	err := ru.db.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (ru *repoUser) GetByID(id int) (entities.User, error) {
	var res entities.User

	err := ru.db.First(&res, id).Error
	if err != nil {
		return entities.User{}, err
	}

	return res, nil
}
