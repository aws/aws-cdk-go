package awsnetworkmanager


// A reference to a Device resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceReference := &DeviceReference{
//   	DeviceArn: jsii.String("deviceArn"),
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   }
//
type DeviceReference struct {
	// The ARN of the Device resource.
	DeviceArn *string `field:"required" json:"deviceArn" yaml:"deviceArn"`
	// The DeviceId of the Device resource.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// The GlobalNetworkId of the Device resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
}

