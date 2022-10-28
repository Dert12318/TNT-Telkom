package test

import (
	"SoalNo6/controllers"

	"github.com/gin-gonic/gin"
)

var aboutInstance = &controllers.UserController{}

func Engine() *gin.Engine {
	engine := gin.Default()

	engine.GET("/api/v1/about-us", GetAboutUs)

	return engine
}

func GetAboutUs(c *gin.Context) {

	// var input models.Cart
	// hasil := input{
	// 	input.KodeProduk: "",
	// 	input.NamaProduk: "",
	// 	input.Kuantitas:  "",
	// }

	// res := models.Carts{hasil}
	// controllers.ResponseSuccess(c, res)
}
