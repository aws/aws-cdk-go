package awsivschat


// The DestinationConfiguration property type describes a location where chat logs will be stored.
//
// Each member represents the configuration of one log destination. For logging, you define only one type of destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigurationProperty := &destinationConfigurationProperty{
//   	cloudWatchLogs: &cloudWatchLogsDestinationConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	firehose: &firehoseDestinationConfigurationProperty{
//   		deliveryStreamName: jsii.String("deliveryStreamName"),
//   	},
//   	s3: &s3DestinationConfigurationProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   }
//
type CfnLoggingConfiguration_DestinationConfigurationProperty struct {
	// An Amazon CloudWatch Logs destination configuration where chat activity will be logged.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// An Amazon Kinesis Data Firehose destination configuration where chat activity will be logged.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// An Amazon S3 destination configuration where chat activity will be logged.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

