package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var a App

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file! %s", err)
	}
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
