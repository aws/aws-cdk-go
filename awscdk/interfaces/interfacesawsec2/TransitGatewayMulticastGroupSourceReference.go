package interfacesawsec2


// A reference to a TransitGatewayMulticastGroupSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMulticastGroupSourceReference := &TransitGatewayMulticastGroupSourceReference{
//   	GroupIpAddress: jsii.String("groupIpAddress"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type TransitGatewayMulticastGroupSourceReference struct {
	// The GroupIpAddress of the TransitGatewayMulticastGroupSource resource.
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The NetworkInterfaceId of the TransitGatewayMulticastGroupSource resource.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The TransitGatewayMulticastDomainId of the TransitGatewayMulticastGroupSource resource.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

