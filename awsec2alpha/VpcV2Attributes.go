package awsec2alpha


// Options to import a VPC created outside of CDK stack.
//
// Example:
//   stack := awscdk.Newstack()
//
//   //Importing a cross acount or cross region VPC
//   importedVpc := awsec2alpha.VpcV2_FromVpcV2Attributes(stack, jsii.String("ImportedVpc"), &VpcV2Attributes{
//   	VpcId: jsii.String("mockVpcID"),
//   	VpcCidrBlock: jsii.String("10.0.0.0/16"),
//   	OwnerAccountId: jsii.String("123456789012"),
//   	Region: jsii.String("us-west-2"),
//   })
//
// Experimental.
type VpcV2Attributes struct {
	// Primary VPC CIDR Block of the imported VPC Can only be IPv4.
	// Experimental.
	VpcCidrBlock *string `field:"required" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// The VPC ID Refers to physical Id of the resource.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the AWS account that owns the imported VPC required in case of cross account VPC as given value will be used to set field account for imported VPC, which then later can be used for establishing VPC peering connection.
	// Default: - constructed with stack account value.
	//
	// Experimental.
	OwnerAccountId *string `field:"optional" json:"ownerAccountId" yaml:"ownerAccountId"`
	// Region in which imported VPC is hosted required in case of cross region VPC as given value will be used to set field region for imported VPC, which then later can be used for establishing VPC peering connection.
	// Default: - constructed with stack region value.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Import Secondary CIDR blocks associated with VPC.
	// Default: - No secondary IP address.
	//
	// Experimental.
	SecondaryCidrBlocks *[]*VPCCidrBlockattributes `field:"optional" json:"secondaryCidrBlocks" yaml:"secondaryCidrBlocks"`
	// Subnets associated with imported VPC.
	// Default: - no subnets provided to be imported.
	//
	// Experimental.
	Subnets *[]*SubnetV2Attributes `field:"optional" json:"subnets" yaml:"subnets"`
	// A VPN Gateway is attached to the VPC.
	// Default: - No VPN Gateway.
	//
	// Experimental.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

