package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nasum/spin/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Connection *gorm.DB

func Init() *gorm.DB {
	url := constants.GetEnv().DataBaseURL
	sqlDB, err := sql.Open("postgres", url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	conn, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	Connection = conn

	return conn
}

func Close() {
	db, err := Connection.DB()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}

	db.Close()
}
