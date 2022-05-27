package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string
	Alamat   string
	NoHP     string
	Password string
}

/* Request Side*/
type RequestUser struct {
	Nama     string `json:"nama" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	NoHP     string `json:"nohp" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (ru RequestUser) Transform() User {
	return User{
		Nama:     ru.Nama,
		Alamat:   ru.Alamat,
		NoHP:     ru.NoHP,
		Password: ru.Password,
	}
}

type RequestLogin struct {
	NoHP     string `json:"nohp" validate:"required"`
	Password string `json:"password" validate:"required"`
}

/* Response Side*/
