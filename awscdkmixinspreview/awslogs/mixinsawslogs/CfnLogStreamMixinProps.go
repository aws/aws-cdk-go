package mixinsawslogs


// Properties for CfnLogStreamPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogStreamMixinProps := &CfnLogStreamMixinProps{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamName: jsii.String("logStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
//
type CfnLogStreamMixinProps struct {
	// The name of the log group where the log stream is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log stream.
	//
	// The name must be unique within the log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-logstreamname
	//
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

