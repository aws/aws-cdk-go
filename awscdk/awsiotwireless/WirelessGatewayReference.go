package awsiotwireless


// A reference to a WirelessGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wirelessGatewayReference := &WirelessGatewayReference{
//   	WirelessGatewayArn: jsii.String("wirelessGatewayArn"),
//   	WirelessGatewayId: jsii.String("wirelessGatewayId"),
//   }
//
type WirelessGatewayReference struct {
	// The ARN of the WirelessGateway resource.
	WirelessGatewayArn *string `field:"required" json:"wirelessGatewayArn" yaml:"wirelessGatewayArn"`
	// The Id of the WirelessGateway resource.
	WirelessGatewayId *string `field:"required" json:"wirelessGatewayId" yaml:"wirelessGatewayId"`
}

