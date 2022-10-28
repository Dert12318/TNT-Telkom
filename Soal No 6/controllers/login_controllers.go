package controllers

import (
	"CartApp/auth"
	"CartApp/models"
	"CartApp/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	uc usecase.LoginUcInterface
}

func NewLoginController(r *gin.RouterGroup, uc usecase.LoginUcInterface) {
	handler := &LoginController{
		uc: uc,
	}

	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
}

func (a LoginController) Login(c *gin.Context) {
	var encrpytData models.EncryptData
	err := c.ShouldBindJSON(&encrpytData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Ops, Error when bind json from body",
		})
		fmt.Printf("[login Controller] error when encode data enkripsi : %v\n", err)
		return
	}

	token, err := a.uc.LoginUser(encrpytData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success",
		"data":            token,
	})
}

func (a *LoginController) Logout(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Error when unauthorized",
		})
		return
	}

	deleted, errs := a.uc.DeleteAuthData(au.AuthUUID)
	if deleted != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Data not found",
		})
		return
	}
	if errs != nil { //if any goes wrong
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Error when unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Successfully logged out",
	})
}
