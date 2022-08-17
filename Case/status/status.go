package status

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

func StatusCreate(c *fiber.Ctx) error {
	var status model.Status
	status.Name = c.FormValue("name")

	db.Exec("INSERT INTO status(name) VALUES($1)", status.Name)
	peopleByte, err := json.MarshalIndent(status, "", "\t")
	log.Println(status)

	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	isvalid.CheckErr(err)
	defer db.Close()

	return c.Status(http.StatusOK).JSON(status)
}
func StatusGet(c *fiber.Ctx) error {
	var status model.Status
	rows, err := db.Query("SELECT * FROM status ORDER BY id DESC")

	var statusAll []model.Status

	for rows.Next() {
		statusRows := rows.Scan(&status.Id, &status.Name)
		if statusRows == nil {
			fmt.Println(&status.Id, &status.Name)
		}
		statusAll = append(statusAll, status)

	}
	peopleByte, err := json.MarshalIndent(statusAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(statusAll)
	}
	log.Println(status)
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(statusAll)

}
func StatusUpdate(c *fiber.Ctx) error {
	var status model.Status
	//entry.Entry_Code = c.FormValue("entry_Code")
	status.Name = c.FormValue("name")
	statusid := c.FormValue("id")
	intVar, _ := strconv.Atoi(statusid)
	status.Id = intVar

	db.Exec("UPDATE status SET name=$1 WHERE id=$2", status.Name, status.Id)

	log.Println(status.Id, status.Name)

	peopleByte, err := json.MarshalIndent(status, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	log.Println(status)
	isvalid.CheckErr(err)
	defer db.Close()

	return c.Status(http.StatusOK).JSON(status)

}
