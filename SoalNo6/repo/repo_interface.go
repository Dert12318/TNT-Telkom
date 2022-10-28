package repo

import (
	"SoalNo6/models"
)

type CartRepoInterface interface {
	AddNewProduct(Product models.Cart) error
	AddExistProduct(Product models.Cart) error
	CheckNamaProduct(Product models.Cart) (models.Cart, error)
	CheckKodeProduct(kodeProduk string) (models.Cart, error)
	DeleteProduct(Product models.Cart) error
	GetAllProduct() (models.Carts, error)
}
