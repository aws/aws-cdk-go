package awsec2


// Properties for defining a `CfnTransitGatewayMulticastGroupMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMulticastGroupMemberProps := &CfnTransitGatewayMulticastGroupMemberProps{
//   	GroupIpAddress: jsii.String("groupIpAddress"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastgroupmember.html
//
type CfnTransitGatewayMulticastGroupMemberProps struct {
	// The IP address assigned to the transit gateway multicast group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastgroupmember.html#cfn-ec2-transitgatewaymulticastgroupmember-groupipaddress
	//
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The group members' network interface IDs to register with the transit gateway multicast group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastgroupmember.html#cfn-ec2-transitgatewaymulticastgroupmember-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastgroupmember.html#cfn-ec2-transitgatewaymulticastgroupmember-transitgatewaymulticastdomainid
	//
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

