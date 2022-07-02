package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func hs256() {
	jwtSecret := "abc"
	token := GenerateHS256JWT(
		"claim1",
		string(jwtSecret),
		time.Hour*24*365,
	)
	fmt.Println(token)
	claims, err := DecodeHS256JWT(jwtSecret, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(claims)
}

// GenerateJWT generate jwt tokens
func GenerateHS256JWT(
	claim1 string,
	secret string,
	expire time.Duration,
) string {
	s := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["claim1"] = claim1
	claims["exp"] = time.Now().Add(expire).Unix()
	t, _ := token.SignedString(s)
	return t
}

func DecodeHS256JWT(jwtKey string, tokenString string) (map[string]interface{}, error) {
	secret := []byte(jwtKey)
	keyvaluemap := make(map[string]interface{})
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// fmt.Println("Signature:", token.Signature)
	if err != nil {
		fmt.Println(err)
		if err.Error() == jwt.ErrSignatureInvalid.Error() {
			for key, val := range claims {
				keyvaluemap[key] = val
			}
			return keyvaluemap, nil
		}
		return keyvaluemap, err
	}
	_ = token
	// fmt.Println(token)
	for key, val := range claims {
		keyvaluemap[key] = val
	}
	return keyvaluemap, nil
}
