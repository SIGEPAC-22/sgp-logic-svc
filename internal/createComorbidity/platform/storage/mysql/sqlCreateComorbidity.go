package mysql

type SqlCreateConmorbility struct {
	NameComorbidity        string `db:"cby_name_comorbidity"`
	DescriptionComorbidity string `db:"cby_comorbidity_description"`
}
