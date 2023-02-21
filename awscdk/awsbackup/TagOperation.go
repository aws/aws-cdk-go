package awsbackup


// An operation that is applied to a key-value pair.
type TagOperation string

const (
	// StringEquals.
	TagOperation_STRING_EQUALS TagOperation = "STRING_EQUALS"
	// Dummy member.
	TagOperation_DUMMY TagOperation = "DUMMY"
)

