package awslogs


// A reference to a LogStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logStreamReference := &LogStreamReference{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamName: jsii.String("logStreamName"),
//   }
//
type LogStreamReference struct {
	// The LogGroupName of the LogStream resource.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The LogStreamName of the LogStream resource.
	LogStreamName *string `field:"required" json:"logStreamName" yaml:"logStreamName"`
}

