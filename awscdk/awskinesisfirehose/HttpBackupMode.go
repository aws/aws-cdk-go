package awskinesisfirehose


// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to the Http endpoint destination.
type HttpBackupMode string

const (
	// Back up only the documents that Kinesis Data Firehose could not deliver.
	HttpBackupMode_FAILED HttpBackupMode = "FAILED"
	// Back up all documents.
	HttpBackupMode_ALL HttpBackupMode = "ALL"
)

