package usecase

import "SoalNo6/models"

type CartInterface interface {
	AddCart(newCart models.Cart) error
	GetAllCart() (models.Carts, error)
	DeleteCart(kodeProduk string) error
}
type ErrorHandlerUsecase interface {
	ResponseError(error interface{}) (int, interface{})
	ValidateRequest(error interface{}) (string, error)
}
