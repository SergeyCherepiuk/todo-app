package middleware

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func validate(c *fiber.Ctx) bool {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	signedToken := c.Cookies("token")
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if validate(c) {
			return c.Next()
		}
		return c.SendStatus(http.StatusUnauthorized)
	}
}
