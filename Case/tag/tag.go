package tag

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/isvalid"
	"main/model"
)

var db = database.Openconnention()

func Tagprocreate(a string, b string) {
	var tagPro model.TagProperties
	tagPro.Entry_id = a
	tagPro.Tag_id = b

	db.Exec("INSERT INTO tagpro(entry_code,tag_code) VALUES($1,$2)", tagPro.Entry_id, tagPro.Tag_id)
	peopleByte, err := json.MarshalIndent(tagPro, "", "\t")
	isvalid.CheckErr(err)
	fmt.Println(peopleByte)
	log.Println(tagPro)

}
