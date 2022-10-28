package application

import (
	"SoalNo6/config/db"
	"SoalNo6/config/env"
	log2 "SoalNo6/config/log"
	"SoalNo6/controllers"
	"SoalNo6/middlewares"
	"SoalNo6/models"
	"SoalNo6/repo/cart_repo"
	user_usecase "SoalNo6/usecase/cart_usecase"
	error2 "SoalNo6/usecase/error"
	"fmt"
	"log"

	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.New()
	router.Use(gin.Recovery())

	env.NewEnv(".env")

	logCustom := log2.NewLogCustom(env.Config)
	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		log.Println("main : init contract")
	}

	dbBase := db.NewDB(env.Config, false).DB
	fmt.Println(dbBase)
	//dbBase.Debug().Migrator().DropTable(models.ExpeditionSchedule{})
	// err := dbBase.Debug().AutoMigrate(models.ExpeditionSchedule{})
	// if err != nil {
	// 	log.Println("main: cannot auto migrate remapping response code")
	// 	return
	// }

	//dbBase.AutoMigrate(nil)
	models.InitTable(dbBase)

	// init db log
	logDb := log2.NewLogDbCustom(dbBase)
	logCustom.LogDb = logDb

	// repository
	repoUser := cart_repo.NewLoginRepo(dbBase)

	errorUc := error2.NewErrorHandlerUsecase()
	//usecase

	ucUser := user_usecase.NewUserUsecase(repoUser)

	router.Use(middlewares.CORSMiddleware())
	newRoute := router.Group("api/v1")

	//newRoute.Use(middlewares.TokenAuthMiddleware())
	middlewares.NewErrorHandler(newRoute, errorUc, logCustom)

	// controller

	controllers.NewUserController(newRoute, ucUser, errorUc, logCustom)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal("error starting server", err)
	}
}
