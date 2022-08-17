package entrycomment

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/isvalid"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var db = database.Openconnention()

func EntryComCreate(c *fiber.Ctx) error {
	var entryComment model.EntryComment
	var users model.User

	entryComment.Entry_id = c.FormValue("entry_id")

	entryComment.User_id = c.FormValue("user_id")

	entryComment.Text = c.FormValue("text")
	den := "True"
	entryComment.Is_Active = den

	entryComment.Create_Date = time.Now()

	rows, err := db.Query("SELECT * FROM users")

	var usersAll []model.User

	for rows.Next() {
		rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)

		usersAll = append(usersAll, users)

	}
	entryCommentuser, err := strconv.Atoi(entryComment.User_id)
	if entryCommentuser == users.Id {
	} else {
		log.Println("User not found")
	}
	db.Exec("INSERT INTO comments(entry_code, user_id, text, create_date, update_date, is_active) VALUES($1,$2,$3,$4,$5,$6)", entryComment.Entry_id, entryComment.User_id, entryComment.Text, entryComment.Create_Date, entryComment.Update_Date, entryComment.Is_Active)
	peopleByte, err := json.MarshalIndent(entryComment, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComment)
	}
	isvalid.CheckErr(err)
	defer db.Close()

	return c.Status(http.StatusOK).JSON(entryComment)

}
func EntryComGet(c *fiber.Ctx) error {
	var entryComment model.EntryComment
	rows, err := db.Query("SELECT * FROM comments ORDER BY id DESC")

	var entryComAll []model.EntryComment

	for rows.Next() {
		rows.Scan(&entryComment.Id, &entryComment.Entry_id, &entryComment.User_id, &entryComment.Text, &entryComment.Create_Date, &entryComment.Update_Date, &entryComment.Is_Active)

		entryComAll = append(entryComAll, entryComment)
		log.Println(entryComAll)
	}
	peopleByte, err := json.MarshalIndent(entryComAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComAll)
	}
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryComAll)

}

func EntryTrueGet(c *fiber.Ctx) error {
	var entryComment model.EntryComment

	rows, err := db.Query("SELECT * FROM comments WHERE is_active='True'")

	var entryComAll []model.EntryComment

	for rows.Next() {
		rows.Scan(&entryComment.Id, &entryComment.Entry_id, &entryComment.User_id, &entryComment.Text, &entryComment.Create_Date, &entryComment.Update_Date, &entryComment.Is_Active)

		entryComAll = append(entryComAll, entryComment)
		log.Println(entryComAll)

	}
	peopleByte, err := json.MarshalIndent(entryComAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComAll)
	}
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryComAll)

}

func EntryComRemove(c *fiber.Ctx) error {
	var entryComment model.EntryComment

	entryCommentEntry_id := c.FormValue("entry_Id")
	entryComment.Entry_id = entryCommentEntry_id

	entryCommentUser_id := c.FormValue("user_Id")
	comInt, _ := strconv.Atoi(entryCommentUser_id)
	entryComment.Id = comInt

	entryComment.Is_Active = "False"

	entryComment.Update_Date = time.Now()

	db.Exec("UPDATE comments SET user_id=$1, update_date=$2, is_active=$3 WHERE id=$4", entryComment.User_id, entryComment.Update_Date, entryComment.Is_Active, entryComment.Id)

	fmt.Println(entryComment.Id, entryComment.Entry_id, entryComment.User_id, entryComment.Text, entryComment.Create_Date, entryComment.Update_Date, entryComment.Is_Active)

	peopleByte, err := json.MarshalIndent(entryComment, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON("peopleByte")
	}
	isvalid.CheckErr(err)
	defer db.Close()

	return c.Status(http.StatusOK).JSON(entryComment)

}
