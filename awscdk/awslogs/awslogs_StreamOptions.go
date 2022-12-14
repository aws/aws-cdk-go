package awslogs


// Properties for a new LogStream created from a LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamOptions := &streamOptions{
//   	logStreamName: jsii.String("logStreamName"),
//   }
//
type StreamOptions struct {
	// The name of the log stream to create.
	//
	// The name must be unique within the log group.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

