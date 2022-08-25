package model

import "time"

type (
	EntrySub struct {
		Id          int
		Entry_Code  string
		Entry_Title string
		Create_Date time.Time
		Update_Date time.Time
		Status      string
		Assig       string
		Tag         string
		Is_Active   string
		EntryRelationSub
		User
		TagProperties
		EntryComment
	}
	EntrySubNot struct {
		Id_         int
		Entry_Code_  string
		Entry_Title_ string
		Create_Date_ time.Time
		Update_Date_ time.Time
		Status_      string
		Assig_       string
		Tag_         string
		Is_Active_   string
	}
	EntryRelationSub struct {
		Id         int
		Main_Entry string
		Sub_Entry  string
		EntrySubNot
		Parent_Entry string
	}
)
type (
	Entry struct {
		Id          int
		Entry_Code  string
		Entry_Title string
		Create_Date time.Time
		Update_Date time.Time
		Status      string
		Assig       string
		Tag         string
		Is_Active   string
		EntryRelation
	}

	User struct {
		Id        int
		User_Name string
		Name      string
		Surname   string
		Is_Active string
	}

	Status struct {
		Id   int
		Name string
	}

	EntryRelation struct {
		Id           int
		Main_Entry   string
		Sub_Entry    string
		Parent_Entry string
	}
	TagArge struct {
		Id   int
		Arge string
	}

	TagAccounting struct {
		Id         int
		Accounting string
	}

	TagTechnical struct {
		Id        int
		Technical string
	}

	TagProperties struct {
		Id         int
		Entry_code string
		Tag_code   string
	}

	EntryComment struct {
		Id          int
		Entry_id    string
		User_id     string
		Text        string
		Create_Date time.Time
		Update_Date time.Time
		Is_Active   string
	}

	EntryAll struct {
		Id          int
		Entry_Code  string
		Entry_Title string
		Create_Date time.Time
		Update_Date time.Time
		Status      string
		Assig       string
		User
		Tag       string
		Is_Active string
		TagProperties
		EntryRelation
		EntryComment
	}

	UserEntry struct {
		Id        int
		User_Name string
		Name      string
		Surname   string
		Is_Active string
		Entry
	}
)
