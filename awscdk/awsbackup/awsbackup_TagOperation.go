package awsbackup


// An operation that is applied to a key-value pair.
// Experimental.
type TagOperation string

const (
	// StringEquals.
	// Experimental.
	TagOperation_STRING_EQUALS TagOperation = "STRING_EQUALS"
	// Dummy member.
	// Experimental.
	TagOperation_DUMMY TagOperation = "DUMMY"
)

