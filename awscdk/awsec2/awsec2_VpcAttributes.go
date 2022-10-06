package awsec2


// Properties that reference an external Vpc.
//
// Example:
//   sg := ec2.securityGroup.fromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.lustreFileSystem.fromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &fileSystemAttributes{
//   	dnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	fileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	securityGroup: sg,
//   })
//
//   vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("Vpc"), &vpcAttributes{
//   	availabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	vpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &instanceProps{
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_T2, ec2.instanceSize_LARGE),
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   fs.connections.allowDefaultPortFrom(inst)
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
	// List of names for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// List of IDs of routing tables for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// List of private subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// List of names for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// List of IDs of routing tables for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// List of public subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// List of names for the public subnets.
	//
	// Must be undefined or have a name for every public subnet group.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// List of IDs of routing tables for the public subnets.
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

