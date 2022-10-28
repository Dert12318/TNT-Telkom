package error

import (
	"CartApp/models"
	"CartApp/models/contract"
	"CartApp/usecase"
	"CartApp/utils"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type errorHandlerUsecase struct {
}

func NewErrorHandlerUsecase() usecase.ErrorHandlerUsecase {
	return &errorHandlerUsecase{}
}

func (e errorHandlerUsecase) ResponseError(A interface{}) (int, interface{}) {
	var T interface{}
	var fieldNameErr string
	var serviceCode string

	if A.(*gin.Error).Meta != nil {
		fieldNameErr = A.(*gin.Error).Meta.(models.ErrMeta).FieldErr
		serviceCode = A.(*gin.Error).Meta.(models.ErrMeta).ServiceCode
	}

	// Check A is a correct error type and assign to T
	if A.(*gin.Error).Err != nil {
		T = A.(*gin.Error).Err
	}

	fmt.Println(serviceCode)

	switch T.(type) {
	case error:
		if _, ok := T.(*pq.Error); ok {
			switch T.(*pq.Error).Code.Name() {
			case "unique_violation":
				return errcntrct.ErrorMessage(http.StatusBadRequest, "", errors.New(contract.ErrGeneralError))
			}
		}

		switch T.(error).Error() {
		case contract.ErrInvalidFieldFormat:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldFormat, serviceCode, fieldNameErr)
		case contract.ErrInvalidFieldMandatory:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldMandatory, serviceCode, fieldNameErr)
		case contract.ErrConflict:
			return responseErrorAdapter(T.(error), http.StatusConflict, "", serviceCode, "")
		case contract.ErrBadRequest:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "")
		case contract.ErrInternalError:
			return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "")
		case contract.ErrCannotSaveToDB:
			return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "")
		case contract.ErrCannotSaveToFile:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "")
		case contract.ErrUnauthorized:
			return responseErrorAdapter(T.(error), http.StatusUnauthorized, "", serviceCode, "")
		case contract.ErrTimeout:
			return responseErrorAdapter(T.(error), http.StatusRequestTimeout, "", serviceCode, "")
		case contract.ErrDataNotFound:
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", serviceCode, "")
		case contract.ErrDateSchedule:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "")
		default:
			return responseErrorAdapter(errors.New(contract.ErrGeneralError), http.StatusInternalServerError, "", serviceCode, "")
		}
	}

	return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "")
}

func (e errorHandlerUsecase) ValidateRequest(T interface{}) (string, error) {
	v := validator.New()
	var errArr error
	var field string
	switch T.(type) {
	case models.User:
		err := v.Struct(T)
		if err == nil {
			return "", nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil
	case models.AboutUsRequest:
		err := v.Struct(T)
		if err == nil {
			return "", nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				if e.Type() == reflect.TypeOf(e.Type()) {
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
				}
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
					if e.Field() == "value" {
						field = e.Field()
						errArr = errors.New(contract.ErrInvalidAmount)
					}
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil
	case models.ExpeditionSchedule:
		err := v.Struct(T)
		if err == nil {
			return "", nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
					if e.Field() == "value" {
						field = e.Field()
						errArr = errors.New(contract.ErrInvalidAmount)
					}
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil
	case models.Blog:
		err := v.Struct(T)
		if err == nil {
			return "", nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
					if e.Field() == "value" {
						field = e.Field()
						errArr = errors.New(contract.ErrInvalidAmount)
					}
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil

	default:
		return "", errors.New(contract.ErrGeneralError)
	}
}

func responseErrorAdapter(errHttpStatus interface{}, httpStatusCode int, ctr string, serviceCode string, fieldErr string) (int, models.ResponseCustomErr) {
	_, errData := errcntrct.ErrorMessage(httpStatusCode, "", errHttpStatus)
	var resp models.ResponseCustomErr
	errCase := strconv.Itoa(httpStatusCode)
	caseCode := strings.Split(errData.Code, "XX")
	resp.ResponseCode = errCase + models.ServiceCode + caseCode[1]
	if strings.Contains(contract.FieldErr, " ") {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, contract.FieldErr)
	} else if ctr == "400XX01" || ctr == "400XX02" {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, utils.LowerCamelCase(fieldErr))
	} else {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg)
	}
	return httpStatusCode, resp
}
