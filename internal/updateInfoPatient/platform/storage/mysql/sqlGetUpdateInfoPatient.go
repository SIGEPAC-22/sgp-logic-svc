package mysql

type sqlGetUpdateInfoPatient struct {
	DocumentType int `db:"pat_id_document_type"`
	Department   int `db:"pat_id_department"`
	Pregnant     int `db:"pfl_pregnant"`
}
