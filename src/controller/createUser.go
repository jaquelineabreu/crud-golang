package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/crud-golang/src/configuration/validation"
	"github.com/jaquelineabreu/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser Controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s/n", err.Error())
		restErr := validation.ValidationUserErro(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
