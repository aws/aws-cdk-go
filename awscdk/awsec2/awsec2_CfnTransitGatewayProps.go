package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayProps := &cfnTransitGatewayProps{
//   	amazonSideAsn: jsii.Number(123),
//   	associationDefaultRouteTableId: jsii.String("associationDefaultRouteTableId"),
//   	autoAcceptSharedAttachments: jsii.String("autoAcceptSharedAttachments"),
//   	defaultRouteTableAssociation: jsii.String("defaultRouteTableAssociation"),
//   	defaultRouteTablePropagation: jsii.String("defaultRouteTablePropagation"),
//   	description: jsii.String("description"),
//   	dnsSupport: jsii.String("dnsSupport"),
//   	multicastSupport: jsii.String("multicastSupport"),
//   	propagationDefaultRouteTableId: jsii.String("propagationDefaultRouteTableId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transitGatewayCidrBlocks: []*string{
//   		jsii.String("transitGatewayCidrBlocks"),
//   	},
//   	vpnEcmpSupport: jsii.String("vpnEcmpSupport"),
//   }
//
type CfnTransitGatewayProps struct {
	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	//
	// The range is 64512 to 65534 for 16-bit ASNs. The default is 64512.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The ID of the default association route table.
	AssociationDefaultRouteTableId *string `field:"optional" json:"associationDefaultRouteTableId" yaml:"associationDefaultRouteTableId"`
	// Enable or disable automatic acceptance of attachment requests.
	//
	// Disabled by default.
	AutoAcceptSharedAttachments *string `field:"optional" json:"autoAcceptSharedAttachments" yaml:"autoAcceptSharedAttachments"`
	// Enable or disable automatic association with the default association route table.
	//
	// Enabled by default.
	DefaultRouteTableAssociation *string `field:"optional" json:"defaultRouteTableAssociation" yaml:"defaultRouteTableAssociation"`
	// Enable or disable automatic propagation of routes to the default propagation route table.
	//
	// Enabled by default.
	DefaultRouteTablePropagation *string `field:"optional" json:"defaultRouteTablePropagation" yaml:"defaultRouteTablePropagation"`
	// The description of the transit gateway.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable DNS support.
	//
	// Enabled by default.
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Indicates whether multicast is enabled on the transit gateway.
	MulticastSupport *string `field:"optional" json:"multicastSupport" yaml:"multicastSupport"`
	// The ID of the default propagation route table.
	PropagationDefaultRouteTableId *string `field:"optional" json:"propagationDefaultRouteTableId" yaml:"propagationDefaultRouteTableId"`
	// The tags for the transit gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The transit gateway CIDR blocks.
	TransitGatewayCidrBlocks *[]*string `field:"optional" json:"transitGatewayCidrBlocks" yaml:"transitGatewayCidrBlocks"`
	// Enable or disable Equal Cost Multipath Protocol support.
	//
	// Enabled by default.
	VpnEcmpSupport *string `field:"optional" json:"vpnEcmpSupport" yaml:"vpnEcmpSupport"`
}

