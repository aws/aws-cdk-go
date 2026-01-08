package interfacesawssagemaker


// A reference to a Device resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceReference := &DeviceReference{
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//   }
//
type DeviceReference struct {
	// The DeviceFleetName of the Device resource.
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
}

