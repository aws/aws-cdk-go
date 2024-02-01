package awspipes


// Represents the Amazon CloudWatch Logs logging configuration settings for the pipe.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-cloudwatchlogslogdestination.html
//
type CfnPipe_CloudwatchLogsLogDestinationProperty struct {
	// The AWS Resource Name (ARN) for the CloudWatch log group to which EventBridge sends the log records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-cloudwatchlogslogdestination.html#cfn-pipes-pipe-cloudwatchlogslogdestination-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}

