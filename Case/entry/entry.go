package entry

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/isvalid"
	"main/model"
	"main/tag"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var db = database.Openconnention()

func TimeNew() int {
	currentTime := time.Now()

	a, _ := fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())
	return a
}
func EntryUpdate(c *fiber.Ctx) error {
	var entry model.Entry
	entry.Entry_Title = c.FormValue("entry_Title")
	statusT := c.FormValue("status")
	assigT := c.FormValue("assig")
	entryid := c.FormValue("userid")
	intEntryid, _ := strconv.Atoi(entryid)
	entry.Id = intEntryid
	entry.Assig = assigT
	entry.Status = statusT
	entry.Update_Date = time.Now()

	db.Exec("UPDATE entry SET entry_title=$1,update_date=$2, status=$3, assig=$4 WHERE id=$5", entry.Entry_Title, entry.Update_Date, entry.Status, entry.Assig, entry.Id)

	fmt.Println(entry.Id, entry.Update_Date, entry.Status, entry.Assig)

	peopleByte, _ := json.MarshalIndent(entry, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entry)
	}

	defer db.Close()

	return c.Status(http.StatusOK).JSON(entry)

}
func EntryRandomCode() string {
	rand.Seed(time.Now().Unix())
	permutation := rand.Int()
	str := strconv.Itoa(permutation)
	return str
}

func EntryCreate(c *fiber.Ctx) error {
	var entrycontrol model.EntryControl
	var tagPro model.TagProperties
	var entry model.Entry
	var users model.User

	entry.Entry_Title = c.FormValue("entry_Title")
	entry.Status = c.FormValue("status")
	entry.Assig = c.FormValue("assig")
	entry.Tag = c.FormValue("tag")

	if entry.Tag == "1000" || entry.Tag == "1001" || entry.Tag == "1002" {

		entry.Entry_Code = EntryRandomCode()
		//
		entry.Create_Date = time.Now()
		entry.Update_Date = time.Now()

		rows, _ := db.Query("SELECT * FROM entry")
		for rows.Next() {
			rows.Scan(&entrycontrol.Id, &entrycontrol.Entry_Title, &entrycontrol.Entry_Code, &entry.Entry_Title, &entrycontrol.Create_Date, &entrycontrol.Update_Date, &entrycontrol.Status, &entrycontrol.Assig, &entrycontrol.Tag)

		}

		if entrycontrol.Entry_Code == entry.Entry_Code {
			fmt.Println("username is used")
		} else if entrycontrol.Entry_Title == entry.Entry_Title {
			fmt.Println("title is used")
		} else if entrycontrol.Entry_Code == entry.Entry_Code {
			entry.Entry_Code = EntryRandomCode()
		} else {

			go tag.Tagprocreate(entry.Entry_Code, entry.Tag)
			db.Exec("INSERT INTO entry(entry_code,entry_title,create_date,update_date,status,assig,tag) VALUES($1,$2,$3,$4,$5,$6,$7)", entry.Entry_Code, entry.Entry_Title, entry.Create_Date, entry.Update_Date, entry.Status, entry.Assig, entry.Tag)

			tagPro.Entry_id = entrycontrol.Entry_Code
			tagPro.Tag_id = entrycontrol.Tag
			var assiguser string
			go UserControlGet()
			userCnt := UserControlGet()
			userCntStr := strconv.Itoa(userCnt)
			assiguser = userCntStr
			if entry.Assig == assiguser {
				assiguserInt, _ := strconv.Atoi(assiguser)
				users.Id = assiguserInt
				users.Is_Active = "True"
				db.Exec("UPDATE users SET is_active=$1 WHERE id=$2", users.Is_Active, users.Id)

				if entry.Tag == "1000" {
					var tagArg model.TagArge
					entry.Entry_Code = tagArg.Arge
					db.Exec("INSERT INTO tagarge(arge) VALUES(" + tagArg.Arge + ")")
					log.Println("tagArge Insert ")
				} else if entry.Tag == "1001" {
					db.Exec("INSERT INTO tagaccounting(accounting) VALUES(" + entry.Entry_Code + ")")

					log.Println("tagaccounting Insert ")

				} else if entry.Tag == "1002" {
					db.Exec("INSERT INTO tagtechnical(technical) VALUES(" + entry.Entry_Code + ")")
					log.Println("tagTechnical Insert ")

				}
			}
			defer rows.Close()
			defer db.Close()
		}
	} else {
		log.Fatalln("Tag category is wrong")
	}

	return c.Status(http.StatusOK).JSON(entry)

}

func EntryGet(c *fiber.Ctx) error {
	var entry model.Entry
	rows, _ := db.Query("SELECT * FROM entry ORDER BY id DESC")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Assig)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryAll)

}
func UserControlGet() int {
	var users model.User

	rows, _ := db.Query("SELECT * FROM users")

	var usersAll []model.User

	for rows.Next() {
		statusRows := rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)
		if statusRows == nil {
			fmt.Println(users.Id, users.User_Name, users.Name, users.Surname, users.Is_Active)
		}
		usersAll = append(usersAll, users)

	}

	defer db.Close()
	defer rows.Close()
	return users.Id
}
func EntryStatusGet(c *fiber.Ctx) error {
	var entry model.Entry
	get := c.FormValue("Get")
	entry.Status = get
	rows, _ := db.Query("SELECT * FROM entry WHERE status=$1", entry.Status)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryAll)

}
func EntryTimeCreGet(c *fiber.Ctx) error {
	var entry model.Entry
	create_Date_Start := c.FormValue("create_Date_Start")
	create_Date_End := c.FormValue("create_Date_End")

	rows, _ := db.Query("SELECT * FROM entry WHERE create_date BETWEEN '" + create_Date_Start + "' AND '" + create_Date_End + "'")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryAll)

}
func EntryTimeUpdGet(c *fiber.Ctx) error {
	var entry model.Entry
	create_Date_Start := c.FormValue("create_Date_Start")
	create_Date_End := c.FormValue("create_Date_End")

	rows, _ := db.Query("SELECT * FROM entry WHERE update_date BETWEEN '" + create_Date_Start + "' AND '" + create_Date_End + "'")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryAll)

}

func EntryAllGet(c *fiber.Ctx) error {
	var entryaa model.Entry

	get := c.FormValue("entry_Code")
	entryaa.Entry_Code = get
	rows, _ := db.Query("SELECT * FROM entry WHERE entry_code=$1", entryaa.Entry_Code)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entryaa.Id, &entryaa.Entry_Code, &entryaa.Entry_Title, &entryaa.Create_Date, &entryaa.Update_Date, &entryaa.Status, &entryaa.Assig, &entryaa.Tag, &entryaa.EntryRelation)

		entryAll = append(entryAll, entryaa)

	}

	entryaa.EntryRelation.Main_Entry = get

	rows2, _ := db.Query("SELECT * FROM entryrelation WHERE main_entry=$1", entryaa.EntryRelation.Main_Entry)

	for rows2.Next() {
		rows2.Scan(&entryaa.EntryRelation.Id, &entryaa.EntryRelation.Main_Entry, &entryaa.EntryRelation.Sub_Entry, &entryaa.EntryRelation.Parent_Entry)

		entryAll = append(entryAll, entryaa)
		log.Println(entryAll)
	}

	if entryaa.EntryRelation.Main_Entry == entryaa.Entry_Code {

		peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
		if err := c.BodyParser(&peopleByte); err != nil {
			return c.Status(http.StatusBadRequest).JSON(entryAll)
		}
	}

	return c.Status(http.StatusOK).JSON(rows)

}
func EntryTagGet(c *fiber.Ctx) error {
	var entry model.Entry
	entry.Tag = c.FormValue("tag")

	rows, _ := db.Query("SELECT * FROM entry WHERE tag=$1", entry.Tag)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryAll)

}

func EntryTopAllGet(c *fiber.Ctx) error {
	var entry model.EntryAll
	entry.Entry_Code = c.FormValue("entry_Code")

	rows, err := db.Query("SELECT * FROM entry WHERE entry_code=$1", entry.Entry_Code)

	var entryAll []model.EntryAll

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag)

	}

	rows2, err := db.Query("SELECT * FROM entryrelation WHERE main_entry=$1", entry.Entry_Code)

	for rows2.Next() {
		rows2.Scan(&entry.EntryRelation.Id, &entry.EntryRelation.Main_Entry, &entry.EntryRelation.Sub_Entry, &entry.EntryRelation.Parent_Entry)

	}

	rows3, err := db.Query("SELECT * FROM users WHERE id=$1", entry.Assig)

	for rows3.Next() {
		rows3.Scan(&entry.User.Id, &entry.User.User_Name, &entry.User.Name, &entry.User.Surname, &entry.User.Is_Active)

	}
	rows4, err := db.Query("SELECT * FROM tagpro WHERE tag_code=$1", entry.Tag)

	for rows4.Next() {
		rows4.Scan(&entry.TagProperties.Id, &entry.TagProperties.Entry_id, &entry.TagProperties.Tag_id)

	}

	rowss, err := db.Query("SELECT * FROM comments WHERE entry_code=$1", entry.Entry_Code)

	for rowss.Next() {
		rowss.Scan(&entry.EntryComment.Id, &entry.EntryComment.Entry_id, &entry.EntryComment.User_id, &entry.EntryComment.Text, &entry.EntryComment.Create_Date, &entry.EntryComment.Update_Date, &entry.EntryComment.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	isvalid.CheckErr(err)
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryAll)

	}

	defer db.Close()
	defer rows.Close()
	defer rows2.Close()
	defer rows3.Close()
	defer rows4.Close()
	defer rowss.Close()
	return c.Status(http.StatusOK).JSON(entryAll)
}
