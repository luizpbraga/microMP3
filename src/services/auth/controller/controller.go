package controller

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = []byte("BANANA")

func keyFunc(token *jwt.Token) (any, error) {
	return secretKey, nil
}

// creates a JWT token with a custom claim (authorized: true) and sets an expiration time.
func generateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	return token.SignedString(secretKey)
}

func Login(c *fiber.Ctx) error {
	// JWT auth token
	// token, err := generateToken()
	auth := c.Get("Authorization")

	if auth == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Token generation failed",
		})
	}

	return c.JSON(fiber.Map{"token": auth})
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
