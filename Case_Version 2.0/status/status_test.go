package status_test

import (
	"main/model"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestStatusCrea(t *testing.T) {
	app := fiber.New()

	app.Post("/StatusCreate", func(c *fiber.Ctx) error {
		body := &model.Status{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/StatusCreate")
	rb.Json(&model.Status{Name: "Status"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Status{Name: "Status"})

}
func TestStatusUpd(t *testing.T) {
	app := fiber.New()

	app.Post("/StatusUpdate", func(c *fiber.Ctx) error {
		body := &model.Status{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/StatusUpdate")
	rb.Json(&model.Status{Id: 1, Name: "Status"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Status{Id: 1, Name: "Status"})

}
