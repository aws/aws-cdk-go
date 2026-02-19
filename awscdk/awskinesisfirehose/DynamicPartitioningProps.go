package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for defining dynamic partitioning.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	DynamicPartitioning: &DynamicPartitioningProps{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	Processors: []IDataProcessor{
//   		firehose.RecordDeAggregationProcessor_Delimited(jsii.String("####")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning.html
//
type DynamicPartitioningProps struct {
	// Whether to enable the dynamic partitioning.
	//
	// You cannot enable dynamic partitioning for an existing Firehose stream that does not have dynamic partitioning already enabled.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning-enable.html
	//
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// The total amount of time that Data Firehose spends on retries.
	//
	// Minimum: Duration.seconds(0)
	// Maximum: Duration.seconds(7200)
	// Default: Duration.seconds(300)
	//
	RetryDuration awscdk.Duration `field:"optional" json:"retryDuration" yaml:"retryDuration"`
}

