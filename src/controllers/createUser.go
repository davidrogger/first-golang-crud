package controller

import (
	"fmt"
	"log"

	"github.com/davidrogger/first-golang-crud/src/configurations/rest_err"
	"github.com/davidrogger/first-golang-crud/src/controllers/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		message := fmt.Sprintf("Incorrect Field, error=%s\n", err.Error())
		log.Print(message)
		restErr := rest_err.NewBadRequestError(message)
		c.JSON(restErr.Code, restErr)
		return
	}

}
