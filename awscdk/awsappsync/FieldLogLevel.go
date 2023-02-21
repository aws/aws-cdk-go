package awsappsync


// log-level for fields in AppSync.
type FieldLogLevel string

const (
	// No logging.
	FieldLogLevel_NONE FieldLogLevel = "NONE"
	// Error logging.
	FieldLogLevel_ERROR FieldLogLevel = "ERROR"
	// All logging.
	FieldLogLevel_ALL FieldLogLevel = "ALL"
)

