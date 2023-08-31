package user

import (
	"log"
	"microservice1/handler/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context) {
	request := VerifyUserReq{}
	response := gin.H{"Status": "false"}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := request.validateAndToUserDto()
	if err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = usecase.ValidateEmailPassword(c, request.Password, user)
	if err != nil {
		log.Println("Error: " + err.Error())
		if err.Error() == "invalid_password" {
			response["message"] = "invalid password"
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response["message"] = "unable to valiadate user due to internal server error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("user validated successfully")
	response["status"] = true
	response["message"] = "successfully validated user"
	c.JSON(http.StatusOK, response)
	return
}

func SaveUser(c *gin.Context) {
	request := SaveUserReq{}
	response := gin.H{"Status": "false"}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := request.validateAndToUserDto()
	if err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = usecase.SaveUserDetails(c, user)
	if err != nil {
		log.Println("Error: " + err.Error())
		if err.Error() == "invalid_payment_details" {
			response["message"] = "payment not initiated"
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response["message"] = "unable to save user due to internal server error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("user saved successfully")
	response["status"] = true
	response["message"] = "successfully saved user"
	c.JSON(http.StatusOK, response)
	return
}