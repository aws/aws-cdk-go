package awss3


// Options for creating a Transfer Acceleration URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transferAccelerationUrlOptions := &transferAccelerationUrlOptions{
//   	dualStack: jsii.Boolean(false),
//   }
//
type TransferAccelerationUrlOptions struct {
	// Dual-stack support to connect to the bucket over IPv6.
	DualStack *bool `field:"optional" json:"dualStack" yaml:"dualStack"`
}

