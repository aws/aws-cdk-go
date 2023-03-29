package awspipes


// The parameters for using an CloudWatch Logs log stream as a target.
//
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
type CfnPipe_PipeTargetCloudWatchLogsParametersProperty struct {
	// The name of the log stream.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// The time the event occurred, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

