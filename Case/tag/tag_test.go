package tag_test

import (
	"main/model"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestTagCre(t *testing.T) {
	app := fiber.New()

	app.Post("/TagCre", func(c *fiber.Ctx) error {
		body := &model.TagProperties{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/TagCre")
	rb.Json(&model.TagProperties{Entry_id: "135341", Tag_id: "234"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.TagProperties{Entry_id: "135341", Tag_id: "234"})

}
