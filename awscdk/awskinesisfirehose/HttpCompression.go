package awskinesisfirehose


// Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination.
type HttpCompression string

const (
	// GZIP.
	HttpCompression_GZIP HttpCompression = "GZIP"
	// NONE.
	HttpCompression_NONE HttpCompression = "NONE"
)

