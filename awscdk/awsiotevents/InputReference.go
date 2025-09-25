package awsiotevents


// A reference to a Input resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputReference := &InputReference{
//   	InputName: jsii.String("inputName"),
//   }
//
type InputReference struct {
	// The InputName of the Input resource.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
}

