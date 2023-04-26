package awsec2


// Properties that reference an external Vpc.
//
// Example:
//   sg := ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.LustreFileSystem_FromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &FileSystemAttributes{
//   	DnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	FileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	SecurityGroup: sg,
//   })
//
//   vpc := ec2.Vpc_FromVpcAttributes(this, jsii.String("Vpc"), &VpcAttributes{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	PublicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	VpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &InstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T2, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	Vpc: Vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
//   fs.Connections.AllowDefaultPortFrom(inst)
//
// Experimental.
type VpcAttributes struct {
	// List of availability zones for the subnets in this VPC.
	// Experimental.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// VPC's identifier.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// List of isolated subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	// Experimental.
	IsolatedSubnetIds *[]*string `field:"optional" json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// List of names for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	// Experimental.
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// List of IDs of routing tables for the isolated subnets.
	//
	// Must be undefined or have a name for every isolated subnet group.
	// Experimental.
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// List of private subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	// Experimental.
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// List of names for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	// Experimental.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// List of IDs of routing tables for the private subnets.
	//
	// Must be undefined or have a name for every private subnet group.
	// Experimental.
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// List of public subnet IDs.
	//
	// Must be undefined or match the availability zones in length and order.
	// Experimental.
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// List of names for the public subnets.
	//
	// Must be undefined or have a name for every public subnet group.
	// Experimental.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// List of IDs of routing tables for the public subnets.
	//
	// Must be undefined or have a name for every public subnet group.
	// Experimental.
	PublicSubnetRouteTableIds *[]*string `field:"optional" json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// VPC's CIDR range.
	// Experimental.
	VpcCidrBlock *string `field:"optional" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// VPN gateway's identifier.
	// Experimental.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

