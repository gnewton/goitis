package main

type TaxonomicUnit struct {
	Tsn                int64      `db:"tsn"`
	Unit_ind1          *string    `db:"unit_ind1"`
	Unit_name1         *string    `dbtsn,"unit_name1"`
	Unit_ind2          *string    `db:"unit_ind2"`
	Unit_name2         *string    `db:"unit_name2"`
	Unit_ind3          *string    `db:"unit_ind3"`
	Unit_name3         *string    `db:"unit_name3"`
	Unit_ind4          *string    `db:"unit_ind4"`
	Unit_name4         *string    `db:"unit_name4"`
	Unnamed_taxon_ind  *string    `db:"unnamed_taxon_ind"`
	Name_usage         *string    `db:"name_usage"`
	Unaccept_reason    *string    `db:"unaccept_reason"`
	Credibility_rtng   *string    `db:"credibility_rtng"`
	Completeness_rtng  *string    `db:"completeness_rtng"`
	Currency_rating    *string    `db:"currency_rating"`
	Phylo_sort_seq     int32      `db:"phylo_sort_seq"`
	Initial_time_stamp *time.Time `db:"initial_time_stamp"`
	Parent_tsn         *int64     `db:"parent_tsn"`
	Taxon_author_id    int64      `db:"taxon_author_id"`
	Hybrid_author_id   int64      `db:"hybrid_author_id"`
	Kingdom_id         uint8      `db:"kingdom_id"`
	Rank_id            uint16     `db:"rank_id"`
	Update_date        []uint8    `db:"update_date"`
	//Update_date *time.Time `db:"update_date"`
	Update_string      string
	Uncertain_prnt_ind *string `db:"uncertain_prnt_ind"`
	N_usage            *string `db:"n_usage"`
	Complete_name      string  `db:"complete_name"`
}
