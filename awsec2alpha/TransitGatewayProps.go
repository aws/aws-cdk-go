package awsec2alpha


// Common properties for creating a Transit Gateway resource.
//
// Example:
//   transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"), &TransitGatewayProps{
//   	DefaultRouteTableAssociation: jsii.Boolean(false),
//   	DefaultRouteTablePropagation: jsii.Boolean(false),
//   })
//
// Experimental.
type TransitGatewayProps struct {
	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	//
	// The range is 64512 to 65534 for 16-bit ASNs.
	// Default: - undefined, 64512 is assigned by CloudFormation.
	//
	// Experimental.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Enable or disable automatic acceptance of cross-account attachment requests.
	// Default: - disable (false).
	//
	// Experimental.
	AutoAcceptSharedAttachments *bool `field:"optional" json:"autoAcceptSharedAttachments" yaml:"autoAcceptSharedAttachments"`
	// Enable or disable automatic association with the default association route table.
	// Default: - enable (true).
	//
	// Experimental.
	DefaultRouteTableAssociation *bool `field:"optional" json:"defaultRouteTableAssociation" yaml:"defaultRouteTableAssociation"`
	// Enable or disable automatic propagation of routes to the default propagation route table.
	// Default: - enable (true).
	//
	// Experimental.
	DefaultRouteTablePropagation *bool `field:"optional" json:"defaultRouteTablePropagation" yaml:"defaultRouteTablePropagation"`
	// The description of the transit gateway.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable DNS support.
	//
	// If dnsSupport is enabled on a VPC Attachment, this also needs to be enabled for the feature to work.
	// Otherwise the resources will still deploy but the feature will not work.
	// Default: - enable (true).
	//
	// Experimental.
	DnsSupport *bool `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enable or disable security group referencing support.
	//
	// If securityGroupReferencingSupport is enabled on a VPC Attachment, this also needs to be enabled for the feature to work.
	// Otherwise the resources will still deploy but the feature will not work.
	// Default: - disable (false).
	//
	// Experimental.
	SecurityGroupReferencingSupport *bool `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
	// The transit gateway CIDR blocks.
	// Default: - none.
	//
	// Experimental.
	TransitGatewayCidrBlocks *[]*string `field:"optional" json:"transitGatewayCidrBlocks" yaml:"transitGatewayCidrBlocks"`
	// Physical name of this Transit Gateway.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayName *string `field:"optional" json:"transitGatewayName" yaml:"transitGatewayName"`
}

