package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes"
)

// Log destination configuration parameters.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   sourceFilter := pipes.NewFilter([]iFilterPattern{
//   	pipes.FilterPattern_FromObject(map[string]interface{}{
//   		"body": map[string][]*string{
//   			// only forward events with customerType B2B or B2C
//   			"customerType": []*string{
//   				jsii.String("B2B"),
//   				jsii.String("B2C"),
//   			},
//   		},
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSqsSource(sourceQueue),
//   	Target: NewSqsTarget(targetQueue),
//   	Filter: sourceFilter,
//   })
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

