package controllers

import (
	"errors"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models/contract"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase"

	"github.com/gin-gonic/gin"
)

type AboutUsController struct {
	uc   usecase.AboutUsUcInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewAboutUsController(r *gin.RouterGroup, uc usecase.AboutUsUcInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &AboutUsController{
		uc:   uc,
		errH: errH,
		logC: logC,
	}

	r.POST("/add-about-us", handler.AddAboutUs)
	r.GET("/about-us", handler.GetAboutUs)
	r.PUT("/update-about-us", handler.UpdateAboutUs)
	r.DELETE("/delete-about-us", handler.DeleteAboutUs)
	r.GET("/about", handler.GetAboutUsById)
}

func (a AboutUsController) GetAboutUs(c *gin.Context) {
	aboutUsList, err := a.uc.GetAboutUs()
	if err != nil {
		a.logC.Error(err, "controller: get all abaout usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, aboutUsList)
}

func (a AboutUsController) AddAboutUs(c *gin.Context) {
	var req models.AboutUsRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := a.errH.ValidateRequest(req)
	if err != nil {
		a.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	about, err := a.uc.AddAbout(req)
	if err != nil {
		a.logC.Error(err, "controller: add about us usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, about)
}

func (a AboutUsController) UpdateAboutUs(c *gin.Context) {
	var req models.AboutUsRequest
	id := c.Query("id")
	idRes, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := a.errH.ValidateRequest(req)
	if err != nil {
		a.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	about, err := a.uc.UpdateData(idRes, req)
	if err != nil {
		a.logC.Error(err, "controller: update about us usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, about)
}

func (a AboutUsController) DeleteAboutUs(c *gin.Context) {
	id := c.Query("id")
	idRes := strings.Split(id, ",")

	err := a.uc.DeleteData(idRes)
	if err != nil {
		a.logC.Error(err, "controller: error when delete data submissions", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, nil)
}

func (a AboutUsController) GetAboutUsById(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		a.logC.Error(err, "controller: convert string to int in param", "", nil, nil, nil)
		c.Error(errors.New(contract.ErrBadRequest))
		c.Abort()
		return
	}

	result, err := a.uc.GetById(idInt)
	if err != nil {
		a.logC.Error(err, "controller: get by id about us usecase", "", nil, idInt, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
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

func TestGetAbouUsMethodTwo(t *testing.T) {
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.GET("/api/test/code_review/repo")

}
