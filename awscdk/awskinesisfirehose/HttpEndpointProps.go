package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for defining an Http destination of a Kinesis Data Firehose delivery stream.
//
// Example:
//   var endpointConfig HttpEndpointConfig
//
//   httpDestination := firehose.NewHttpEndpoint(&HttpEndpointProps{
//   	EndpointConfig: EndpointConfig,
//   })
//
type HttpEndpointProps struct {
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
	// Describes the configuration of the Http endpoint to which Kinesis Firehose delivers data.
	EndpointConfig *HttpEndpointConfig `field:"required" json:"endpointConfig" yaml:"endpointConfig"`
	// Describes the metadata sent to the Http endpoint destination.
	// Default: - None.
	//
	Attributes *[]*HttpAttribute `field:"optional" json:"attributes" yaml:"attributes"`
	// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to the Http endpoint destination.
	// Default: - Failed data only.
	//
	BackupMode HttpBackupMode `field:"optional" json:"backupMode" yaml:"backupMode"`
	// The buffering options that can be used before data is delivered to the specified destination.
	// Default: - None.
	//
	BufferingHints *HttpBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Compress the body of a request before sending the request to the destination.
	// Default: - None.
	//
	RequestCompression HttpCompression `field:"optional" json:"requestCompression" yaml:"requestCompression"`
	// The total amount of time that Kinesis Data Firehose spends on retries.
	// Default: - None.
	//
	RetryOptions *HttpRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

