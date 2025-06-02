package awslogs


// Properties for defining a `CfnLogStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogStreamProps := &CfnLogStreamProps{
//   	LogGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	LogStreamName: jsii.String("logStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
//
type CfnLogStreamProps struct {
	// The name of the log group where the log stream is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log stream.
	//
	// The name must be unique within the log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-logstreamname
	//
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

