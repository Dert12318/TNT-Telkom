package usecase

import "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"

type AboutUsUcInterface interface {
	AddAbout(v models.AboutUsRequest) (models.AboutUsRequest, error)
	GetAboutUs() (models.AboutUsRequest, error)
	GetById(id int) (models.AboutUsRequest, error)
	UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error)
	DeleteData(id []string) error
}

type ExpeditionUcInterface interface {
	AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	GetById(id int) (models.ExpeditionSchedule, error)
	GetAll(limit, offset, uri string) (models.ResponseDataPagination, error)
	Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	DeleteData(id []string) error
	GetByRoute(route string) ([]models.ExpeditionSchedule, error)
}

type UserUcInterface interface {
	AddUser(v models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	UpdateData(id int, v models.User) (models.User, error)
	DeleteData(id []string) error
}

type LoginUcInterface interface {
	LoginUser(encrpytData models.EncryptData) (models.TokenStruct, error)
	DeleteAuthData(givenUuid string) (int, error)
}

type ErrorHandlerUsecase interface {
	ResponseError(error interface{}) (int, interface{})
	ValidateRequest(error interface{}) (string, error)
}

type BlogUcInterface interface {
	AddBlog(v models.Blog) (models.Blog, error)
	GetAll() ([]models.Blog, error)
	GetById(id int) (models.Blog, error)
	UpdateData(id int, v models.Blog) (models.Blog, error)
	DeleteData(id []string) error
}
