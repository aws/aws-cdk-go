package awsiotwireless


// LoRaWAN wireless gateway object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANGatewayProperty := &LoRaWANGatewayProperty{
//   	GatewayEui: jsii.String("gatewayEui"),
//   	RfRegion: jsii.String("rfRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessgateway-lorawangateway.html
//
type CfnWirelessGateway_LoRaWANGatewayProperty struct {
	// The gateway's EUI value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessgateway-lorawangateway.html#cfn-iotwireless-wirelessgateway-lorawangateway-gatewayeui
	//
	GatewayEui *string `field:"required" json:"gatewayEui" yaml:"gatewayEui"`
	// The frequency band (RFRegion) value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessgateway-lorawangateway.html#cfn-iotwireless-wirelessgateway-lorawangateway-rfregion
	//
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

