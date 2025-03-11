package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "User was successfully created."})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	err = user.Validate()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login successful", "token": token})
}
