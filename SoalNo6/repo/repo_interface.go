package repo

import "github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"

type LoginRepoInterface interface {
	AddProduct(Product models.Cart) error
}
