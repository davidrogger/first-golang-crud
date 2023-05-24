package controller

import (
	"fmt"
	"net/http"

	"github.com/davidrogger/first-golang-crud/src/configurations/logger"
	"github.com/davidrogger/first-golang-crud/src/configurations/validation"
	"github.com/davidrogger/first-golang-crud/src/controllers/model/request"
	"github.com/davidrogger/first-golang-crud/src/controllers/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	LogField := zapcore.Field{
		Key:    "journey",
		String: "createUser",
	}

	// outra opção zap.String("jorney", "createUser") para definir o campo de log

	logger.Info("Init CreateUser controller", LogField)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		message := fmt.Sprintf("Incorrect Field, error=%s\n", err.Error())
		logger.Error(message, err, LogField)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("User Created successfully", LogField)

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)

}
