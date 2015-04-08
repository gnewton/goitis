package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

var db *sqlx.DB

//const DB_FILE = "./data/itisSqlite/ITIS.sqlite"

const DB_FILE = "file:///home/newtong/gocode/src/github.com/gnewton/goitis/data/itisSqlite/ITIS.sqlite?mode=ro"

func tu(c web.C, w http.ResponseWriter, r *http.Request) {
	tsn := c.URLParams["tsn"]
	tu, err := getTaxonomicUnits2(tsn)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	if tu == nil {
		w.WriteHeader(404)
		w.Write([]byte(tsn + " not found"))
		return
	}

	b, err := json.Marshal(tu)
	w.Write(b)
}

func main() {
	if true {
		tu := new(TaxonomicUnit)
		indexAsKeyValue(tu)
		return
	}
	var err error
	db, err = connect()
	if err != nil {
		log.Fatal(err)
	}
	goji.Get("/taxonomicunits/:tsn", tu)
	goji.Serve()
}

func connect() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", DB_FILE)
}

// func getTaxonomicUnits(tsn string) (*TaxonomicUnit, error) {

// 	tu := TaxonomicUnit{}

// 	//err := baseDb.Get(&tu, "select * from taxonomic_units where tsn="+tsn)
// 	sql := "select * from taxonomic_units where tsn=" + tsn
// 	fmt.Println(sql)
// 	err := baseDb.Get(&tu, sql)

// 	if err != nil && err.Error() == "sql: no rows in result set" {
// 		return nil, nil
// 	}

// 	return &tu, err

// }

func getTaxonomicUnits2(tsn string) (*TaxonomicUnit, error) {

	tu := TaxonomicUnit{}
	db, err := connect()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Preparex("select * from taxonomic_units where tsn=?")
	defer stmt.Close()
	fmt.Println(*stmt)
	row := stmt.QueryRowx(tsn)

	err = row.StructScan(&tu)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
	}

	if tu.Update_date_tmp != nil {
		//tu.Update_date = new(string)
		tmp := string(tu.Update_date_tmp)
		tu.Update_date = &tmp
	}
	return &tu, err

}
