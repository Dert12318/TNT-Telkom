package repo

import (
	"github.com/Dert12318/TNT-Telkom.git/models"
)

type LoginRepoInterface interface {
	AddProduct(Product models.Cart) error
}
