package handler

import (
	"main/entry"
	"main/entrycomment"
	"main/entryrelation"
	"main/status"
	"main/user"

	"github.com/gofiber/fiber/v2"
)

func Handlers() {
	app := fiber.New()

	app.Post("/EntryTopAllGet", entry.EntryTopAllGet)
	app.Post("/UsersEntryAllGet", user.UsersEntry)

	app.Post("/EntryComCreate", entrycomment.EntryComCreate)
	app.Get("/EntryComGet", entrycomment.EntryComGet)
	app.Post("/EntryComRemove", entrycomment.EntryComRemove)
	app.Get("/EntryTrueGet", entrycomment.EntryTrueGet)

	app.Post("/EntryAllGet", entry.EntryAllGet)
	app.Post("/EntryTagGet", entry.EntryTagGet)

	app.Get("/EntryGet", entry.EntryGet)
	app.Get("/EntryeReltGet", entryrelation.EntryeReltGet)
	app.Get("/StatusGet", status.StatusGet)
	app.Get("/UsersGet", user.UsersGet)

	app.Post("/EntryeReltMainGet", entryrelation.EntryeReltMainGet)
	app.Post("/EntryStatusGet", entry.EntryStatusGet)
	app.Post("/UsersActiveGet", user.UsersActiveGet)
	app.Post("/EntryTimeCreGet", entry.EntryTimeCreGet)
	app.Post("/EntryTimeUpdGet", entry.EntryTimeUpdGet)

	app.Post("/EntryUpdate", entry.EntryUpdate)
	app.Post("/EntryReltUpdate", entryrelation.EntryReltUpdate)
	app.Post("/StatusUpdate", status.StatusUpdate)

	app.Post("/EntryCreate", entry.EntryCreate)
	app.Post("/EntryReltCreate", entryrelation.EntryReltCreate)
	app.Post("/StatusCreate", status.StatusCreate)
	app.Post("/UserCreate", user.UserCreate)

	app.Listen(":8080")
}
