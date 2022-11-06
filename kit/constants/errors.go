package constants

import "errors"

var ErrorDataError = errors.New("Error informacion incompleta")
var ErrorDataBase = errors.New("Error al comunicarse a la BD")
var ZeroRows = errors.New("zero rows, no data process")
