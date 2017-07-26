package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/vshaveyko/h24-static-api/services/encryption"
)

type User struct {
	Base

	EncryptedPassword string `json:"-"`
	Email             string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password          string `json:"password" sql:"-"`
}

// func (u *User) BeforeSave() (err error) {
//
//   if u.Password != "" {
//     u.EncryptedPassword = encryption.Encrypt(u.Password)
//   }
//
//   return
// }
//
func (u *User) IsCorrectPassword(password string) bool {
	return true
	// TODO: change to correct if needed
	// return encryption.CompareHashAndPassword(u.EncryptedPassword, password)
}
