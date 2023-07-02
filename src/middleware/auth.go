package middleware

import (
	"net/http"
	"os"

	"github.com/SergeyCherepiuk/todo-app/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func validate(signedToken string) bool {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(signedToken, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func getUserId(signedToken string) (uint64, error) {
	claims := repositories.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(signedToken, &claims, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		signedToken := c.Cookies("token", "")
		if !validate(signedToken) {
			return c.SendStatus(http.StatusUnauthorized)
		}
		userId, err := getUserId(signedToken)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		c.Locals("userId", userId)
		return c.Next()
	}
}
