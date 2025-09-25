package awsec2


// A reference to a TransitGatewayMulticastGroupMember resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMulticastGroupMemberReference := &TransitGatewayMulticastGroupMemberReference{
//   	GroupIpAddress: jsii.String("groupIpAddress"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type TransitGatewayMulticastGroupMemberReference struct {
	// The GroupIpAddress of the TransitGatewayMulticastGroupMember resource.
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The NetworkInterfaceId of the TransitGatewayMulticastGroupMember resource.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The TransitGatewayMulticastDomainId of the TransitGatewayMulticastGroupMember resource.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

