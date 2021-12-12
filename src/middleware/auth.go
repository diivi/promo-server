package middleware

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(400).JSON(fiber.Map{"message": "Unauthorized"})
	}
	return c.Next()
}

func GetUserId(c *fiber.Ctx) (uint,error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return 0,err
	}
	payload := token.Claims.(*jwt.StandardClaims)
	id,_ := strconv.Atoi(payload.Subject)
	return uint(id),nil
}
//change jwt secret
