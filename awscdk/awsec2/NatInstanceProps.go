package awsec2


// Properties for a NAT instance.
//
// Example:
//   // Configure the `natGatewayProvider` when defining a Vpc
//   natGatewayProvider := ec2.NatProvider_Instance(&NatInstanceProps{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   })
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &VpcProps{
//   	NatGatewayProvider: NatGatewayProvider,
//
//   	// The 'natGateways' parameter now controls the number of NAT instances
//   	NatGateways: jsii.Number(2),
//   })
//
type NatInstanceProps struct {
	// Instance type of the NAT instance.
	InstanceType InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
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
}

