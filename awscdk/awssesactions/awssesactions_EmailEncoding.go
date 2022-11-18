package awssesactions


// The type of email encoding to use for a SNS action.
type EmailEncoding string

const (
	// Base 64.
	EmailEncoding_BASE64 EmailEncoding = "BASE64"
	// UTF-8.
	EmailEncoding_UTF8 EmailEncoding = "UTF8"
)

