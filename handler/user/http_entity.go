package user

import (
	"errors"
	"microservice1/domain/entity"
)

type VerifyUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SaveUserReq struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

func (req VerifyUserReq) validateAndToUserDto() (entity.User, error) {
	if len(req.Email) == 0 || len(req.Password) == 0 {
		return entity.User{}, errors.New("invalid email or password")
	}
	return entity.User{
		Email: req.Email,
	}, nil
}

func (req SaveUserReq) validateAndToUserDto() (entity.User, error) {
	if len(req.Phone) != 10 {
		return entity.User{}, errors.New("invalid phone number")
	}
	return entity.User{
		Email:    req.Email,
		Phone:    req.Phone,
		FullName: req.FullName,
	}, nil
}
