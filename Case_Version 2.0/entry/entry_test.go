package entry_test

import (
	"main/model"
	"testing"
	"time"

	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
)

func TestEntryUpdate(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryUpdate", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryUpdate")
	rb.Json(&model.Entry{Id: 1, Entry_Title: "Hello", Status: "In Progress", Assig: "2"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Id: 1, Entry_Title: "Hello", Status: "In Progress", Assig: "2"})

}
func TestEntryCreate(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryCreate", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryCreate")
	rb.Json(&model.Entry{Entry_Title: "Hello", Status: "In Progress", Assig: "2", Tag: "1001"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Entry_Title: "Hello", Status: "In Progress", Assig: "2", Tag: "1001"})

}
func TestEntryStatusGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryStatusGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryStatusGet")
	rb.Json(&model.Entry{Status: "In Progress"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Status: "In Progress"})

}
func TestEntryTimeCreGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryTimeCreGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryTimeCreGet")
	rb.Json(&model.Entry{Create_Date: time.Now()})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Create_Date: time.Now()})

}
func TestEntryTimeUpdGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryTimeUpdGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryTimeUpdGet")
	rb.Json(&model.Entry{Create_Date: time.Now()})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Create_Date: time.Now()})

}
func TestEntryAllGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryAllGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryAllGet")
	rb.Json(&model.Entry{Entry_Code: "212314124124141"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Entry_Code: "212314124124141"})

}
func TestEntryTagGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryTagGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryTagGet")
	rb.Json(&model.Entry{Tag: "1002"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Tag: "1002"})

}
func TestEntryTopAllGet(t *testing.T) {
	app := fiber.New()

	app.Post("/EntryTopAllGet", func(c *fiber.Ctx) error {
		body := &model.Entry{}
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(body)
	})

	rb := httptest.Post("/EntryTopAllGet")
	rb.Json(&model.Entry{Entry_Code: "1241513534214"})

	test := rb.Test(t, app)
	test.Status(200)
	test.Json(&model.Entry{Entry_Code: "1241513534214"})

}
