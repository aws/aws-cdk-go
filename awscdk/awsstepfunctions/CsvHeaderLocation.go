package awsstepfunctions


// CSV header location options.
type CsvHeaderLocation string

const (
	// Headers will be read from first row of CSV file.
	CsvHeaderLocation_FIRST_ROW CsvHeaderLocation = "FIRST_ROW"
	// Headers are provided in CSVHeaders property.
	CsvHeaderLocation_GIVEN CsvHeaderLocation = "GIVEN"
)

