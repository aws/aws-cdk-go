package interfacesawsiotwireless


// A reference to a WirelessDevice resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wirelessDeviceReference := &WirelessDeviceReference{
//   	WirelessDeviceArn: jsii.String("wirelessDeviceArn"),
//   	WirelessDeviceId: jsii.String("wirelessDeviceId"),
//   }
//
type WirelessDeviceReference struct {
	// The ARN of the WirelessDevice resource.
	WirelessDeviceArn *string `field:"required" json:"wirelessDeviceArn" yaml:"wirelessDeviceArn"`
	// The Id of the WirelessDevice resource.
	WirelessDeviceId *string `field:"required" json:"wirelessDeviceId" yaml:"wirelessDeviceId"`
}

