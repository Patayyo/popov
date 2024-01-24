package handlers

//"log"
//"strconv"

//"github.com/gofiber/fiber/v2"
//"github.com/gorepos/usercartv2/internal/store"

/*func (uh *UsersHandler) AddItemToUserCartHandler(c *fiber.Ctx) error {
	userIDStr := c.Params("UserID")
	itemID := c.Params("ItemID")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Printf("Error converting UserID to int: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UserID"})
	}

	log.Printf("Adding item to user's cart (UserID: %d, ItemID: %s)", userID, itemID)

	user, err := uh.App.S.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user"})
	}

	if user == nil {
		log.Printf("User not found with ID: %d", userID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	item, err := uh.App.S.GetItemByID(itemID)
	if err != nil {
		log.Printf("Error fetching item by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get item"})
	}

	if item == nil {
		log.Printf("Item not found with ID: %s", itemID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	if err := uh.App.S.AddItemToUserCart(userID, *item); err != nil {
		log.Printf("Error adding item to user's cart: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to user's cart"})
	}

	log.Printf("Item added to user's cart successfully")
	return c.SendString("Item added to user's cart successfully")
}*/

/*func (uh *UsersHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	userIDStr := c.Params("UserID")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Printf("Error converting UserID to int: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UserID"})
	}

	log.Printf("Fetching user by ID: %d", userID)

	user, err := uh.App.S.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user"})
	}

	if user == nil {
		log.Printf("User not found with ID: %d", userID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	log.Printf("Successfully fetched user by ID: %d", userID)
	return c.JSON(user)
}*/

/*func (uh *UsersHandler) AddUserHandler(c *fiber.Ctx) error {
	var newUser store.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	log.Printf("Adding new user: %v", newUser)

	if err := uh.App.S.AddUser(newUser); err != nil {
		log.Printf("Error adding new user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add user")
	}

	log.Printf("Successfully added new user")
	return c.SendString("User added successfully")
}*/

/*func (uh *UsersHandler) DeleteUserHandler(c *fiber.Ctx) error {
	userID := c.Params("UserID")

	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UserID is required"})
	}

	// Преобразование строки в число (int)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UserID format"})
	}

	// Вызов метода удаления пользователя из вашей службы
	err = uh.App.S.DeleteUser(userIDInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendString("User deleted successfully")
}*/

/*func (uh *UsersHandler) AddUserRoleHandler(c *fiber.Ctx) error {
	userID := c.Params("UserID")
	role := c.Params("Role")

	// Преобразуем userID в int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Добавляем роль пользователю в хранилище данных
	if err := uh.App.S.AddUserRole(userIDInt, role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add role to user"})
	}

	return c.SendString("Role added to user successfully")
}*/
