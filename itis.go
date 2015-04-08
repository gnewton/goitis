package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

type IsTable interface {
	Name() string
	KeyFieldName() string
	Prepare(interface{})
	Clear() bool
}

type IsTable2 struct {
}

func (it *IsTable2) KeyFieldName() string {
	return "Tsn"
}

func (it *IsTable2) Name() string {
	return "taxonomic_units"
}

func (it *IsTable2) Prepare(interface{}) {
	// if .Update_date_tmp != nil {
	// 	tmp := string(t.Update_date_tmp)
	// 	t.Update_date = &tmp
	// }
}

func (t *TaxonomicUnit) Instance() interface{} {
	return new(TaxonomicUnit)
}

func (t *TaxonomicUnit) KeyFieldName() string {
	return "Tsn"
}

func (t *TaxonomicUnit) Clear() bool {
	t.Tsn = 0
	t.Complete_name = nil
	t.Completeness_rtng = nil
	t.Credibility_rtng = nil
	t.Currency_rating = nil
	t.Hybrid_author_id = 0
	t.Initial_time_stamp = nil
	t.Kingdom_id = 0
	t.N_usage = nil
	t.Name_usage = nil
	t.Parent_tsn = nil
	t.Phylo_sort_seq = 0
	t.Rank_id = 0
	t.Taxon_author_id = 0
	t.Unaccept_reason = nil
	t.Uncertain_prnt_ind = nil
	t.Unit_ind1 = nil
	t.Unit_ind2 = nil
	t.Unit_ind3 = nil
	t.Unit_ind4 = nil
	t.Unit_name1 = nil
	t.Unit_name2 = nil
	t.Unit_name3 = nil
	t.Unit_name4 = nil
	t.Unnamed_taxon_ind = nil
	t.Update_date = nil
	tmp := new([]uint8)
	t.Update_date_tmp = *tmp
	t.Update_string = nil
	return false
}

func (t *TaxonomicUnit) Name() string {
	return "taxonomic_units"
}

func (t *TaxonomicUnit) Prepare(interface{}) {
	if t.Update_date_tmp != nil {
		tmp := string(t.Update_date_tmp)
		t.Update_date = &tmp
	}
}

type TaxonomicUnit struct {
	Tsn int64 `db:"tsn"   json:"tsn"`

	Complete_name      *string    `db:"complete_name"  json:"completeName,omitempty"`
	Completeness_rtng  *string    `db:"completeness_rtng"  json:",omitempty"`
	Credibility_rtng   *string    `db:"credibility_rtng"  json:",omitempty"`
	Currency_rating    *string    `db:"currency_rating"  json:"currencyRating,omitempty"`
	Hybrid_author_id   int64      `db:"hybrid_author_id"  json:",omitempty"`
	Initial_time_stamp *time.Time `db:"initial_time_stamp"  json:",omitempty"`
	Kingdom_id         uint8      `db:"kingdom_id"  json:",omitempty"`
	N_usage            *string    `db:"n_usage"  json:",omitempty"`
	Name_usage         *string    `db:"name_usage"  json:",omitempty"`
	Parent_tsn         *int64     `db:"parent_tsn"  json:",omitempty"`
	Phylo_sort_seq     int32      `db:"phylo_sort_seq"  json:",omitempty"`
	Rank_id            uint16     `db:"rank_id"  json:",omitempty"`
	Taxon_author_id    int64      `db:"taxon_author_id"  json:",omitempty"`
	Unaccept_reason    *string    `db:"unaccept_reason"  json:",omitempty"`
	Uncertain_prnt_ind *string    `db:"uncertain_prnt_ind"  json:",omitempty"`
	Unit_ind1          *string    `db:"unit_ind1" json:",omitempty"`
	Unit_ind2          *string    `db:"unit_ind2"  json:",omitempty"`
	Unit_ind3          *string    `db:"unit_ind3"  json:",omitempty"`
	Unit_ind4          *string    `db:"unit_ind4"  json:",omitempty"`
	Unit_name1         *string    `dbtsn,"unit_name1"  json:",omitempty"`
	Unit_name2         *string    `db:"unit_name2"  json:",omitempty"`
	Unit_name3         *string    `db:"unit_name3"  json:",omitempty"`
	Unit_name4         *string    `db:"unit_name4"  json:",omitempty"`
	Unnamed_taxon_ind  *string    `db:"unnamed_taxon_ind"  json:",omitempty"`
	Update_date        *string    `db:"update_date"  json:"updateDate,omitempty"`
	Update_date_tmp    []uint8    `db:"update_date"  json:"-"`
	Update_string      *string    `json:"-"`
}

func indexAsKeyValue(table IsTable) {

	db, err := connect()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "select * from " + table.Name()
	fmt.Println(sql)
	rows, err := db.Queryx(sql)
	if err != nil {
		log.Fatal(err)
	}
	barType := reflect.ValueOf(table).Elem().Type()

	for rows.Next() {
		v := reflect.New(barType).Interface()

		err = rows.StructScan(v)
		keyFieldName := reflect.ValueOf(v).Elem().FieldByName(table.KeyFieldName())
		fmt.Println(keyFieldName.Int())
		fmt.Println("==")

		var keyString string

		switch keyFieldName.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			keyString = strconv.FormatInt(keyFieldName.Int(), 10)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			keyString = strconv.FormatUint(keyFieldName.Uint(), 10)

		case reflect.String:
			keyString = string(keyFieldName.Bytes())
		}

		fmt.Println(keyString)

		b, _ := json.Marshal(v)
		fmt.Println(string(b))

	}
}
