package awslogs


// Attributes for importing a LogStream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logStreamAttributes := &LogStreamAttributes{
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamName: jsii.String("logStreamName"),
//   }
//
type LogStreamAttributes struct {
	// The name of the log group.
	// Default: - When not provided, logStreamRef will throw an error.
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The name of the log stream.
	LogStreamName *string `field:"required" json:"logStreamName" yaml:"logStreamName"`
}

