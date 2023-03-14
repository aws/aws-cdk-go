package awsiotwireless


// LoRaWAN wireless gateway object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANGatewayProperty := &loRaWANGatewayProperty{
//   	gatewayEui: jsii.String("gatewayEui"),
//   	rfRegion: jsii.String("rfRegion"),
//   }
//
type CfnWirelessGateway_LoRaWANGatewayProperty struct {
	// The gateway's EUI value.
	GatewayEui *string `field:"required" json:"gatewayEui" yaml:"gatewayEui"`
	// The frequency band (RFRegion) value.
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

