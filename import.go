package importdata

import (
	"importdata/csv"
	"importdata/sav"
)

type ImportFile interface {
	Import(fileName string) int
}

type HandleImport func(fileName string) int

func (imp HandleImport) Import(fileName string) int {
	return imp(fileName)
}

var ImportCsvFile = HandleImport(csv.Import)
var ImportSavFile = HandleImport(sav.Import)
