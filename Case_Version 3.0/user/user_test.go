package user_test

import (
	"main/model"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestHandlerCreate(t *testing.T) {
	app := fiber.New()

	app.Post("/UserCreate", func(c *fiber.Ctx) error {
		body := &model.User{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/UserCreate")
	rb.Json(&model.User{User_Name: "fs", Name: "furkan", Surname: "samaraz", Is_Active: "False"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.User{User_Name: "fs", Name: "furkan", Surname: "samaraz", Is_Active: "False"})

}
func TestUsersEntryGet(t *testing.T) {
	app := fiber.New()

	app.Post("/UsersEntryAllGet", func(c *fiber.Ctx) error {
		body := &model.UserEntry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/UsersEntryAllGet")
	rb.Json(&model.UserEntry{Id: 2})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.UserEntry{Id: 2})

}
