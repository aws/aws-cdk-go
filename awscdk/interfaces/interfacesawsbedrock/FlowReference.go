package interfacesawsbedrock


// A reference to a Flow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowReference := &FlowReference{
//   	FlowArn: jsii.String("flowArn"),
//   }
//
type FlowReference struct {
	// The Arn of the Flow resource.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
}

