package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/crud-golang/src/configuration/logger"
	"github.com/jaquelineabreu/crud-golang/src/configuration/validation"
	"github.com/jaquelineabreu/crud-golang/src/controller/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidationUserErro(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("journey", "createUser"),
	)

	fmt.Println(userRequest)
}
