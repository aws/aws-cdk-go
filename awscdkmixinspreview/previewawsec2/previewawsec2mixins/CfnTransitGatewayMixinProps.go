package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTransitGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayMixinProps := &CfnTransitGatewayMixinProps{
//   	AmazonSideAsn: jsii.Number(123),
//   	AssociationDefaultRouteTableId: jsii.String("associationDefaultRouteTableId"),
//   	AutoAcceptSharedAttachments: jsii.String("autoAcceptSharedAttachments"),
//   	DefaultRouteTableAssociation: jsii.String("defaultRouteTableAssociation"),
//   	DefaultRouteTablePropagation: jsii.String("defaultRouteTablePropagation"),
//   	Description: jsii.String("description"),
//   	DnsSupport: jsii.String("dnsSupport"),
//   	EncryptionSupport: jsii.String("encryptionSupport"),
//   	MulticastSupport: jsii.String("multicastSupport"),
//   	PropagationDefaultRouteTableId: jsii.String("propagationDefaultRouteTableId"),
//   	SecurityGroupReferencingSupport: jsii.String("securityGroupReferencingSupport"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayCidrBlocks: []*string{
//   		jsii.String("transitGatewayCidrBlocks"),
//   	},
//   	VpnEcmpSupport: jsii.String("vpnEcmpSupport"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html
//
type CfnTransitGatewayMixinProps struct {
	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	//
	// The range is 64512 to 65534 for 16-bit ASNs. The default is 64512.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-amazonsideasn
	//
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The ID of the default association route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-associationdefaultroutetableid
	//
	AssociationDefaultRouteTableId *string `field:"optional" json:"associationDefaultRouteTableId" yaml:"associationDefaultRouteTableId"`
	// Enable or disable automatic acceptance of attachment requests.
	//
	// Disabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-autoacceptsharedattachments
	//
	AutoAcceptSharedAttachments *string `field:"optional" json:"autoAcceptSharedAttachments" yaml:"autoAcceptSharedAttachments"`
	// Enable or disable automatic association with the default association route table.
	//
	// Enabled by default. If `DefaultRouteTableAssociation` is set to enable, AWS Transit Gateway will create the default transit gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetableassociation
	//
	DefaultRouteTableAssociation *string `field:"optional" json:"defaultRouteTableAssociation" yaml:"defaultRouteTableAssociation"`
	// Enable or disable automatic propagation of routes to the default propagation route table.
	//
	// Enabled by default. If `DefaultRouteTablePropagation` is set to enable, AWS Transit Gateway will create the default transit gateway route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetablepropagation
	//
	DefaultRouteTablePropagation *string `field:"optional" json:"defaultRouteTablePropagation" yaml:"defaultRouteTablePropagation"`
	// The description of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable DNS support.
	//
	// Enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-dnssupport
	//
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enable or disable encryption support.
	//
	// Disabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-encryptionsupport
	//
	EncryptionSupport *string `field:"optional" json:"encryptionSupport" yaml:"encryptionSupport"`
	// Indicates whether multicast is enabled on the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-multicastsupport
	//
	MulticastSupport *string `field:"optional" json:"multicastSupport" yaml:"multicastSupport"`
	// The ID of the default propagation route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-propagationdefaultroutetableid
	//
	PropagationDefaultRouteTableId *string `field:"optional" json:"propagationDefaultRouteTableId" yaml:"propagationDefaultRouteTableId"`
	// Enables you to reference a security group across VPCs attached to a transit gateway (TGW).
	//
	// Use this option to simplify security group management and control of instance-to-instance traffic across VPCs that are connected by transit gateway. You can also use this option to migrate from VPC peering (which was the only option that supported security group referencing) to transit gateways (which now also support security group referencing). This option is disabled by default and there are no additional costs to use this feature.
	//
	// For important information about this feature, see [Create a transit gateway](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-transit-gateways.html#create-tgw) in the *AWS Transit Gateway Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-securitygroupreferencingsupport
	//
	SecurityGroupReferencingSupport *string `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
	// The tags for the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The transit gateway CIDR blocks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-transitgatewaycidrblocks
	//
	TransitGatewayCidrBlocks *[]*string `field:"optional" json:"transitGatewayCidrBlocks" yaml:"transitGatewayCidrBlocks"`
	// Enable or disable Equal Cost Multipath Protocol support.
	//
	// Enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-vpnecmpsupport
	//
	VpnEcmpSupport *string `field:"optional" json:"vpnEcmpSupport" yaml:"vpnEcmpSupport"`
}

