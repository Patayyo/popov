package handlers

import (
	"github.com/gorepos/usercartv2/internal/application"
)

type UsersHandler struct {
	App *application.Application
}

/*func (uh *UsersHandler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := uh.App.S.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	return c.JSON(users)
}*/
