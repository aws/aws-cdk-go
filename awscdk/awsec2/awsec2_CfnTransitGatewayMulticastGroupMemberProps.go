package awsec2


// Properties for defining a `CfnTransitGatewayMulticastGroupMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMulticastGroupMemberProps := &cfnTransitGatewayMulticastGroupMemberProps{
//   	groupIpAddress: jsii.String("groupIpAddress"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	transitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type CfnTransitGatewayMulticastGroupMemberProps struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The group members' network interface IDs to register with the transit gateway multicast group.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

