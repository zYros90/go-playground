package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func rs256() {
	token := GenerateRS256JWT(
		"abc",
		"abcd",
	)
	fmt.Println("token: ", token)

	fmt.Println("validating")

	pubKey, err := ioutil.ReadFile("./jwtRS256.key.pub")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ValidateRS256JWT(token, pubKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}

// GenerateRS256JWT generate jwt tokens""
func GenerateRS256JWT(
	claim1 string,
	claim2 string,
) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["claim1"] = claim1
	claims["claim2"] = claim2

	prvKey, _ := ioutil.ReadFile("./jwtRS256.key")
	tok, err := Create(prvKey, time.Hour, claims)
	if err != nil {
		log.Fatalln(err)
	}
	return tok
}

func Create(privateKey []byte, ttl time.Duration, claims jwt.MapClaims) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()
	// claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix() // The time at which the token was issued.
	claims["nbf"] = now.Unix() // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateRS256JWT(token string, publicKey []byte) (map[string]interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}
	keyvaluemap := make(map[string]interface{})
	for key, val := range claims {
		keyvaluemap[key] = val
	}
	return keyvaluemap, nil
}
