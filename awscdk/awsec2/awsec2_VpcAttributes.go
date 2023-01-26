package awsec2


// Properties that reference an external Vpc.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("VPC"), &vpcAttributes{
//   	vpcId: jsii.String("vpc-1234"),
//   	availabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//
//   	// Either pass literals for all IDs
//   	publicSubnetIds: []*string{
//   		jsii.String("s-12345"),
//   		jsii.String("s-67890"),
//   	},
//
//   	// OR: import a list of known length
//   	privateSubnetIds: awscdk.Fn.importListValue(jsii.String("PrivateSubnetIds"), jsii.Number(2)),
//
//   	// OR: split an imported string to a list of known length
//   	isolatedSubnetIds: awscdk.Fn.split(jsii.String(","), ssm.stringParameter.valueForStringParameter(this, jsii.String("MyParameter")), jsii.Number(2)),
//   })
//
type VpcAttributes struct {
	// List of availability zones for the subnets in this VPC.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// VPC's identifier.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// List of isolated subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	IsolatedSubnetIds *[]*string `field:"optional" json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// List of IPv4 CIDR blocks for the isolated subnets.
	//
	// Must be undefined or have an entry for every isolated subnet group.
	IsolatedSubnetIpv4CidrBlocks *[]*string `field:"optional" json:"isolatedSubnetIpv4CidrBlocks" yaml:"isolatedSubnetIpv4CidrBlocks"`
	// List of names for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// List of IDs of route tables for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// List of private subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// List of IPv4 CIDR blocks for the private subnets.
	//
	// Must be undefined or have an entry for every private subnet group.
	PrivateSubnetIpv4CidrBlocks *[]*string `field:"optional" json:"privateSubnetIpv4CidrBlocks" yaml:"privateSubnetIpv4CidrBlocks"`
	// List of names for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// List of IDs of route tables for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// List of public subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// List of IPv4 CIDR blocks for the public subnets.
	//
	// Must be undefined or have an entry for every public subnet group.
	PublicSubnetIpv4CidrBlocks *[]*string `field:"optional" json:"publicSubnetIpv4CidrBlocks" yaml:"publicSubnetIpv4CidrBlocks"`
	// List of names for the public subnets.
	//
	// Must be undefined or have a name for every public subnet group.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// List of IDs of route tables for the public subnets.
	//
	// Must be undefined or have a name for every public subnet group.
	PublicSubnetRouteTableIds *[]*string `field:"optional" json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// The region the VPC is in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// VPC's CIDR range.
	VpcCidrBlock *string `field:"optional" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// VPN gateway's identifier.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

