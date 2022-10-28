package application

import (
	"CartApp/config/db"
	"CartApp/config/env"
	log2 "CartApp/config/log"
	"CartApp/controllers"
	"CartApp/middlewares"
	"CartApp/models"
	"CartApp/repo/about_us"
	"CartApp/repo/blog"
	"CartApp/repo/expedition_schedule_rp"
	"CartApp/repo/login_repo"
	"CartApp/repo/user_repo"
	"CartApp/usecase/about_usecase"
	blog2 "CartApp/usecase/blog"
	error2 "CartApp/usecase/error"
	"CartApp/usecase/expedition_schedule_uc"
	"CartApp/usecase/login_usecase"
	"CartApp/usecase/user_usecase"
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
