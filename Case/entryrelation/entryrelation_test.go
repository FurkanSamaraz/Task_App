package entryrelation_test

import (
	"main/model"
	"testing"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestEntryeReltMainGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryeReltMainGet", func(c *fiber.Ctx) error {
		body := &model.EntryRelation{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryeReltMainGet")
	rb.Json(&model.EntryRelation{Main_Entry: "Status"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.EntryRelation{Main_Entry: "Status"})

}
func TestEntryReltUpdate(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryReltUpdate", func(c *fiber.Ctx) error {
		body := &model.EntryRelation{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryReltUpdate")
	rb.Json(&model.EntryRelation{Id: 1, Main_Entry: "124415251251", Sub_Entry: "null", Parent_Entry: "null"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.EntryRelation{Id: 1, Main_Entry: "124415251251", Sub_Entry: "null", Parent_Entry: "null"})

}
func TestEntryReltCreate(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryReltCreate", func(c *fiber.Ctx) error {
		body := &model.EntryRelation{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryReltCreate")
	rb.Json(&model.EntryRelation{Main_Entry: "124415251251", Sub_Entry: "null", Parent_Entry: "null"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.EntryRelation{Main_Entry: "124415251251", Sub_Entry: "null", Parent_Entry: "null"})

}
