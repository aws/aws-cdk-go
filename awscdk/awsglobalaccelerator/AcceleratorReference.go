package awsglobalaccelerator


// A reference to a Accelerator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorReference := &AcceleratorReference{
//   	AcceleratorArn: jsii.String("acceleratorArn"),
//   }
//
type AcceleratorReference struct {
	// The AcceleratorArn of the Accelerator resource.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
}

