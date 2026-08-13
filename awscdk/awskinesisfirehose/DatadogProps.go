package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Props for defining a Datadog destination of a Kinesis Data Firehose delivery stream.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiKey Secret
//
//   datadogDestination := firehose.NewDatadog(&DatadogProps{
//   	ApiKey: ApiKey,
//   	Endpoint: firehose.DatadogEndpoint_LOGS_US1(),
//   })
//
type DatadogProps struct {
	// Configuration that determines whether to log errors during data transformation or delivery failures, and specifies the CloudWatch log group for storing error logs.
	// Default: - errors will be logged and a log group will be created for you.
	//
	LoggingConfig ILoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Default: - no data transformation will occur.
	//
	// Deprecated: Use `processors` instead.
	Processor IDataProcessor `field:"optional" json:"processor" yaml:"processor"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Default: - no data transformation will occur.
	//
	Processors *[]IDataProcessor `field:"optional" json:"processors" yaml:"processors"`
	// The IAM role associated with this destination.
	//
	// Assumed by Amazon Data Firehose to invoke processors and write to destinations.
	// Default: - a role will be created with default permissions.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The configuration for backing up source records to S3.
	// Default: - source records will not be backed up to S3.
	//
	S3Backup *DestinationS3BackupProps `field:"optional" json:"s3Backup" yaml:"s3Backup"`
	// The API key used to authenticate with Datadog.
	//
	// Delivered to Firehose through AWS Secrets Manager (Firehose retrieves it at runtime rather
	// than embedding it in the template).
	ApiKey awssecretsmanager.ISecret `field:"required" json:"apiKey" yaml:"apiKey"`
	// The Datadog endpoint to send data to.
	Endpoint DatadogEndpoint `field:"required" json:"endpoint" yaml:"endpoint"`
	// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to Datadog.
	// Default: HttpBackupMode.FAILED
	//
	BackupMode HttpBackupMode `field:"optional" json:"backupMode" yaml:"backupMode"`
	// Buffering hints for data delivery to the Datadog endpoint.
	// Default: - interval of 60 seconds, size of 4 MiB.
	//
	BufferingHints *HttpBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Content encoding applied to the request body before delivery.
	// Default: HttpCompression.GZIP
	//
	RequestCompression HttpCompression `field:"optional" json:"requestCompression" yaml:"requestCompression"`
	// Retry behavior when Kinesis Data Firehose cannot deliver data to Datadog.
	// Default: - duration of 60 seconds.
	//
	RetryOptions *HttpRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Datadog tags to apply for filtering.
	// Default: - No tags.
	//
	Tags *[]*HttpAttribute `field:"optional" json:"tags" yaml:"tags"`
}

