package user

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/isvalid"
	"main/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var db = database.Openconnention()

func UserCreate(c *fiber.Ctx) error {
	var users model.User

	users.User_Name = c.FormValue("User_Name")
	users.Name = c.FormValue("Name")
	users.Surname = c.FormValue("Surname")
	users.Is_Active = c.FormValue("Is_Active")

	db.Exec("INSERT INTO users(user_Name,name,surname,is_Active) VALUES($1,$2,$3,$4)", users.User_Name, users.Name, users.Surname, users.Is_Active)
	peopleByte, err := json.MarshalIndent(users, "", "\t")
	isvalid.CheckErr(err)
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(users)
	}
	defer db.Close()

	return c.Status(http.StatusOK).JSON(users)
}

func UsersGet(c *fiber.Ctx) error {
	var users model.User

	rows, err := db.Query("SELECT * FROM users ORDER BY id DESC")

	var usersAll []model.User

	for rows.Next() {
		statusRows := rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)
		if statusRows == nil {
			fmt.Println(users.Id, users.User_Name, users.Name, users.Surname, users.Is_Active)
		}
		usersAll = append(usersAll, users)

	}
	peopleByte, err := json.MarshalIndent(usersAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(usersAll)
	}
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(usersAll)

}
func UsersActiveGet(c *fiber.Ctx) error {
	var users model.User

	isActive := c.FormValue("is_Active")

	users.Is_Active = isActive

	rows, err := db.Query("SELECT * FROM users WHERE is_active=$1", users.Is_Active)

	var usersAll []model.User

	for rows.Next() {
		statusRows := rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)
		if statusRows == nil {
			fmt.Println(users.Id, users.User_Name, users.Name, users.Surname, users.Is_Active)
		}
		usersAll = append(usersAll, users)

	}
	peopleByte, err := json.MarshalIndent(usersAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(usersAll)
	}
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(usersAll)

}
func UsersEntry(c *fiber.Ctx) error {
	var userEntry model.UserEntry
	var userEntryAll []model.UserEntry

	userEntryId := c.FormValue("user_id")
	userInt, _ := strconv.Atoi(userEntryId)
	userEntry.Id = userInt

	rowsOne, err := db.Query("SELECT * FROM users WHERE id=$1", userEntry.Id)

	for rowsOne.Next() {
		rowsOne.Scan(&userEntry.Id, &userEntry.User_Name, &userEntry.Name, &userEntry.Surname, &userEntry.Is_Active)

	}
	rows, err := db.Query("SELECT * FROM entry WHERE assig=$1", userEntry.Id)

	for rows.Next() {
		rows.Scan(&userEntry.Entry.Id, &userEntry.Entry.Entry_Code, &userEntry.Entry.Entry_Title, &userEntry.Entry.Create_Date, &userEntry.Entry.Update_Date, &userEntry.Entry.Status, &userEntry.Entry.Assig, &userEntry.Entry.Tag)

	}
	rows2, err := db.Query("SELECT * FROM entryrelation WHERE main_entry=$1", userEntry.Entry.Entry_Code)

	for rows2.Next() {
		rows2.Scan(&userEntry.Entry.EntryRelation.Id, &userEntry.Entry.EntryRelation.Main_Entry, &userEntry.Entry.EntryRelation.Sub_Entry, &userEntry.Entry.EntryRelation.Parent_Entry)

		userEntryAll = append(userEntryAll, userEntry)
		log.Println(userEntryAll)

	}

	peopleByte, err := json.MarshalIndent(userEntryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(userEntryAll)
	}
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	defer rows2.Close()
	defer rowsOne.Close()

	return c.Status(http.StatusOK).JSON(userEntryAll)
}
