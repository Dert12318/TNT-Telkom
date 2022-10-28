package application

import (
	"fmt"
	"log"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/db"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/env"
	log2 "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/controllers"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/middlewares"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo/about_us"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo/blog"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo/expedition_schedule_rp"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo/login_repo"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo/user_repo"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/about_usecase"
	blog2 "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/blog"
	error2 "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/error"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/expedition_schedule_uc"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/login_usecase"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase/user_usecase"

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
	err := dbBase.Debug().AutoMigrate(models.ExpeditionSchedule{})
	if err != nil {
		log.Println("main: cannot auto migrate remapping response code")
		return
	}
	//dbBase.AutoMigrate(nil)
	models.InitTable(dbBase)

	// init db log
	logDb := log2.NewLogDbCustom(dbBase)
	logCustom.LogDb = logDb

	// repository
	repoUser := user_repo.NewUserRepo(dbBase)
	repoLogin := login_repo.NewLoginRepo(dbBase)
	aboutRepo := about_us.NewAboutUsRepo(dbBase, logCustom)
	esRepo := expedition_schedule_rp.NewExpeditionRepo(dbBase, logCustom)
	blogRepo := blog.NewBlogRepo(dbBase, logCustom)

	errorUc := error2.NewErrorHandlerUsecase()
	//usecase

	ucLogin := login_usecase.NewLoginUsecase(repoLogin)
	ucUser := user_usecase.NewUserUsecase(repoUser)
	abtUc := about_usecase.NewAboutUsUsecase(aboutRepo, logCustom)
	esUc := expedition_schedule_uc.NewEsUc(esRepo, logCustom)
	blogUc := blog2.NewBlogUc(blogRepo, logCustom)

	router.Use(middlewares.CORSMiddleware())
	newRoute := router.Group("api/v1")

	//newRoute.Use(middlewares.TokenAuthMiddleware())
	middlewares.NewErrorHandler(newRoute, errorUc, logCustom)

	// controller

	controllers.NewUserController(newRoute, ucUser, errorUc, logCustom)
	controllers.NewAboutUsController(newRoute, abtUc, errorUc, logCustom)
	controllers.NewExpeditionController(newRoute, esUc, errorUc, logCustom)
	controllers.NewBlogController(newRoute, blogUc, errorUc, logCustom)
	controllers.NewDaerahController(newRoute, dbBase)
	router.Use(middlewares.TokenAuthMiddlewareCustom(repoLogin))
	controllers.NewLoginController(newRoute, ucLogin)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal("error starting server", err)
	}
}
