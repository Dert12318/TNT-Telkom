package repo

import "CartApp/models"

type LoginRepoInterface interface {
	AddProduct(Product models.Cart) error
}
