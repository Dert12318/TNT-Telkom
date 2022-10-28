package cart_repo

import (
	// "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"SoalNo6/models"
	"SoalNo6/repo"
	"fmt"

	"gorm.io/gorm"
)

type CartRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) repo.CartRepoInterface {
	return &CartRepoStruct{db}
}

func (a CartRepoStruct) AddNewProduct(Product models.Cart) error {
	tx := a.db.Begin()
	err3 := a.db.Debug().Create(&Product).Error
	if err3 != nil {
		tx.Rollback()
		return err3
	}
	tx.Commit()
	return nil
}
func (a CartRepoStruct) AddExistProduct(Product models.Cart) error {
	tx := a.db.Begin()
	err3 := a.db.Debug().Model(&models.Cart{}).Where("namaProduk = ?", Product.NamaProduk).Update("kuantitas", Product.Kuantitas).Error
	if err3 != nil {
		tx.Rollback()
		return err3
	}
	tx.Commit()
	return nil
}
func (a CartRepoStruct) CheckNamaProduct(Product models.Cart) (models.Cart, error) {
	var data models.Cart
	tx := a.db.Begin()
	err := a.db.Debug().Where("namaProduk = ?", Product.NamaProduk).Find(&data).Error
	fmt.Println("error repo_login line 44 : ", err)
	if err != nil {
		tx.Rollback()
		return models.Cart{}, err
	}
	tx.Commit()
	return data, nil
}
func (a CartRepoStruct) CheckKodeProduct(kodeProduk string) (models.Cart, error) {
	var data models.Cart
	tx := a.db.Begin()
	err := a.db.Debug().Where("kodeProduk = ?", kodeProduk).Find(&data).Error
	fmt.Println("error repo_login line 56 : ", err)
	if err != nil {
		tx.Rollback()
		return models.Cart{}, err
	}
	tx.Commit()
	return data, nil
}
func (a CartRepoStruct) DeleteProduct(Product models.Cart) error {
	var data models.Cart
	tx := a.db.Begin()
	err := a.db.Debug().Delete(&data, Product.KodeProduk).Error
	fmt.Println("error repo_login line 68 : ", err)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (a CartRepoStruct) GetAllProduct() (models.Carts, error) {
	var data models.Carts
	tx := a.db.Begin()
	err := a.db.Debug().Find(&data).Error
	fmt.Println("error repo_login line 80 : ", err)
	if err != nil {
		tx.Rollback()
		return models.Carts{}, err
	}
	tx.Commit()
	return data, nil
}
