package entrycomment_test

import (
	"main/model"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestEntryeReltMainGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryComCreate", func(c *fiber.Ctx) error {
		body := &model.EntryComment{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryComCreate")
	rb.Json(&model.EntryComment{Entry_id: "242423124", User_id: "2", Text: "Hello", Is_Active: "True"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.EntryComment{Entry_id: "242423124", User_id: "2", Text: "Hello", Is_Active: "True"})

}
func TestEntryComRemove(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryComRemove", func(c *fiber.Ctx) error {
		body := &model.EntryComment{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryComRemove")
	rb.Json(&model.EntryComment{Entry_id: "141241414", User_id: "2", Is_Active: "False"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.EntryComment{Entry_id: "141241414", User_id: "2", Is_Active: "False"})

}
