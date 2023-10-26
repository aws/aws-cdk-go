package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetCloudWatchLogsParametersProperty := &PipeTargetCloudWatchLogsParametersProperty{
//   	LogStreamName: jsii.String("logStreamName"),
//   	Timestamp: jsii.String("timestamp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html
//
type CfnPipe_PipeTargetCloudWatchLogsParametersProperty struct {
	// The name of the log stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-logstreamname
	//
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// The time the event occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-timestamp
	//
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

