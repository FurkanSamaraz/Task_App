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
)

var db = database.Openconnention()

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var users model.User

	users.User_Name = r.FormValue("User_Name")
	users.Name = r.FormValue("Name")
	users.Surname = r.FormValue("Surname")
	users.Is_Active = r.FormValue("Is_Active")

	db.Exec("INSERT INTO users(user_Name,name,surname,is_Active) VALUES($1,$2,$3,$4)", users.User_Name, users.Name, users.Surname, users.Is_Active)
	peopleByte, err := json.MarshalIndent(users, "", "\t")
	isvalid.CheckErr(err)
	w.Write(peopleByte)

}

func UsersGet(w http.ResponseWriter, r *http.Request) {
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
	w.Write(peopleByte)
	isvalid.CheckErr(err)
	defer rows.Close()

}
func UsersActiveGet(w http.ResponseWriter, r *http.Request) {
	var users model.User

	isActive := r.FormValue("is_Active")

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
	w.Write(peopleByte)
	isvalid.CheckErr(err)

	defer rows.Close()

}
func UsersEntry(w http.ResponseWriter, r *http.Request) {
	var userEntry model.UserEntry
	var userEntryAll []model.UserEntry

	userEntryId := r.FormValue("user_id")
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
	peopleByte, err := json.MarshalIndent(userEntry, "", "\t")
	w.Write(peopleByte)
	isvalid.CheckErr(err)

	defer rows.Close()
	defer rows2.Close()
	defer rowsOne.Close()

}
