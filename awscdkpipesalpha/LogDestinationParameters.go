package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes"
)

// Log destination configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   logDestinationParameters := &LogDestinationParameters{
//   	CloudwatchLogsLogDestination: &CloudwatchLogsLogDestinationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	FirehoseLogDestination: &FirehoseLogDestinationProperty{
//   		DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	},
//   	S3LogDestination: &S3LogDestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		OutputFormat: jsii.String("outputFormat"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// Experimental.
type LogDestinationParameters struct {
	// The logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-cloudwatchlogslogdestination
	//
	// Default: - none.
	//
	// Experimental.
	CloudwatchLogsLogDestination *awspipes.CfnPipe_CloudwatchLogsLogDestinationProperty `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// The Amazon Kinesis Data Firehose logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-firehoselogdestination
	//
	// Default: - none.
	//
	// Experimental.
	FirehoseLogDestination *awspipes.CfnPipe_FirehoseLogDestinationProperty `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// The Amazon S3 logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-s3logdestination
	//
	// Default: - none.
	//
	// Experimental.
	S3LogDestination *awspipes.CfnPipe_S3LogDestinationProperty `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

