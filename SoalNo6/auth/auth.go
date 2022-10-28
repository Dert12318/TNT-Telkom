package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/repo"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CreateToken(authD models.Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUUID
	claims["username"] = authD.Username
	claims["email"] = authD.Email
	claims["role"] = authD.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// Token Validation
func TokenValidCustom(r *http.Request, loginRepo repo.LoginRepoInterface) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	emptyStruct := models.Auth{}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string) //convert the interface to string
		if !ok {
			return err
		}
		email, done := claims["email"].(string)
		if !done {
			return err
		}
		checkUsernameAndAuth, err := loginRepo.GetAuthByEmailAndAuthID(email, authUuid)
		if err != nil {
			return err
		}
		if checkUsernameAndAuth.DeletedAt != emptyStruct.DeletedAt {
			return errors.New("1")
		}
	}

	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*models.Auth, error) {
	var res models.Auth
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims, "LLLLLLL")
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string)
		fmt.Println(authUuid, "OOOOOOOOOOOO")
		if !ok {
			return nil, err
		}
		username, done := claims["username"].(string)
		if !done {
			return nil, err
		}
		fmt.Println(username, ">>>>>>>>>>>")
		email, done := claims["email"].(string)
		//userIdRes, err := strconv.Atoi(userId)
		//if err != nil {
		//	return nil, err
		//}
		if !done {
			return nil, err
		}

		res.AuthUUID = authUuid
		res.Username = username
		res.Email = email
		return &res, nil

	}
	return nil, err
}
