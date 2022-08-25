package entry

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/isvalid"
	"main/model"
	"net/http"
	"strconv"
	"time"
)

var db = database.Openconnention()

func EntryUpdate(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	entryid := r.FormValue("entry_id")
	entry.Entry_Title = r.FormValue("entry_Title")
	statusT := r.FormValue("status")
	assigT := r.FormValue("assig")
	entry.Tag = r.FormValue("tag")

	if statusT == "Done" {
		entry.Is_Active = "False"
	} else {
		entry.Is_Active = "True"

	}
	intEntryid, _ := strconv.Atoi(entryid)
	entry.Id = intEntryid
	entry.Assig = assigT
	entry.Status = statusT
	entry.Update_Date = time.Now()

	db.Exec("UPDATE entry SET entry_title=$1,update_date=$2, status=$3, assig=$4,tag=$5,is_active=$6 WHERE id=$7", entry.Entry_Title, entry.Update_Date, entry.Status, entry.Assig, entry.Tag, entry.Is_Active, entry.Id)

	fmt.Println(entry.Id, entry.Update_Date, entry.Status, entry.Assig)

	peopleByte, _ := json.MarshalIndent(entry, "", "\t")
	w.Write(peopleByte)

	defer db.Close()

}
func EntryRandomCode(srcs []byte) string {

	src := []byte(srcs)

	s := hex.EncodeToString(src)

	return s
}

func EntryCreate(w http.ResponseWriter, r *http.Request) {
	var (
		tagPro model.TagProperties
		entry  model.Entry

		users    model.User
		statu    model.Status
		con      = false
		conS     = false
		conA     = false
		conC     = false
		statuAll []model.Status
		usersAll []model.User
	)

	db := database.Openconnention()

	entry.Entry_Title = r.FormValue("entry_Title")
	entry.Status = r.FormValue("status")
	entry.Assig = r.FormValue("assig")
	entry.Tag = r.FormValue("tag")
	entry.Entry_Code = EntryRandomCode([]byte(entry.Entry_Title))

	//----------------------------------------------------------------------------------------

	row, err := db.Query("SELECT * FROM status")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		row.Scan(&statu.Id, &statu.Name)
		statuAll = append(statuAll, statu)
		if entry.Status == statu.Name {

			conS = true
		}
	}

	//----------------------------------------------------------------------------------------
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)
		usersAll = append(usersAll, users)
		eAsg := strconv.Itoa(users.Id)
		if entry.Assig == eAsg {

			conA = true
		}
	}
	//----------------------------------------------------------------------------------------

	rowss, err := db.Query("SELECT * FROM entry")
	if err != nil {
		log.Fatal(err)
	}
	var a, b, c, d, e, f, g, v string
	var rs int

	for rowss.Next() {
		rowss.Scan(&rs, &a, &b, &c, &d, &e, &f, &g, &v)

		if entry.Entry_Title == b {

			con = true
		}
		if entry.Entry_Code == a {

			conC = true
		}
	}

	if entry.Tag == "1000" || entry.Tag == "1001" || entry.Tag == "1002" {

		//----------------------------------------------------------------------------------------
		switch {
		case conC == true:
			fmt.Fprintf(w, "error entry code")
		case conS == false:
			fmt.Fprintf(w, "error statu")
		case conA == false:
			fmt.Fprintf(w, "error assig")
		case con == true:
			fmt.Fprintf(w, "error entry title")

		default:

			users.Is_Active = "True"
			db.Exec("UPDATE user SET is_Active=$1 WHERE id=$2", users.Is_Active, users.Id)

			entry.Create_Date = time.Now()
			if entry.Status == "Done" {
				entry.Is_Active = "False"
			} else {
				entry.Is_Active = "True"

			}
			db.Exec("INSERT INTO entry(entry_code,entry_title,create_date,update_date,status,assig,tag,is_active) VALUES($1,$2,$3,$4,$5,$6,$7,$8)", entry.Entry_Code, entry.Entry_Title, entry.Create_Date, entry.Update_Date, entry.Status, entry.Assig, entry.Tag, entry.Is_Active)
			Tagprocreate(entry.Entry_Code, entry.Tag)
			tagPro.Entry_code = entry.Entry_Code
			tagPro.Tag_code = entry.Tag

			peopleByte, _ := json.MarshalIndent(entry, "", "\t")

			users.Is_Active = "True"
			EA, _ := strconv.Atoi(entry.Assig)
			users.Id = EA
			db.Exec("UPDATE users SET is_active=$1 WHERE id=$2", users.Is_Active, users.Id)

			log.Println(users.Is_Active, users.Name, users.Surname)

			w.Write(peopleByte)
			switch true {
			case entry.Tag == "1000":
				tagarge(entry)
			case entry.Tag == "1001":
				tagaccounting(entry)
			case entry.Tag == "1002":
				tagtechnical(entry)
			}
		}
		defer row.Close()
		defer rowss.Close()
		defer rows.Close()
		defer db.Close()
	} else {
		fmt.Println("tag hatalı")
	}
}

func Tagprocreate(a string, b string) {
	var tagPro model.TagProperties
	tagPro.Entry_code = a
	tagPro.Tag_code = b

	_, err := db.Exec("INSERT INTO tagpro(entry_code,tag_code) VALUES($1,$2)", tagPro.Entry_code, tagPro.Tag_code)
	isvalid.CheckErr(err)

}

func tagaccounting(entry model.Entry) {

	_, err := db.Exec("INSERT INTO tagaccounting(accounting) VALUES(" + entry.Entry_Code + ")")
	isvalid.CheckErr(err)
	log.Println("tagaccounting Insert ")
}
func tagtechnical(entry model.Entry) {

	_, err := db.Exec("INSERT INTO tagtechnical(technical) VALUES(" + entry.Entry_Code + ")")
	log.Println("tagTechnical Insert ")
	isvalid.CheckErr(err)
}
func tagarge(entry model.Entry) {

	var tagArg model.TagArge
	entry.Entry_Code = tagArg.Arge
	_, err := db.Exec("INSERT INTO tagarge(arge) VALUES(" + tagArg.Arge + ")")
	isvalid.CheckErr(err)
	log.Println("tagArge Insert ")
}
func EntryGet(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	rows, _ := db.Query("SELECT * FROM entry ORDER BY id DESC")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Assig, &entry.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)

}
func EntryStatusGet(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	get := r.FormValue("Get")
	entry.Status = get
	rows, _ := db.Query("SELECT * FROM entry WHERE status=$1", entry.Status)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)

	defer rows.Close()

}
func EntryTimeCreGet(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	create_Date_Start := r.FormValue("create_Date_Start")
	create_Date_End := r.FormValue("create_Date_End")

	rows, _ := db.Query("SELECT * FROM entry WHERE create_date BETWEEN '" + create_Date_Start + "' AND '" + create_Date_End + "'")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)

	defer rows.Close()

}
func EntryTimeUpdGet(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	create_Date_Start := r.FormValue("create_Date_Start")
	create_Date_End := r.FormValue("create_Date_End")

	rows, _ := db.Query("SELECT * FROM entry WHERE update_date BETWEEN '" + create_Date_Start + "' AND '" + create_Date_End + "'")

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)

	defer rows.Close()

}
func EntryAllGet(w http.ResponseWriter, r *http.Request) {
	var entryaa model.Entry

	get := r.FormValue("entry_Code")
	entryaa.Entry_Code = get
	rows, _ := db.Query("SELECT * FROM entry WHERE entry_code=$1", entryaa.Entry_Code)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entryaa.Id, &entryaa.Entry_Code, &entryaa.Entry_Title, &entryaa.Create_Date, &entryaa.Update_Date, &entryaa.Status, &entryaa.Assig, &entryaa.Tag, &entryaa.EntryRelation, &entryaa.Is_Active)

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
		w.Write(peopleByte)
	}

}
func EntryTagGet(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	entry.Tag = r.FormValue("tag")

	rows, _ := db.Query("SELECT * FROM entry WHERE tag=$1", entry.Tag)

	var entryAll []model.Entry

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)

}
func EntryTopAllGet(w http.ResponseWriter, r *http.Request) {
	var entry model.EntryAll
	entry.Entry_Code = r.FormValue("entry_Code")

	rows, err := db.Query("SELECT * FROM entry WHERE entry_code=$1", entry.Entry_Code)

	var entryAll []model.EntryAll

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

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
		rows4.Scan(&entry.TagProperties.Id, &entry.TagProperties.Entry_code, &entry.TagProperties.Tag_code)

	}

	rowss, err := db.Query("SELECT * FROM comments WHERE entry_code=$1", entry.Entry_Code)

	for rowss.Next() {
		rowss.Scan(&entry.EntryComment.Id, &entry.EntryComment.Entry_id, &entry.EntryComment.User_id, &entry.EntryComment.Text, &entry.EntryComment.Create_Date, &entry.EntryComment.Update_Date, &entry.EntryComment.Is_Active)

		entryAll = append(entryAll, entry)
		log.Println(entryAll)
	}
	isvalid.CheckErr(err)
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)
}

func EntrySubFull(w http.ResponseWriter, r *http.Request) {
	var entry model.EntrySub
	entry.Entry_Code = r.FormValue("entry_Code")

	rows, err := db.Query("SELECT * FROM entry WHERE entry_code=$1", entry.Entry_Code)

	var entryAll []model.EntrySub

	for rows.Next() {
		rows.Scan(&entry.Id, &entry.Entry_Code, &entry.Entry_Title, &entry.Create_Date, &entry.Update_Date, &entry.Status, &entry.Assig, &entry.Tag, &entry.Is_Active)

	}

	rows2, err := db.Query("SELECT * FROM entryrelation WHERE main_entry=$1", entry.Entry_Code)

	for rows2.Next() {
		rows2.Scan(&entry.EntryRelationSub.Id, &entry.EntryRelationSub.Main_Entry, &entry.EntryRelationSub.Sub_Entry, &entry.EntryRelationSub.Parent_Entry)

	}

	rows5, err := db.Query("SELECT * FROM entry WHERE entry_code=$1", entry.EntryRelationSub.Sub_Entry)

	for rows5.Next() {
		rows5.Scan(&entry.EntryRelationSub.Id_, &entry.EntryRelationSub.Entry_Code_, &entry.EntryRelationSub.Entry_Title_, &entry.EntryRelationSub.Create_Date_, &entry.EntryRelationSub.Update_Date_, &entry.EntryRelationSub.Status_, &entry.EntryRelationSub.Assig_, &entry.EntryRelationSub.Tag_, &entry.EntryRelationSub.Is_Active_)

	}

	rows3, err := db.Query("SELECT * FROM users WHERE id=$1", entry.Assig)

	for rows3.Next() {
		rows3.Scan(&entry.User.Id, &entry.User.User_Name, &entry.User.Name, &entry.User.Surname, &entry.User.Is_Active)

	}
	rows4, err := db.Query("SELECT * FROM tagpro WHERE tag_code=$1", entry.Tag)

	for rows4.Next() {
		rows4.Scan(&entry.TagProperties.Id, &entry.TagProperties.Entry_code, &entry.TagProperties.Tag_code)

	}

	rowss, err := db.Query("SELECT * FROM comments WHERE entry_code=$1", entry.Entry_Code)

	for rowss.Next() {
		rowss.Scan(&entry.EntryComment.Id, &entry.EntryComment.Entry_id, &entry.EntryComment.User_id, &entry.EntryComment.Text, &entry.EntryComment.Create_Date, &entry.EntryComment.Update_Date, &entry.EntryComment.Is_Active)

		log.Println(entryAll)
	}
	entryAll = append(entryAll, entry)
	isvalid.CheckErr(err)
	peopleByte, _ := json.MarshalIndent(entryAll, "", "\t")
	w.Write(peopleByte)
}
