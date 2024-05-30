package jwtToken

import (
	"Arqui_Soft_I/backend/dto"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateUserToken(userDto dto.UserDto) (string, error) {
	claims := jwt.MapClaims{
		"id":       userDto.ID,
		"username": userDto.Username,
		"rol":      userDto.Rol,
		"nombre":   userDto.Nombre,
		"apellido": userDto.Apellido,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := "secreto"
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := "secreto"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
