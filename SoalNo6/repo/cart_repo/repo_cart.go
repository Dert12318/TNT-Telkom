package cart_repo

import (
	"github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"fmt"

	"gorm.io/gorm"
)

type LoginRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) repo.LoginRepoInterface {
	return &LoginRepoStruct{db}
}

func (a LoginRepoStruct) AddProduct(Product models.Cart) error {
	var data models.Cart
	tx := a.db.Begin()
	err := a.db.Debug().Where("addProduct = ?", Product.NamaProduk).Find(&data).Error
	fmt.Println("error repo_login line 23 : ", err)
	if err != nil {
		err2 := a.db.Debug().Create(&Product).Error
		if err2 != nil{
			tx.Rollback()
			return err2
		}
	} else {
		data.Kuantitas := data.Kuantitas + Product.Kuantitas
		err2 := a.db.Debug().Create(&data).Error
		if err2 != nil{
			tx.Rollback()
			return err3
		}
	}
	tx.Commit()
	return nil
}
