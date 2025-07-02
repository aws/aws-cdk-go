package awsstepfunctions


// Delimiter used in CSV file.
type CsvDelimiter string

const (
	// Comma delimiter.
	CsvDelimiter_COMMA CsvDelimiter = "COMMA"
	// Pipe delimiter.
	CsvDelimiter_PIPE CsvDelimiter = "PIPE"
	// Semicolon delimiter.
	CsvDelimiter_SEMICOLON CsvDelimiter = "SEMICOLON"
	// Space delimiter.
	CsvDelimiter_SPACE CsvDelimiter = "SPACE"
	// Tab delimiter.
	CsvDelimiter_TAB CsvDelimiter = "TAB"
)

