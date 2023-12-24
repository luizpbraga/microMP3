package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/luizpbraga/microMP3/src/database"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func keyFunc(token *jwt.Token) (any, error) {
	return secretKey, nil
}

type authEncoded struct {
	email, password string
}

// Decods the Authorization HEADER
func decodeAuth(auth string) (*authEncoded, error) {
	// since auth stats with "Bearer " , we ignore this part
	decodedCredential, err := base64.StdEncoding.DecodeString(auth[7:])
	if err != nil {
		return nil, fmt.Errorf("Decodification error: %w", err)
	}

	credential := strings.SplitN(string(decodedCredential), ":", 2)
	if len(credential) != 2 {
		return nil, errors.New("Bad Request: Invalid Credential")
	}

	return &authEncoded{email: credential[0], password: credential[1]}, nil
}

// creates a JWT token with a custom claim and sets an expiration time.
func generateTokenFromEmail(email string) (string, error) {
	// novo token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	// Definir os claims do token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Assinar o token com a secret key
	return token.SignedString(secretKey)
}

// Login usa o header Authorization para verificar a assinatura de um usuario
func Login(c *fiber.Ctx) error {
	// JWT encoded_auth token
	encodedAuth := c.Get("Authorization")
	if encodedAuth == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing Credentials",
		})
	}

	decodedAuth, err := decodeAuth(encodedAuth)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing Credentials",
		})
	}

	var email, password string
	if err := database.Db.QueryRow("SELECT email, password FROM user WHERE email = ?", decodedAuth.email).Scan(&email, &password); err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database Failed",
		})
	}

	if decodedAuth.email != email || decodedAuth.password != password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	// novo token JWT
	token := jwt.New(jwt.SigningMethodHS256)
	// Definir os claims do token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Assinar o token com a secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "JWT error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "token is " + tokenString,
	})
}

// // middleware responsible for validating the JWT token sent in the Authorization header.
// func authMiddleware(c *fiber.Ctx) error {
// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized access",
// 		})
// 	}
//
// 	// see keyFunc
// 	parsedToken, err := jwt.Parse(token, keyFunc)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized access",
// 		})
// 	}
//
// 	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
// 		fmt.Println(claims["authorized"])
// 		return c.Next()
// 	}
//
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": "Unauthorized access",
// 	})
// }
