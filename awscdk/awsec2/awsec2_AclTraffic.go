package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The traffic that is configured using a Network ACL entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclTraffic := awscdk.Aws_ec2.aclTraffic.allTraffic()
//
// Experimental.
type AclTraffic interface {
	// Experimental.
	ToTrafficConfig() *AclTrafficConfig
}

// The jsii proxy struct for AclTraffic
type jsiiProxy_AclTraffic struct {
	_ byte // padding
}

// Experimental.
func NewAclTraffic_Override(a AclTraffic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.AclTraffic",
		nil, // no parameters
		a,
	)
}

// Apply the ACL entry to all traffic.
// Experimental.
func AclTraffic_AllTraffic() AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"allTraffic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Apply the ACL entry to ICMP traffic of given type and code.
// Experimental.
func AclTraffic_Icmp(props *AclIcmp) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"icmp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Apply the ACL entry to ICMPv6 traffic of given type and code.
//
// Requires an IPv6 CIDR block.
// Experimental.
func AclTraffic_Icmpv6(props *AclIcmp) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"icmpv6",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Apply the ACL entry to TCP traffic on a given port.
// Experimental.
func AclTraffic_TcpPort(port *float64) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"tcpPort",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Apply the ACL entry to TCP traffic on a given port range.
// Experimental.
func AclTraffic_TcpPortRange(startPort *float64, endPort *float64) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"tcpPortRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// Apply the ACL entry to UDP traffic on a given port.
// Experimental.
func AclTraffic_UdpPort(port *float64) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"udpPort",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// Apply the ACL entry to UDP traffic on a given port range.
// Experimental.
func AclTraffic_UdpPortRange(startPort *float64, endPort *float64) AclTraffic {
	_init_.Initialize()

	var returns AclTraffic

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.AclTraffic",
		"udpPortRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclTraffic) ToTrafficConfig() *AclTrafficConfig {
	var returns *AclTrafficConfig

	_jsii_.Invoke(
		a,
		"toTrafficConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

