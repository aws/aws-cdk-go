package awsbedrock


// A reference to a FlowVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowVersionReference := &FlowVersionReference{
//   	FlowArn: jsii.String("flowArn"),
//   	Version: jsii.String("version"),
//   }
//
type FlowVersionReference struct {
	// The FlowArn of the FlowVersion resource.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The Version of the FlowVersion resource.
	Version *string `field:"required" json:"version" yaml:"version"`
}

