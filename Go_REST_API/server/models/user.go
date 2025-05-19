package models

import (
    "example.com/testserver/db"
	"example.com/testserver/utils"
	"errors"
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
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(hashedPassword, u.Email)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userID
	u.Password = hashedPassword
	return nil
}



func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"	
	row := db.DB.QueryRow(query, u.Email)

	var storedPassword string
	err := row.Scan(&u.ID, &storedPassword)

	if err != nil {
		return err
	}	

	passwordIsValid := utils.CheckPasswordHash(u.Password, storedPassword)
	if !passwordIsValid {
		return errors.New("Credentials are not valid")
	}
	
	return nil
}


