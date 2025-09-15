package awssagemaker


// A reference to a DeviceFleet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceFleetReference := &DeviceFleetReference{
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//   }
//
type DeviceFleetReference struct {
	// The DeviceFleetName of the DeviceFleet resource.
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
}

