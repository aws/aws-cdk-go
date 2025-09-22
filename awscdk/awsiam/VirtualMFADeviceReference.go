package awsiam


// A reference to a VirtualMFADevice resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualMFADeviceReference := &VirtualMFADeviceReference{
//   	SerialNumber: jsii.String("serialNumber"),
//   }
//
type VirtualMFADeviceReference struct {
	// The SerialNumber of the VirtualMFADevice resource.
	SerialNumber *string `field:"required" json:"serialNumber" yaml:"serialNumber"`
}

