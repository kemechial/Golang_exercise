package models

import (
    "example.com/testserver/db"
)

type User struct {
	ID       int64  `json:"id"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (u *User) Save()  error {
	query := "INSERT INTO users (password, email) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Password, u.Email)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userID
	return nil
}

