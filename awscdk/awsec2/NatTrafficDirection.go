package awsec2


// Direction of traffic to allow all by default.
//
// Example:
//   natGatewayProvider := ec2.NatProvider_InstanceV2(&NatInstanceProps{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   	DefaultAllowedTraffic: ec2.NatTrafficDirection_NONE,
//   })
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	NatGatewayProvider: NatGatewayProvider,
//   })
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   	AllowAllOutbound: jsii.Boolean(false),
//   })
//   securityGroup.AddEgressRule(ec2.Peer_AnyIpv4(), ec2.Port_Tcp(jsii.Number(443)))
//   for _, gatewayInstance := range natGatewayProvider.gatewayInstances {
//   	gatewayInstance.AddSecurityGroup(securityGroup)
//   }
//
type NatTrafficDirection string

const (
	// Allow all outbound traffic and disallow all inbound traffic.
	NatTrafficDirection_OUTBOUND_ONLY NatTrafficDirection = "OUTBOUND_ONLY"
	// Allow all outbound and inbound traffic.
	NatTrafficDirection_INBOUND_AND_OUTBOUND NatTrafficDirection = "INBOUND_AND_OUTBOUND"
	// Disallow all outbound and inbound traffic.
	NatTrafficDirection_NONE NatTrafficDirection = "NONE"
)

