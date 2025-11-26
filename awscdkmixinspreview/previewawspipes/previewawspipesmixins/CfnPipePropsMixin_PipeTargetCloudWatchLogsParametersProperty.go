package previewawspipesmixins


// The parameters for using an CloudWatch Logs log stream as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeTargetCloudWatchLogsParametersProperty := &PipeTargetCloudWatchLogsParametersProperty{
//   	LogStreamName: jsii.String("logStreamName"),
//   	Timestamp: jsii.String("timestamp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html
//
type CfnPipePropsMixin_PipeTargetCloudWatchLogsParametersProperty struct {
	// The name of the log stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-logstreamname
	//
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// A [dynamic path parameter](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) to a field in the payload containing the time the event occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
	//
	// The value cannot be a static timestamp as the provided timestamp would be applied to all events delivered by the Pipe, regardless of when they are actually delivered.
	//
	// If no dynamic path parameter is provided, the default value is the time the invocation is processed by the Pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-timestamp
	//
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

