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
)

var db = database.Openconnention()

func StatusCreate(w http.ResponseWriter, r *http.Request) {
	var status model.Status
	status.Name = r.FormValue("name")

	db.Exec("INSERT INTO status(name) VALUES($1)", status.Name)
	peopleByte, err := json.MarshalIndent(status, "", "\t")
	log.Println(status)
	w.Write(peopleByte)
	isvalid.CheckErr(err)

}
func StatusGet(w http.ResponseWriter, r *http.Request) {
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
	w.Write(peopleByte)
	log.Println(status)
	isvalid.CheckErr(err)

	defer rows.Close()

}
func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	var status model.Status
	//entry.Entry_Code = c.FormValue("entry_Code")
	status.Name = r.FormValue("name")
	statusid := r.FormValue("id")
	intVar, _ := strconv.Atoi(statusid)
	status.Id = intVar

	db.Exec("UPDATE status SET name=$1 WHERE id=$2", status.Name, status.Id)

	log.Println(status.Id, status.Name)

	peopleByte, err := json.MarshalIndent(status, "", "\t")
	w.Write(peopleByte)
	log.Println(status)
	isvalid.CheckErr(err)

}
