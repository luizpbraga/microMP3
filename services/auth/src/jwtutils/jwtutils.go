package jwtutils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func keyFunc(token *jwt.Token) (any, error) {
	return secretKey, nil
}

func validateFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("invalid signing method")
	}
	return secretKey, nil
}

type authEncoded struct {
	Email, Password string
}

// Decods the Authorization HEADER
func DecodeAuth(auth string) (*authEncoded, error) {
	// since auth stats with "Bearer " , we ignore this part
	decodedCredential, err := base64.StdEncoding.DecodeString(auth[7:])
	if err != nil {
		return nil, fmt.Errorf("Decodification error: %w", err)
	}

	credential := strings.SplitN(string(decodedCredential), ":", 2)
	if len(credential) != 2 {
		return nil, errors.New("Bad Request: Invalid Credential")
	}

	return &authEncoded{Email: credential[0], Password: credential[1]}, nil
}

func ValidateToken(encodedJwt string) (*jwt.Token, error) {
	if encodedJwt == "" || !strings.HasPrefix("Bearer ", encodedJwt) {
		return nil, fmt.Errorf("Bad request")
	}

	tokenString := strings.Split(encodedJwt, " ")[1]
	return jwt.Parse(tokenString, validateFunc)
}

// creates a JWT token with a custom claim and sets an expiration time.
func GenerateTokenFromEmail(email string, admin bool) (string, error) {
	// novo token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	// Definir os claims do token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Assinar o token com a secret key
	return token.SignedString(secretKey)
}
