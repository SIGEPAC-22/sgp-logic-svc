package mysql

type SqlCreateComorbidity struct {
	NameComorbidity        string `db:"cby_name_comorbidity"`
	DescriptionComorbidity string `db:"cby_comorbidity_description"`
}
