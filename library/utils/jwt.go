package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret []byte
	jwtExpire time.Duration = 3 * time.Hour
)

type Claims struct {
	Avatar       string   `json:"avatar"`
	Roles        []string `json:"role"`
	Introduction string   `json:"introduction"`
	Name         string   `json:"name"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
//func GenerateToken(username, password string) (string, error) {
func GenerateToken(name, avatar, introduction string, roles int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(jwtExpire)
	//EncodeMD5(username),
	//EncodeMD5(password),
	claims := Claims{
		Avatar: avatar,
		//Roles:        []string{roles},
		Roles:        []string{strconv.Itoa(roles)},
		Introduction: introduction,
		Name:         name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
