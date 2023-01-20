package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetCloudWatchLogsParametersProperty := &pipeTargetCloudWatchLogsParametersProperty{
//   	logStreamName: jsii.String("logStreamName"),
//   	timestamp: jsii.String("timestamp"),
//   }
//
type CfnPipe_PipeTargetCloudWatchLogsParametersProperty struct {
	// `CfnPipe.PipeTargetCloudWatchLogsParametersProperty.LogStreamName`.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// `CfnPipe.PipeTargetCloudWatchLogsParametersProperty.Timestamp`.
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

