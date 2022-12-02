package mysql

type sqlGetUpdatePatientFile struct {
	StatePatient string `db:"pfl_id_state_patient"`
	HighDate     string `db:"pfl_high_date"`
	LowDate      string `db:"pfl_low_date"`
}
