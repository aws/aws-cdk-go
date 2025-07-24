package awsiotwireless


// A list of optional LoRaWAN application information, which can be used for geolocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationProperty := &ApplicationProperty{
//   	DestinationName: jsii.String("destinationName"),
//   	FPort: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html
//
type CfnWirelessDevice_ApplicationProperty struct {
	// The name of the position data destination that describes the IoT rule that processes the device's position data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-destinationname
	//
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
	// The name of the new destination for the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-fport
	//
	FPort *float64 `field:"optional" json:"fPort" yaml:"fPort"`
	// Application type, which can be specified to obtain real-time position information of your LoRaWAN device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-application.html#cfn-iotwireless-wirelessdevice-application-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

