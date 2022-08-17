package database_test

import (
	"log"
	"main/database"
	"os"
	"testing"
)

func TestDB(m *testing.M) {
	db := database.Openconnention()
	if db != nil {
		log.Fatal("cannot connect to db:", db)
	}

	os.Exit(m.Run())
}
