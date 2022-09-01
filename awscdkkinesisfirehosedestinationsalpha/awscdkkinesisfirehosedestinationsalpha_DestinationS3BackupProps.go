// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for defining an S3 backup destination.
//
// S3 backup is available for all destinations, regardless of whether the final destination is S3 or not.
//
// Example:
//   // Enable backup of all source records (to an S3 bucket created by CDK).
//   var bucket bucket
//   // Explicitly provide an S3 bucket to which all source records will be backed up.
//   var backupBucket bucket
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.backupMode_ALL,
//   			},
//   		}),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Bucket"), &deliveryStreamProps{
//   	destinations: []*iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				bucket: backupBucket,
//   			},
//   		}),
//   	},
//   })
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Prefix"), &deliveryStreamProps{
//   	destinations: []*iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.*backupMode_ALL,
//   				dataOutputPrefix: jsii.String("mybackup"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type DestinationS3BackupProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The S3 bucket that will store data and failed records.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Indicates the mode by which incoming records should be backed up to S3, if any.
	//
	// If `bucket` is provided, this will be implicitly set to `BackupMode.ALL`.
	// Experimental.
	Mode BackupMode `field:"optional" json:"mode" yaml:"mode"`
}

