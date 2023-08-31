package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"microservice1/config"
	"microservice1/domain"
	"microservice1/response"
	"net/http"
)

func VerifyUserPaymentDetails(ctx context.Context, email string) (bool, error) {
	body := map[string]string {
		"email": email,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println("Error: " + err.Error())
		return false, errors.New("unable_to_unmarshal_body")
	}

	basePath := config.Configurations.BaseURL
	request, err := http.NewRequest("POST", fmt.Sprintf("%v/payment/verify", basePath), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Println("Error: " + err.Error())
		return false, errors.New("unable_to_create_request")
	}

	client := http.Client {
		Timeout: domain.InternalTimeout,
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Println("Error: " + err.Error())
		return false, errors.New("unable_to_network_call")
	}
	defer resp.Body.Close()

	respBytes, err := json.Marshal(resp.Body)
	if err != nil {
		log.Println("Error: " + err.Error())
		return false, errors.New("unable_to_marshal_response")
	}

	var paymentResponse response.VerifyPaymentResponse
	err = json.Unmarshal(respBytes, &paymentResponse)
	if err != nil {
		log.Println("Error: " + err.Error())
		return false, errors.New("unable_to_unmarshal_response_bytes")
	}
	return paymentResponse.IsPaymentVerified, nil
}