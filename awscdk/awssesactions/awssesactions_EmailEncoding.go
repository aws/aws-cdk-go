package awssesactions


// The type of email encoding to use for a SNS action.
// Experimental.
type EmailEncoding string

const (
	// Base 64.
	// Experimental.
	EmailEncoding_BASE64 EmailEncoding = "BASE64"
	// UTF-8.
	// Experimental.
	EmailEncoding_UTF8 EmailEncoding = "UTF8"
)

