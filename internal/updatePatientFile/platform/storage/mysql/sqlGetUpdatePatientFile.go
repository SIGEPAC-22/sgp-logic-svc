package mysql

import "time"

type sqlGetUpdatePatientFile struct {
	StatePatient int        `db:"pfl_id_state_patient"`
	HighDate     *time.Time `db:"pfl_high_date"`
	LowDate      *time.Time `db:"pfl_low_date"`
}

type SqlGetSymptomPatient struct {
	IdSymptom int `db:"fhs_id_symptom"`
}
