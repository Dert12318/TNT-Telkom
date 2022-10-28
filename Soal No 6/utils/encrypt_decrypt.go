package utils

import (
	"CartApp/models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func KeyEncrypt(keyStr string, cryptoText string) string {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return Encrypt(keyBytes[:], cryptoText)
}

// Encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// this is function for decrypt
func KeyDecrypt(keyStr string, cryptoText string) (string, error) {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return Decrypt(keyBytes[:], cryptoText), nil
}

// Decrypt from base64 to Decrypted string
func Decrypt(key []byte, cryptoText string) string {
	//var c *gin.Context
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		fmt.Printf("Error when decrypt from data encrypt %v", err)
		//ErrorMessage(c, http.StatusInternalServerError, "Invalid Encrypt")
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		fmt.Println("CIPHERTEXT TOO SHORT")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

func HashPassword(user *models.User) (*models.User, error) {
	if len(user.Password) == 0 {
		return nil, errors.New("password should not be empty!")
	}
	bytePassword := []byte(user.Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		fmt.Printf("[UserController.SetPassword] Error when generate password with error: %v\n", err)
		return nil, err
	}
	user.Password = string(passwordHash)
	return user, nil
}

func CheckPasswordHash(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
