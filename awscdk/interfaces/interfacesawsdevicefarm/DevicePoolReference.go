package interfacesawsdevicefarm


// A reference to a DevicePool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devicePoolReference := &DevicePoolReference{
//   	DevicePoolArn: jsii.String("devicePoolArn"),
//   }
//
type DevicePoolReference struct {
	// The Arn of the DevicePool resource.
	DevicePoolArn *string `field:"required" json:"devicePoolArn" yaml:"devicePoolArn"`
}

