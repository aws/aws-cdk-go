package awsecs


// The type of compression the GELF driver uses to compress each log message.
type GelfCompressionType string

const (
	GelfCompressionType_GZIP GelfCompressionType = "GZIP"
	GelfCompressionType_ZLIB GelfCompressionType = "ZLIB"
	GelfCompressionType_NONE GelfCompressionType = "NONE"
)

