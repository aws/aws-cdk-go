package awsappflow


// A reference to a Flow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowReference := &FlowReference{
//   	FlowArn: jsii.String("flowArn"),
//   	FlowName: jsii.String("flowName"),
//   }
//
type FlowReference struct {
	// The ARN of the Flow resource.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The FlowName of the Flow resource.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
}

