package awscdkcloudassemblyschema


// Query input for looking up a VPC.
type VpcContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Filters to apply to the VPC.
	//
	// Filter parameters are the same as passed to DescribeVpcs.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html
	//
	Filter *map[string]*string `field:"required" json:"filter" yaml:"filter"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// Whether to populate the subnetGroups field of the `VpcContextResponse`, which contains potentially asymmetric subnet groups.
	// Default: false.
	//
	ReturnAsymmetricSubnets *bool `field:"optional" json:"returnAsymmetricSubnets" yaml:"returnAsymmetricSubnets"`
	// Whether to populate the `vpnGatewayId` field of the `VpcContextResponse`, which contains the VPN Gateway ID, if one exists.
	//
	// You can explicitly
	// disable this in order to avoid the lookup if you know the VPC does not have
	// a VPN Gatway attached.
	// Default: true.
	//
	ReturnVpnGateways *bool `field:"optional" json:"returnVpnGateways" yaml:"returnVpnGateways"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	// Default: 'aws-cdk:subnet-name'.
	//
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
}

