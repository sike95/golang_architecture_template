package controllers

import (
	"golang_architectire_template/domain/users"
	"golang_architectire_template/services"
	"golang_architectire_template/utils/rest_errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getUserId(userIdParam string) (int64, *rest_errors.RestErr) {
	userId, _ := strconv.ParseInt(userIdParam, 10, 64)

	return userId, nil
}

func Create(c *fiber.Ctx) error {
	var user users.User
	result, _ := services.UsersService.CreateUser(user)
	// if saveErr != nil {
	// 	return c.Status(saveErr.Status).JSON(saveErr)
	// }
	return c.Status(fiber.StatusCreated).JSON(result.Marshall(c.Get("X-Public") == "true"))
}
func Get(c *fiber.Ctx) error {
	userId, _ := getUserId(c.Params("user_id"))
	// if idErr != nil {
	// 	return c.Status(idErr.Status).JSON(idErr)
	// }
	user, _ := services.UsersService.GetUser(userId)
	// if getErr != nil {
	// 	return c.Status(getErr.Status).JSON(getErr)
	// }

	return c.Status(fiber.StatusOK).JSON(user.Marshall(c.Get("X-Public") == "true"))
}

func Update(c *fiber.Ctx) error {
	userId, _ := getUserId(c.Params("user_id"))
	// if idErr != nil {
	// 	return c.Status(idErr.Status).JSON(idErr)
	// }

	var user users.User
	// if err := c.BodyParser(&user); err != nil {
	// 	restErr := rest_errors.NewBadRequestError("invalid json body")
	// 	return c.Status(restErr.Status).JSON(restErr)
	// }

	user.Id = userId
	isPartial := c.Method() == fiber.MethodPatch

	result, _ := services.UsersService.UpdateUser(isPartial, user)
	// if err != nil {
	// 	return c.Status(err.Status).JSON(err)
	// }
	return c.Status(fiber.StatusOK).JSON(result.Marshall(c.Get("X-Public") == "true"))
}

func Delete(c *fiber.Ctx) error {
	userId, _ := getUserId(c.Params("user_id"))
	// if idErr != nil {
	// 	return c.Status(idErr.Status).JSON(idErr)
	// }

	services.UsersService.DeleteUser(userId)
	// if err != nil {
	// 	return c.Status(err.Status).JSON(err)
	// }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "deleted"})
}

func Search(c *fiber.Ctx) error {
	status := c.Query("status")

	users, _ := services.UsersService.SearchUser(status)
	// if err != nil {
	// 	return c.Status(err.Status).JSON(err)
	// }
	return c.Status(fiber.StatusOK).JSON(users.Marshall(c.Get("X-Public") == "true"))
}

func Login(c *fiber.Ctx) error {
	var request users.LoginRequest
	user, _ := services.UsersService.LoginUser(request)
	// if err != nil {
	// 	return c.Status(err.Status).JSON(err)
	// }
	return c.Status(fiber.StatusOK).JSON(user.Marshall(c.Get("X-Public") == "true"))
}
