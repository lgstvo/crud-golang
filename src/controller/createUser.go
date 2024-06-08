package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lgstvo/crud-golang/src/configuration/validation"
	"github.com/lgstvo/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
