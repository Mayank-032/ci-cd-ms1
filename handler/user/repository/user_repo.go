package repository

import (
	"context"
	"errors"
	"log"
	"microservice1/database"
	"microservice1/domain/entity"
)

func FetchPassword(ctx context.Context, userPassword string, user entity.User) (string, error) {
	db := database.DB

	sqlQuery := `
		SELECT password FROM users WHERE email = ?
	`
	
	var password string
	err := db.QueryRowContext(ctx, sqlQuery, user.Email).Scan(&password)
	if err != nil {
		log.Println("Error: " + err.Error())
		return "", errors.New("unable_to_make_db_call")
	}
	return password, nil
}

func Save(ctx context.Context, user entity.User) error {
	db := database.DB

	sqlQuery := `
		INSERT INTO user(name, email, phone) 
		VALUES(?, ?, ?)
		ON DUPLICATE KEY UPDATE name=values(name), phone=values(phone)
	`
	
	_, err := db.ExecContext(ctx, sqlQuery, user.FullName, user.Email, user.Phone)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable_to_make_db_call")
	}
	return nil
}