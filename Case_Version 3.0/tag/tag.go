package tag

import (
	"main/database"
	"main/isvalid"
	"main/model"
)

var db = database.Openconnention()

func Tagprocreate(a string, b string) {
	var tagPro model.TagProperties
	tagPro.Entry_code = a
	tagPro.Tag_code = b

	_, err := db.Exec("INSERT INTO tagpro(entry_code,tag_code) VALUES($1,$2)", tagPro.Entry_code, tagPro.Tag_code)
	isvalid.CheckErr(err)

}
