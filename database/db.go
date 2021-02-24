package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB(dbHost, dbUser, dbPassword, dbName string) *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 dbUser,
		Passwd:               dbPassword,
		Net:                  "tcp",
		Addr:                 dbHost,
		DBName:               dbName,
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
			"charset":   "utf8",
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
