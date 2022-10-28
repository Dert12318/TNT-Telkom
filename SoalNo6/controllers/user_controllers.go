package controllers

import (
	"strconv"
	"strings"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc   usecase.UserUcInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewUserController(r *gin.RouterGroup, uc usecase.UserUcInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &UserController{
		uc:   uc,
		errH: errH,
		logC: logC,
	}

	r.POST("/add-user", handler.AddUser)
	r.GET("/all-user", handler.GetAllUser)
	r.GET("/user/:id", handler.GetUserByID)
	r.PUT("/update-user/:id", handler.UpdateUser)
	r.DELETE("/delete-user", handler.DeleteUser)
}

func (a UserController) AddUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := a.errH.ValidateRequest(user)
	if err != nil {
		a.logC.Error(err, "controller: Validate request data", "", nil, user, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	result, err := a.uc.AddUser(user)
	if err != nil {
		a.logC.Error(err, "controller: add user usecase", "", nil, user, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
}

func (a UserController) GetAllUser(c *gin.Context) {

	allDataUser, err := a.uc.GetAll()
	if err != nil {
		a.logC.Error(err, "controller: get user usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, allDataUser)
}

func (a UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	data, err := a.uc.GetById(id)
	if err != nil {
		a.logC.Error(err, "controller: get blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, data)
}

func (a UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	err = c.ShouldBindJSON(&user)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := a.errH.ValidateRequest(user)
	if err != nil {
		a.logC.Error(err, "controller: Validate request data", "", nil, user, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	data, err := a.uc.UpdateData(id, user)
	if err != nil {
		a.logC.Error(err, "controller: update blog usecase", "", nil, user, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, data)
}

func (a UserController) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	idRes := strings.Split(id, ",")

	err := a.uc.DeleteData(idRes)
	if err != nil {
		a.logC.Error(err, "controller: delete blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, nil)
}
