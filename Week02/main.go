package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	ID 		string
	name 	string
}

var DB *sql.DB

func (u User) getUserById(ID string) (User, error) {
	err := DB.Table("user").Where("id = ?", ID).Find(u).Error
	if err != nil {
		return errors.Wrap(err, "record not found")
	}
	return u, nil
}