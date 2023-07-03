package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SergeyCherepiuk/todo-app/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func NewOwnerMiddleware(idLabel string, check func(uint64, uint64) (bool, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(utils.CopyString(c.Params(idLabel)), 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(controllers.MessageResponse{
				Message: fmt.Sprintf("%s is not a valid id", c.Params(idLabel)),
			})
		}
		userId := c.Locals("userId").(uint64)

		isOwner, err := check(id, userId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(controllers.MessageResponse{
				Message: fmt.Sprintf("Internal server error: %s", err.Error()),
			})
		}

		if !isOwner {
			return c.Status(http.StatusNotFound).JSON(controllers.MessageResponse{
				Message: fmt.Sprintf("User with id %d has no item with id %d", userId, id),
			})
		}
		return c.Next()
	}
}
