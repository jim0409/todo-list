package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ep(err error) {
	if err != nil {
		panic(err)
	}
}

func GetCode(key string) (string, error) {
	inputNoSpacesUpper := strings.ToUpper(key)
	encodeKey, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		fmt.Println(err)
	}
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(encodeKey, toBytes(epochSeconds/30))
	return fmt.Sprintf("%06d", pwd), nil
}

func oneTimePassword(key []byte, value []byte) uint32 {
	hmacSha1 := hmac.New(sha1.New, key) // sign the value using HMAC-SHA1
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)
	pwd := number % 1000000

	return pwd
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

var jwtKey = []byte("test")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func EncryptJwt(name string, role string) string {

	claims := &Claims{
		Username:       name,
		Role:           role,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func DecryptJwt(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	_ = tkn

	if err != nil {
		return nil, err
	}
	return claims, nil
}

func ValidatedJWT(tkn string) (string, string, error) {
	c, err := DecryptJwt(tkn)
	if err != nil {
		return "", "", err
	}

	return c.Username, c.Role, nil
}
