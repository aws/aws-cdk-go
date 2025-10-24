package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// CloudWatch Logs target properties.
//
// Example:
//   var sourceQueue Queue
//   var targetLogGroup LogGroup
//
//
//   logGroupTarget := targets.NewCloudWatchLogsTarget(targetLogGroup, &CloudWatchLogsTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: logGroupTarget,
//   })
//
// Experimental.
type CloudWatchLogsTargetParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// The name of the log stream.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-logstreamname
	//
	// Default: - none.
	//
	// Experimental.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// The JSON path expression that references the timestamp in the payload.
	//
	// This is the time that the event occurred, as a JSON path expression in the payload.
	//
	// Example:
	//   "$.data.timestamp"
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetcloudwatchlogsparameters.html#cfn-pipes-pipe-pipetargetcloudwatchlogsparameters-timestamp
	//
	// Default: - current time.
	//
	// Experimental.
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}

