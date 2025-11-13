package interfacesawssagemaker


// A reference to a Device resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceReference := &DeviceReference{
//   	DeviceName: jsii.String("deviceName"),
//   }
//
type DeviceReference struct {
	// The Device/DeviceName of the Device resource.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
}

