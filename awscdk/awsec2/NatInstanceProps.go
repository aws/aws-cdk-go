package awsec2


// Properties for a NAT instance.
//
// Example:
//   var instanceType instanceType
//
//
//   provider := ec2.NatProvider_InstanceV2(&NatInstanceProps{
//   	InstanceType: InstanceType,
//   	DefaultAllowedTraffic: ec2.NatTrafficDirection_OUTBOUND_ONLY,
//   })
//   ec2.NewVpc(this, jsii.String("TheVPC"), &VpcProps{
//   	NatGatewayProvider: provider,
//   })
//   provider.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(80)))
//
type NatInstanceProps struct {
	// Instance type of the NAT instance.
	InstanceType InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Specifying the CPU credit type for burstable EC2 instance types (T2, T3, T3a, etc).
	//
	// The unlimited CPU credit option is not supported for T3 instances with dedicated host (`host`) tenancy.
	// Default: - T2 instances are standard, while T3, T4g, and T3a instances are unlimited.
	//
	CreditSpecification CpuCredits `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Direction to allow all traffic through the NAT instance by default.
	//
	// By default, inbound and outbound traffic is allowed.
	//
	// If you set this to another value than INBOUND_AND_OUTBOUND, you must
	// configure the NAT instance's security groups in another way, either by
	// passing in a fully configured Security Group using the `securityGroup`
	// property, or by configuring it using the `.securityGroup` or
	// `.connections` members after passing the NAT Instance Provider to a Vpc.
	// Default: NatTrafficDirection.INBOUND_AND_OUTBOUND
	//
	DefaultAllowedTraffic NatTrafficDirection `field:"optional" json:"defaultAllowedTraffic" yaml:"defaultAllowedTraffic"`
	// Name of SSH keypair to grant access to instance.
	// Default: - No SSH access will be possible.
	//
	// Deprecated: - Use `keyPair` instead - https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2-readme.html#using-an-existing-ec2-key-pair
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The SSH keypair to grant access to the instance.
	// Default: - No SSH access will be possible.
	//
	KeyPair IKeyPair `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The machine image (AMI) to use.
	//
	// By default, will do an AMI lookup for the latest NAT instance image.
	//
	// If you have a specific AMI ID you want to use, pass a `GenericLinuxImage`. For example:
	//
	// ```ts
	// ec2.NatProvider.instance({
	//   instanceType: new ec2.InstanceType('t3.micro'),
	//   machineImage: new ec2.GenericLinuxImage({
	//     'us-east-2': 'ami-0f9c61b5a562a16af'
	//   })
	// })
	// ```.
	// Default: - Latest NAT instance image.
	//
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Security Group for NAT instances.
	// Default: - A new security group will be created.
	//
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Custom user data to run on the NAT instances.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html#create-nat-ami
	//
	// Default: UserData.forLinux().addCommands(...NatInstanceProviderV2.DEFAULT_USER_DATA_COMMANDS);  - Appropriate user data commands to initialize and configure the NAT instances
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

