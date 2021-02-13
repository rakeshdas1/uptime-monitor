package main

import (
	"os"
	"testing"
	"log"
)

var a App
func TestMain(m *testing.M) {
	a.Initialize(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	ensureTableExists()
	code := m.Run()
	os.Exit(code)
}

func ensureTableExists() {
	const tableCheckQuery = "SELECT to_regclass('public.security_uptime');"
	if _, err := a.DB.Exec(tableCheckQuery); err != nil {
		log.Fatal(err)
	}
}