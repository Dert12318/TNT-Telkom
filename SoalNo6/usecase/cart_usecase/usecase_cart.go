package user_usecase

import (
	"SoalNo6/models"
	"SoalNo6/repo"
	"SoalNo6/usecase"
)

type CartUsecaseStruct struct {
	repo repo.CartRepoInterface
}

func NewUserUsecase(repo repo.CartRepoInterface) usecase.CartInterface {
	return &CartUsecaseStruct{
		repo: repo,
	}
}

func (a CartUsecaseStruct) AddCart(newCart models.Cart) error {
	res, err := a.repo.CheckNamaProduct(newCart)
	if err != nil {
		// Nama produk tidak ada
		err2 := a.repo.AddNewProduct(newCart)
		if err2 != nil {
			return err2
		} else {
			return nil
		}
	} else {
		res.KodeProduk = res.KodeProduk + newCart.KodeProduk
		err3 := a.repo.AddNewProduct(res)
		if err3 != nil {
			return err3
		} else {
			return nil
		}
	}
}
func (a CartUsecaseStruct) GetAllCart() (models.Carts, error) {
	res, err := a.repo.GetAllProduct()
	if err != nil {
		return models.Carts{}, err
	}
	return res, nil
}
func (a CartUsecaseStruct) DeleteCart(kodeProduk string) error {
	res, err := a.repo.CheckKodeProduct(kodeProduk)
	if err != nil {
		err2 := a.repo.DeleteProduct(res)
		if err2 != nil {
			return err2
		} else {
			return nil
		}
	}
	return err
}
