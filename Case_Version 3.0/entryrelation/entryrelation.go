package entryrelation

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

func EntryeReltMainGet(w http.ResponseWriter, r *http.Request) {
	var entryrelt model.EntryRelation
	main_entry := r.FormValue("main_entry")
	entryrelt.Main_Entry = main_entry

	rows, err := db.Query("SELECT * FROM entryrelation WHERE main_entry=$1", entryrelt.Main_Entry)

	var entryReltAll []model.EntryRelation

	for rows.Next() {
		rows.Scan(&entryrelt.Id, &entryrelt.Main_Entry, &entryrelt.Sub_Entry, &entryrelt.Parent_Entry)

		entryReltAll = append(entryReltAll, entryrelt)
		log.Println(entryReltAll)

	}
	peopleByte, err := json.MarshalIndent(entryReltAll, "", "\t")
	w.Write(peopleByte)
	isvalid.CheckErr(err)
	defer rows.Close()

}
func EntryeReltGet(w http.ResponseWriter, r *http.Request) {
	var entryrelt model.EntryRelation
	rows, err := db.Query("SELECT * FROM entryrelation ORDER BY id DESC")

	var entryReltAll []model.EntryRelation

	for rows.Next() {
		rows.Scan(&entryrelt.Id, &entryrelt.Main_Entry, &entryrelt.Sub_Entry, &entryrelt.Parent_Entry)

		entryReltAll = append(entryReltAll, entryrelt)
		log.Println(entryReltAll)
	}
	peopleByte, err := json.MarshalIndent(entryReltAll, "", "\t")
	w.Write(peopleByte)
	isvalid.CheckErr(err)
	defer db.Close()
	defer rows.Close()

}

func EntryReltUpdate(w http.ResponseWriter, r *http.Request) {
	var entryrelt model.EntryRelation

	entryrelt.Main_Entry = r.FormValue("main_Entry")
	entryrelt.Sub_Entry = r.FormValue("sub_Entry")
	entryrelt.Parent_Entry = r.FormValue("parent_Entry")
	entryreltId := r.FormValue("userid")
	intVar, _ := strconv.Atoi(entryreltId)
	entryrelt.Id = intVar
	db.Exec("UPDATE entryrelation SET main_entry=$1, sub_entry=$2, parent_entry=$3 WHERE id=$4", entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry, entryrelt.Id)

	fmt.Println(entryrelt.Id, entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry)

	peopleByte, err := json.MarshalIndent(entryrelt, "", "\t")
	w.Write(peopleByte)
	isvalid.CheckErr(err)
	defer db.Close()

}

func EntryReltCreate(w http.ResponseWriter, r *http.Request) {
	var entryrelt model.EntryRelation

	entryrelt.Main_Entry = r.FormValue("Main_Entry")
	entryrelt.Sub_Entry = r.FormValue("Sub_Entry")
	entryrelt.Parent_Entry = r.FormValue("Parent_Entry")
	switch {
	case entryrelt.Main_Entry == entryrelt.Sub_Entry:
		log.Println("main and sub cannot be the same")
	case entryrelt.Parent_Entry == entryrelt.Main_Entry:
		log.Println("parent and main cannot be the same")
	case entryrelt.Parent_Entry == entryrelt.Sub_Entry:
		log.Println("parent and sub cannot be the same")
	default:

		db.Exec("INSERT INTO entryrelation(main_entry,sub_entry,parent_entry) VALUES($1,$2,$3)", entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry)
		peopleByte, err := json.MarshalIndent(entryrelt, "", "\t")

		w.Write(peopleByte)
		isvalid.CheckErr(err)
		defer db.Close()
	}

}
