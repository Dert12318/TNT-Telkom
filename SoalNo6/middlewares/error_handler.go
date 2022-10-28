package middlewares

import (
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	ErrorHandlerUsecase usecase.ErrorHandlerUsecase
	log                 *log.LogCustom
}

func NewErrorHandler(r *gin.RouterGroup, ehus usecase.ErrorHandlerUsecase, log *log.LogCustom) {
	handler := &ErrorHandler{
		ErrorHandlerUsecase: ehus,
		log:                 log,
	}

	r.Use(handler.errorHandler)
}

func (eh *ErrorHandler) errorHandler(c *gin.Context) {
	c.Next()

	errorToPrint := c.Errors.Last()
	if errorToPrint != nil {
		_, v := eh.ErrorHandlerUsecase.ResponseError(errorToPrint)
		s := v.(models.ResponseCustomErr)
		c.JSON(eh.ErrorHandlerUsecase.ResponseError(errorToPrint))
		eh.log.Error(errorToPrint, "middlewares/errorHandler", "", nil, nil, s)
		c.Abort()
		return
	}
}
