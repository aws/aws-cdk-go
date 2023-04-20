package awsappsync


// log-level for fields in AppSync.
// Experimental.
type FieldLogLevel string

const (
	// No logging.
	// Experimental.
	FieldLogLevel_NONE FieldLogLevel = "NONE"
	// Error logging.
	// Experimental.
	FieldLogLevel_ERROR FieldLogLevel = "ERROR"
	// All logging.
	// Experimental.
	FieldLogLevel_ALL FieldLogLevel = "ALL"
)

