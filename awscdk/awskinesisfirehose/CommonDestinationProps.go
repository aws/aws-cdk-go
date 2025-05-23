package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Generic properties for defining a delivery stream destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var compression compression
//   var dataProcessor iDataProcessor
//   var key key
//   var loggingConfig iLoggingConfig
//   var role role
//   var size size
//
//   commonDestinationProps := &CommonDestinationProps{
//   	LoggingConfig: loggingConfig,
//   	Processor: dataProcessor,
//   	Role: role,
//   	S3Backup: &DestinationS3BackupProps{
//   		Bucket: bucket,
//   		BufferingInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   		BufferingSize: size,
//   		Compression: compression,
//   		DataOutputPrefix: jsii.String("dataOutputPrefix"),
//   		EncryptionKey: key,
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		LoggingConfig: loggingConfig,
//   		Mode: awscdk.Aws_kinesisfirehose.BackupMode_ALL,
//   	},
//   }
//
type CommonDestinationProps struct {
	// Configuration that determines whether to log errors during data transformation or delivery failures, and specifies the CloudWatch log group for storing error logs.
	// Default: - errors will be logged and a log group will be created for you.
	//
	LoggingConfig ILoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Default: - no data transformation will occur.
	//
	Processor IDataProcessor `field:"optional" json:"processor" yaml:"processor"`
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
}

