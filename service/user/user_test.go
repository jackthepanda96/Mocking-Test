package user_test

import (
	"errors"
	"go-clean/entities"
	repoMock "go-clean/mocks"
	"go-clean/service/user"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAdd(t *testing.T) {
	repo := new(repoMock.UserRepository)
	insertData := entities.RequestUser{Nama: "jerry", NoHP: "1234", Alamat: "surabaya", Password: "1234"}
	returnData := entities.User{Model: gorm.Model{ID: uint(1)}, Nama: "jerry", NoHP: "1234", Alamat: "surabaya", Password: "1234"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := user.New(repo, validator.New())

		res, err := srv.Add(insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.Nama, res.Nama)
		assert.Equal(t, returnData.Alamat, res.Alamat)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(entities.User{}, errors.New("there is some error")).Once()
		srv := user.New(repo, validator.New())

		res, err := srv.Add(entities.RequestUser{})
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Nama)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}

func TestListUsers(t *testing.T) {
	repo := new(repoMock.UserRepository)
	returnData := []entities.User{{Model: gorm.Model{ID: uint(1)}, Nama: "jerry", NoHP: "1234", Alamat: "surabaya", Password: "1234"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()

		srv := user.New(repo, validator.New())

		res, err := srv.ListUsers()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("data not found")).Once()

		srv := user.New(repo, validator.New())

		res, err := srv.ListUsers()
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestMyProfile(t *testing.T) {
	repo := new(repoMock.UserRepository)
	returnData := entities.User{Model: gorm.Model{ID: uint(1)}, Nama: "jerry", NoHP: "1234", Alamat: "surabaya", Password: "1234"}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetByID", 1).Return(returnData, nil).Once()

		srv := user.New(repo, validator.New())

		_, err := srv.MyProfile(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Out of Index", func(t *testing.T) {
		repo.On("GetByID", 2).Return(entities.User{}, errors.New("error out of index")).Once()

		srv := user.New(repo, validator.New())

		_, err := srv.MyProfile(2)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})
}
