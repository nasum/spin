package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nasum/spin/constants"
	"github.com/nasum/spin/route"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	url := constants.GetEnv().DataBaseURL
	sqlDB, err := sql.Open("postgres", url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	db.AutoMigrate()

	route.Init()
}
