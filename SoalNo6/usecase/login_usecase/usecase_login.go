package login_usecase

import (
	"encoding/json"
	"fmt"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/auth"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/env"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/usecase"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/utils"

	"github.com/pkg/errors"
)

type LoginUsecaseStruct struct {
	repo repo.LoginRepoInterface
}

func NewLoginUsecase(repo repo.LoginRepoInterface) usecase.LoginUcInterface {
	return &LoginUsecaseStruct{
		repo: repo,
	}
}

func (a LoginUsecaseStruct) LoginUser(encrpytData models.EncryptData) (models.TokenStruct, error) {
	key := env.Config.EncKey

	decrypt, err := utils.KeyDecrypt(key, encrpytData.Encrypt)

	fmt.Println(key, "WKWKWKWK")

	userRequest := models.UserRequest{}
	err = json.Unmarshal([]byte(decrypt), &userRequest)
	if err != nil {
		errs := errors.New("Error when decrypt")
		fmt.Println("[LoginUsecaseStruct.LoginUser] ", errs)
		return models.TokenStruct{}, errs
	}
	//to do check validate userRequest is null ?

	//check to userRequest to database
	userData, err := a.repo.LoginUser(userRequest.Email)

	//compare password
	valid := utils.CheckPasswordHash(userData.Password, []byte(userRequest.Password))

	if valid != true {
		errs := errors.New("Invalid Password")
		fmt.Println("[LoginUsecaseStruct.LoginUser] ", errs)
		return models.TokenStruct{}, errs
	}
	//Insert to table auth for log login
	dataFix := models.Auth{
		Username: userData.Username,
		Email:    userData.Email,
		Role:     userData.Role,
	}

	userAuth, err := a.repo.CreateAuth(dataFix)
	if err != nil {
		fmt.Println("[LoginUsecaseStruct.LoginUser] ", err)
		return models.TokenStruct{}, err
	}

	token, err := auth.CreateToken(userAuth)
	if err != nil {
		fmt.Println("[LoginUsecaseStruct.LoginUser] ", err)
		return models.TokenStruct{}, err
	}

	var JWT models.TokenStruct
	JWT.Token = token

	return JWT, nil

}

func (a *LoginUsecaseStruct) DeleteAuthData(givenUuid string) (int, error) {
	return a.repo.DeleteAuthData(givenUuid)
}
