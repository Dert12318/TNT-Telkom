package controllers

import (
	"CartApp/config/log"
	"CartApp/models"
	"CartApp/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	uc   usecase.BlogUcInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewBlogController(r *gin.RouterGroup, uc usecase.BlogUcInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &BlogController{uc: uc,
		errH: errH,
		logC: logC,
	}

	r.POST("/add-blog", handler.Create)
	r.GET("/blogs", handler.Get)
	r.GET("/blog", handler.GetById)
	r.PUT("/update-blog", handler.Update)
	r.DELETE("/delete-blog", handler.Delete)
}

func (b BlogController) Create(c *gin.Context) {
	var req models.Blog
	if err := c.ShouldBindJSON(&req); err != nil {
		b.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := b.errH.ValidateRequest(req)
	if err != nil {
		b.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	result, err := b.uc.AddBlog(req)
	if err != nil {
		b.logC.Error(err, "controller: add blog usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
}

func (b BlogController) Get(c *gin.Context) {

	result, err := b.uc.GetAll()
	if err != nil {
		b.logC.Error(err, "controller: get blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
}

func (b BlogController) GetById(c *gin.Context) {
	id := c.Query("id")
	idRes, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	result, err := b.uc.GetById(idRes)
	if err != nil {
		b.logC.Error(err, "controller: get blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
}

func (b BlogController) Update(c *gin.Context) {
	var req models.Blog
	id := c.Query("id")
	idRes, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		b.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := b.errH.ValidateRequest(req)
	if err != nil {
		b.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	result, err := b.uc.UpdateData(idRes, req)
	if err != nil {
		b.logC.Error(err, "controller: update blog usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, result)
}

func (b BlogController) Delete(c *gin.Context) {
	id := c.Query("id")
	idRes := strings.Split(id, ",")

	err := b.uc.DeleteData(idRes)
	if err != nil {
		b.logC.Error(err, "controller: delete blog usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	ResponseSuccess(c, nil)
}
