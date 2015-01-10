package main

import (
	//"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"strconv"
	"time"
	//	"os"
)

var baseDb *sqlx.DB

const DB_FILE = "./itisSqlite122314/ITIS.sqlite"

func main() {
	var err error
	baseDb, err = connect()
	if err != nil {
		log.Fatal(err)
	}
	const maxTsn = 955136
	for i := 1; i < maxTsn; i++ {
		randomTsn := rand.Intn(maxTsn)
		tu, err := getTaxonomicUnits(strconv.Itoa(randomTsn))
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println(i)
		if tu != nil {
			tu.Update_string = string(tu.Update_date)
			fmt.Printf("%+v\n", *tu)
		} else {
			fmt.Println("*")
		}
	}

}

func connect() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", DB_FILE)
}

func getTaxonomicUnits(tsn string) (*TaxonomicUnit, error) {

	tu := TaxonomicUnit{}

	//err := baseDb.Get(&tu, "select * from taxonomic_units where tsn="+tsn)
	err := baseDb.Get(&tu, "select * from taxonomic_units where tsn="+tsn)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	return &tu, err

}
