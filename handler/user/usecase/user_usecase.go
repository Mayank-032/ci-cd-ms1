package usecase

import (
	"context"
	"errors"
	"log"
	"microservice1/domain/entity"
	"microservice1/handler/user/repository"
)

func ValidateEmailPassword(ctx context.Context, password string, user entity.User) error {
	fetchedPassword, err := repository.FetchPassword(ctx, password, user)
	if err != nil {
		return errors.New("unable_to_fetch_password")
	}

	if password != fetchedPassword {
		log.Println("Error: password not matched")
		return errors.New("invalid_password")
	}
	return errors.New("unable_to_validate_user")
}

func SaveUserDetails(ctx context.Context, user entity.User) error {
	isPaymentVerified, err := repository.VerifyUserPaymentDetails(ctx, user.Email)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable_to_verify_payment_details")
	}

	if !isPaymentVerified {
		log.Println("Error: payment_details_not_verified")
		return errors.New("inavalid_payment_details")
	}

	err = repository.Save(ctx, user)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable_save_user")
	}
	return errors.New("unable to ")
}