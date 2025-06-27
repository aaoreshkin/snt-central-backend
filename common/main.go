package common

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

var supportedType = map[string]struct{}{
	"svg":  {},
	"pdf":  {},
	"doc":  {},
	"docx": {},
	"jpg":  {},
	"jpeg": {},
	"png":  {},
}

// Hash the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil, err
}

func GetExtension(header *multipart.FileHeader) (string, error) {
	extension := filepath.Ext(header.Filename)
	if len(extension) == 0 {
		return "", errors.New("file has no extension")
	}
	extension = extension[1:]

	if _, ok := supportedType[extension]; !ok {
		return "", fmt.Errorf("unsupported file format: %s", extension)
	}

	return extension, nil
}

func GenerateHex() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := hex.EncodeToString(randomBytes)

	return randomString, nil
}
