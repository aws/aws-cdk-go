package awslogs


// Properties for defining a `CfnLogStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogStreamProps := &cfnLogStreamProps{
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   }
//
type CfnLogStreamProps struct {
	// The name of the log group where the log stream is created.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log stream.
	//
	// The name must be unique within the log group.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

