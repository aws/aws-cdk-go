package awsec2


// Properties for defining a `CfnTransitGatewayMulticastGroupSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMulticastGroupSourceProps := &cfnTransitGatewayMulticastGroupSourceProps{
//   	groupIpAddress: jsii.String("groupIpAddress"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	transitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type CfnTransitGatewayMulticastGroupSourceProps struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The group sources' network interface IDs to register with the transit gateway multicast group.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

