package interfacesawsiotwireless


// A reference to a DeviceProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProfileReference := &DeviceProfileReference{
//   	DeviceProfileArn: jsii.String("deviceProfileArn"),
//   	DeviceProfileId: jsii.String("deviceProfileId"),
//   }
//
type DeviceProfileReference struct {
	// The ARN of the DeviceProfile resource.
	DeviceProfileArn *string `field:"required" json:"deviceProfileArn" yaml:"deviceProfileArn"`
	// The Id of the DeviceProfile resource.
	DeviceProfileId *string `field:"required" json:"deviceProfileId" yaml:"deviceProfileId"`
}

