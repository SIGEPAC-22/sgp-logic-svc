package mysql

type SqlCreateSymptom struct {
	NameSymptom        string `db:"stm_name_symptons"`
	DescriptionSymptom string `db:"stm_sympton_description"`
}
