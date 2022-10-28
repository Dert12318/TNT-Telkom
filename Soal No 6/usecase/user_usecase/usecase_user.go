package user_usecase

import (
	"CartApp/models"
	"CartApp/repo"
	"CartApp/usecase"
	"CartApp/utils"
)

type UserUsecaseStruct struct {
	repo repo.UserRepoInterface
}

func NewUserUsecase(repo repo.UserRepoInterface) usecase.UserUcInterface {
	return &UserUsecaseStruct{
		repo: repo,
	}
}

func (a UserUsecaseStruct) AddUser(v models.User) (models.User, error) {

	userHash, err := utils.HashPassword(&v)

	user, err := a.repo.AddUser(*userHash)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (a UserUsecaseStruct) GetAll() ([]models.User, error) {
	AllUser, err := a.repo.GetAll()

	if err != nil {
		return []models.User{}, err
	}

	return AllUser, nil

}

func (a UserUsecaseStruct) GetById(id int) (models.User, error) {
	user, err := a.repo.GetById(id)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a UserUsecaseStruct) UpdateData(id int, v models.User) (models.User, error) {
	user, err := a.repo.UpdateData(id, v)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a UserUsecaseStruct) DeleteData(id []string) error {
	err := a.repo.DeleteData(id)

	if err != nil {
		return err
	}

	return nil
}
