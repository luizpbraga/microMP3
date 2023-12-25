package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/luizpbraga/microMP3/services/auth/src/database/connection"
	"github.com/luizpbraga/microMP3/services/auth/src/jwtutils"
	"log"
)

// Login usa o header Authorization para verificar a assinatura de um usuario
func Login(c *fiber.Ctx) error {
	// JWT encoded_auth token
	encodedAuth := c.Get("Authorization")
	if encodedAuth == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing Credentials",
		})
	}

	decodedAuth, err := jwtutils.DecodeAuth(encodedAuth)

	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing Credentials",
		})
	}

	var email, password string

	if err := connection.Db.QueryRow("SELECT email, password FROM user WHERE email = ?", decodedAuth.Email).Scan(&email, &password); err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database Failed",
		})
	}

	if decodedAuth.Email != email || decodedAuth.Password != password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	tokenString, err := jwtutils.GenerateTokenFromEmail(email, true)
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

func Validate(c *fiber.Ctx) error {
	encodedJwt := c.Get("Authorization")
	token, err := jwtutils.ValidateToken(encodedJwt)
	if err != nil {
		log.Print(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	email := claims["email"].(string)
	return c.Status(fiber.StatusOK).SendString(email)
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
