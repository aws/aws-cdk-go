package awsec2


// Properties for a NAT instance.
//
// Example:
//   // Configure the `natGatewayProvider` when defining a Vpc
//   natGatewayProvider := ec2.natProvider.instance(&natInstanceProps{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   })
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
//   	natGatewayProvider: natGatewayProvider,
//
//   	// The 'natGateways' parameter now controls the number of NAT instances
//   	natGateways: jsii.Number(2),
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
	DefaultAllowedTraffic NatTrafficDirection `field:"optional" json:"defaultAllowedTraffic" yaml:"defaultAllowedTraffic"`
	// Name of SSH keypair to grant access to instance.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The machine image (AMI) to use.
	//
	// By default, will do an AMI lookup for the latest NAT instance image.
	//
	// If you have a specific AMI ID you want to use, pass a `GenericLinuxImage`. For example:
	//
	// ```ts
	// ec2.NatProvider.instance({
	//    instanceType: new ec2.InstanceType('t3.micro'),
	//    machineImage: new ec2.GenericLinuxImage({
	//      'us-east-2': 'ami-0f9c61b5a562a16af'
	//    })
	// })
	// ```.
	MachineImage IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Security Group for NAT instances.
	SecurityGroup ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

