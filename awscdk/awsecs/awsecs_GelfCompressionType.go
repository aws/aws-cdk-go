package awsecs


// The type of compression the GELF driver uses to compress each log message.
// Experimental.
type GelfCompressionType string

const (
	// Experimental.
	GelfCompressionType_GZIP GelfCompressionType = "GZIP"
	// Experimental.
	GelfCompressionType_ZLIB GelfCompressionType = "ZLIB"
	// Experimental.
	GelfCompressionType_NONE GelfCompressionType = "NONE"
)

