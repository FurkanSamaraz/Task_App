package handler

import (
	"main/entry"
	"main/entrycomment"
	"main/entryrelation"
	"main/status"
	"main/user"
	"net/http"
)

func Handlers() {

	http.HandleFunc("/EntryTopAllGet", entry.EntryTopAllGet)
	http.HandleFunc("/UsersEntryAllGet", user.UsersEntry)

	http.HandleFunc("/EntryComCreate", entrycomment.EntryComCreate)
	http.HandleFunc("/EntryComGet", entrycomment.EntryComGet)
	http.HandleFunc("/EntryComRemove", entrycomment.EntryComRemove)
	http.HandleFunc("/EntryTrueGet", entrycomment.EntryTrueGet)

	http.HandleFunc("/EntryAllGet", entry.EntryAllGet)
	http.HandleFunc("/EntryTagGet", entry.EntryTagGet)

	http.HandleFunc("/EntryGet", entry.EntryGet)
	http.HandleFunc("/EntryeReltGet", entryrelation.EntryeReltGet)
	http.HandleFunc("/StatusGet", status.StatusGet)
	http.HandleFunc("/UsersGet", user.UsersGet)

	http.HandleFunc("/EntryeReltMainGet", entryrelation.EntryeReltMainGet)
	http.HandleFunc("/EntryStatusGet", entry.EntryStatusGet)
	http.HandleFunc("/UsersActiveGet", user.UsersActiveGet)
	http.HandleFunc("/EntryTimeCreGet", entry.EntryTimeCreGet)
	http.HandleFunc("/EntryTimeUpdGet", entry.EntryTimeUpdGet)

	http.HandleFunc("/EntryUpdate", entry.EntryUpdate)
	http.HandleFunc("/EntryReltUpdate", entryrelation.EntryReltUpdate)
	http.HandleFunc("/StatusUpdate", status.StatusUpdate)

	http.HandleFunc("/EntryCreate", entry.EntryCreate)
	http.HandleFunc("/EntryReltCreate", entryrelation.EntryReltCreate)
	http.HandleFunc("/StatusCreate", status.StatusCreate)
	http.HandleFunc("/UserCreate", user.UserCreate)

	http.ListenAndServe(":8080", nil)
}
