package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Peer object factories (to be used in Security Group management).
//
// The static methods on this object can be used to create peer objects
// which represent a connection partner in Security Group rules.
//
// Use this object if you need to represent connection partners using plain IP
// addresses, or a prefix list ID.
//
// If you want to address a connection partner by Security Group, you can just
// use the Security Group (or the construct that contains a Security Group)
// directly, as it already implements `IPeer`.
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
type Peer interface {
}

// The jsii proxy struct for Peer
type jsiiProxy_Peer struct {
	_ byte // padding
}

func NewPeer() Peer {
	_init_.Initialize()

	j := jsiiProxy_Peer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Peer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewPeer_Override(p Peer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Peer",
		nil, // no parameters
		p,
	)
}

// Any IPv4 address.
func Peer_AnyIpv4() IPeer {
	_init_.Initialize()

	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"anyIpv4",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any IPv6 address.
func Peer_AnyIpv6() IPeer {
	_init_.Initialize()

	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"anyIpv6",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create an IPv4 peer from a CIDR.
func Peer_Ipv4(cidrIp *string) IPeer {
	_init_.Initialize()

	if err := validatePeer_Ipv4Parameters(cidrIp); err != nil {
		panic(err)
	}
	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"ipv4",
		[]interface{}{cidrIp},
		&returns,
	)

	return returns
}

// Create an IPv6 peer from a CIDR.
func Peer_Ipv6(cidrIp *string) IPeer {
	_init_.Initialize()

	if err := validatePeer_Ipv6Parameters(cidrIp); err != nil {
		panic(err)
	}
	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"ipv6",
		[]interface{}{cidrIp},
		&returns,
	)

	return returns
}

// A prefix list.
func Peer_PrefixList(prefixListId *string) IPeer {
	_init_.Initialize()

	if err := validatePeer_PrefixListParameters(prefixListId); err != nil {
		panic(err)
	}
	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"prefixList",
		[]interface{}{prefixListId},
		&returns,
	)

	return returns
}

// A security group ID.
func Peer_SecurityGroupId(securityGroupId *string, sourceSecurityGroupOwnerId *string) IPeer {
	_init_.Initialize()

	if err := validatePeer_SecurityGroupIdParameters(securityGroupId); err != nil {
		panic(err)
	}
	var returns IPeer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Peer",
		"securityGroupId",
		[]interface{}{securityGroupId, sourceSecurityGroupOwnerId},
		&returns,
	)

	return returns
}

