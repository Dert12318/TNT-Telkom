package controllers

import (
	"SoalNo6/config/log"
	"SoalNo6/models"
	"SoalNo6/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc   usecase.CartInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewUserController(r *gin.RouterGroup, uc usecase.CartInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &UserController{
		uc:   uc,
		errH: errH,
		logC: logC,
	}

	r.POST("/add-cart", handler.AddCart)
	r.GET("/get-cart", handler.GetAllCart)
	r.DELETE("/delete-cart:id", handler.DeleteCart)
}

func (a UserController) AddCart(c *gin.Context) {
	user := models.Cart{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	// jika butuh validasi

	// fieldErr, err := a.errH.ValidateRequest(user)
	// if err != nil {
	// 	a.logC.Error(err, "controller: Validate request data", "", nil, user, nil)
	// 	c.Error(err).SetMeta(models.ErrMeta{
	// 		ServiceCode: models.ServiceCode,
	// 		FieldErr:    fieldErr,
	// 	})
	// 	c.Abort()
	// 	return
	// }

	err2 := a.uc.AddCart(user)
	if err2 != nil {
		a.logC.Error(err2, "controller: add user usecase", "", nil, user, nil)
		c.Error(err2)
		c.Abort()
		return
	}

	ResponseSuccess(c, nil)
}

func (a UserController) GetAllCart(c *gin.Context) {

	allDataCarts, err := a.uc.GetAllCart()
	if err != nil {
		a.logC.Error(err, "controller: get user usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, allDataCarts)
}

func (a UserController) DeleteCart(c *gin.Context) {
	id := c.Query("id")

	idRes := strings.Split(id, ",")

	err := a.uc.DeleteCart(idRes[1])
	if err != nil {
		a.logC.Error(err, "controller: delete blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, nil)
}

func ResponseSuccess(c *gin.Context, data interface{}) *gin.Context {
	res := models.ResponseSuccess{
		ResponseCode:    "0000",
		ResponseMessage: "success",
		Data:            data,
	}
	c.JSON(200, res)

	return c
}
