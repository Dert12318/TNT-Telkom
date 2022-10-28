package test

import (
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/controllers"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"github.com/gin-gonic/gin"
)

var aboutInstance = &controllers.AboutUsController{}

func Engine() *gin.Engine {
	engine := gin.Default()

	engine.GET("/api/v1/about-us", GetAboutUs)

	return engine
}

func GetAboutUs(c *gin.Context) {

	var misiDetail []models.MisiDetail
	var perRekananDetail []models.PerusahaanRekananDetail
	misiDetail = append(misiDetail, models.MisiDetail{
		Item: "Misi 1",
	})
	perRekananDetail = append(perRekananDetail, models.PerusahaanRekananDetail{
		NamaPerusahaan: "Perusahaan 1",
	})

	res := models.AboutUsRequest{
		Profil:            "this is profile",
		Visi:              "this is visi",
		Motto:             "this is motto",
		Office:            "this is office",
		Warehouse:         "this is warehouse",
		Misi:              misiDetail,
		PerusahaanRekanan: perRekananDetail,
		Email:             "this is email",
		NoTelp:            "this is no telp",
		SocialMedia: models.SocialMediaDetail{
			Facebook:  "this is facebook",
			Twitter:   "this is twitter",
			Instagram: "this is instagram",
		},
	}
	controllers.ResponseSuccess(c, res)
}
