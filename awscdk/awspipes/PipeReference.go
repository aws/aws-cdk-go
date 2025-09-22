package awspipes


// A reference to a Pipe resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeReference := &PipeReference{
//   	PipeArn: jsii.String("pipeArn"),
//   	PipeName: jsii.String("pipeName"),
//   }
//
type PipeReference struct {
	// The ARN of the Pipe resource.
	PipeArn *string `field:"required" json:"pipeArn" yaml:"pipeArn"`
	// The Name of the Pipe resource.
	PipeName *string `field:"required" json:"pipeName" yaml:"pipeName"`
}

