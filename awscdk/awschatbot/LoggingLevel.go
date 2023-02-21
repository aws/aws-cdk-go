package awschatbot


// Logging levels include ERROR, INFO, or NONE.
type LoggingLevel string

const (
	// ERROR.
	LoggingLevel_ERROR LoggingLevel = "ERROR"
	// INFO.
	LoggingLevel_INFO LoggingLevel = "INFO"
	// NONE.
	LoggingLevel_NONE LoggingLevel = "NONE"
)

