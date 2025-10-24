package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for defining an S3 backup destination.
//
// S3 backup is available for all destinations, regardless of whether the final destination is S3 or not.
//
// Example:
//   // Enable backup of all source records (to an S3 bucket created by CDK).
//   var bucket Bucket
//   // Explicitly provide an S3 bucket to which all source records will be backed up.
//   var backupBucket Bucket
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Mode: firehose.BackupMode_ALL,
//   		},
//   	}),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Bucket"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Bucket: backupBucket,
//   		},
//   	}),
//   })
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Prefix"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Mode: firehose.BackupMode_ALL,
//   			DataOutputPrefix: jsii.String("mybackup"),
//   		},
//   	}),
//   })
//
type DestinationS3BackupProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(0)
	// Maximum: Duration.seconds(900)
	// Default: Duration.seconds(300)
	//
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Amazon Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1) when record data format conversion is disabled, Size.mebibytes(64) when it is enabled
	// Maximum: Size.mebibytes(128)
	// Default: Size.mebibytes(5) when record data format conversion is disabled, Size.mebibytes(128) when it is enabled
	//
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Amazon Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Default: - UNCOMPRESSED.
	//
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Amazon Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Default: "YYYY/MM/DD/HH".
	//
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Default: - Data is not encrypted.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Amazon Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Default: "YYYY/MM/DD/HH".
	//
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The S3 bucket that will store data and failed records.
	// Default: - If `mode` is set to `BackupMode.ALL` or `BackupMode.FAILED`, a bucket will be created for you.
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Configuration that determines whether to log errors during data transformation or delivery failures, and specifies the CloudWatch log group for storing error logs.
	// Default: - errors will be logged and a log group will be created for you.
	//
	LoggingConfig ILoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Indicates the mode by which incoming records should be backed up to S3, if any.
	//
	// If `bucket` is provided, this will be implicitly set to `BackupMode.ALL`.
	// Default: - If `bucket` is provided, the default will be `BackupMode.ALL`. Otherwise,
	// source records are not backed up to S3.
	//
	Mode BackupMode `field:"optional" json:"mode" yaml:"mode"`
}

