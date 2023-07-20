package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayProps := &CfnTransitGatewayProps{
//   	AmazonSideAsn: jsii.Number(123),
//   	AssociationDefaultRouteTableId: jsii.String("associationDefaultRouteTableId"),
//   	AutoAcceptSharedAttachments: jsii.String("autoAcceptSharedAttachments"),
//   	DefaultRouteTableAssociation: jsii.String("defaultRouteTableAssociation"),
//   	DefaultRouteTablePropagation: jsii.String("defaultRouteTablePropagation"),
//   	Description: jsii.String("description"),
//   	DnsSupport: jsii.String("dnsSupport"),
//   	MulticastSupport: jsii.String("multicastSupport"),
//   	PropagationDefaultRouteTableId: jsii.String("propagationDefaultRouteTableId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
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
type CfnTransitGatewayProps struct {
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
	// Enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-defaultroutetableassociation
	//
	DefaultRouteTableAssociation *string `field:"optional" json:"defaultRouteTableAssociation" yaml:"defaultRouteTableAssociation"`
	// Enable or disable automatic propagation of routes to the default propagation route table.
	//
	// Enabled by default.
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
	// Indicates whether multicast is enabled on the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-multicastsupport
	//
	MulticastSupport *string `field:"optional" json:"multicastSupport" yaml:"multicastSupport"`
	// The ID of the default propagation route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgateway.html#cfn-ec2-transitgateway-propagationdefaultroutetableid
	//
	PropagationDefaultRouteTableId *string `field:"optional" json:"propagationDefaultRouteTableId" yaml:"propagationDefaultRouteTableId"`
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

