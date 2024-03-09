package awsec2


// Direction of traffic to allow all by default.
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
type NatTrafficDirection string

const (
	// Allow all outbound traffic and disallow all inbound traffic.
	NatTrafficDirection_OUTBOUND_ONLY NatTrafficDirection = "OUTBOUND_ONLY"
	// Allow all outbound and inbound traffic.
	NatTrafficDirection_INBOUND_AND_OUTBOUND NatTrafficDirection = "INBOUND_AND_OUTBOUND"
	// Disallow all outbound and inbound traffic.
	NatTrafficDirection_NONE NatTrafficDirection = "NONE"
)

