package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for classes that provide the connection-specification parts of a security group rule.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V3_8_X(),
//   	Vpc: Vpc,
//   })
//
//   cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(2181)))
//   cluster.connections.AllowFrom(ec2.Peer_Ipv4(jsii.String("1.2.3.4/8")), ec2.Port_Tcp(jsii.Number(9094)))
//
type Port interface {
	// Whether the rule containing this port range can be inlined into a securitygroup or not.
	CanInlineRule() *bool
	// Produce the ingress/egress rule JSON for the given connection.
	ToRuleJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Port
type jsiiProxy_Port struct {
	_ byte // padding
}

func (j *jsiiProxy_Port) CanInlineRule() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canInlineRule",
		&returns,
	)
	return returns
}


func NewPort(props *PortProps) Port {
	_init_.Initialize()

	if err := validateNewPortParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Port{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Port",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPort_Override(p Port, props *PortProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Port",
		[]interface{}{props},
		p,
	)
}

// A single AH port.
func Port_Ah() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"ah",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All ICMP traffic.
func Port_AllIcmp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"allIcmp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All ICMPv6 traffic.
func Port_AllIcmpV6() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"allIcmpV6",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any TCP traffic.
func Port_AllTcp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"allTcp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All traffic.
func Port_AllTraffic() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"allTraffic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any UDP traffic.
func Port_AllUdp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"allUdp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A single ESP port.
func Port_Esp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"esp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ICMP ping (echo) traffic.
func Port_IcmpPing() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"icmpPing",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All codes for a single ICMP type.
func Port_IcmpType(type_ *float64) Port {
	_init_.Initialize()

	if err := validatePort_IcmpTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"icmpType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// A specific combination of ICMP type and code.
// See: https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml
//
func Port_IcmpTypeAndCode(type_ *float64, code *float64) Port {
	_init_.Initialize()

	if err := validatePort_IcmpTypeAndCodeParameters(type_, code); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"icmpTypeAndCode",
		[]interface{}{type_, code},
		&returns,
	)

	return returns
}

// A single TCP port.
func Port_Tcp(port *float64) Port {
	_init_.Initialize()

	if err := validatePort_TcpParameters(port); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A TCP port range.
func Port_TcpRange(startPort *float64, endPort *float64) Port {
	_init_.Initialize()

	if err := validatePort_TcpRangeParameters(startPort, endPort); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"tcpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// A single UDP port.
func Port_Udp(port *float64) Port {
	_init_.Initialize()

	if err := validatePort_UdpParameters(port); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"udp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A UDP port range.
func Port_UdpRange(startPort *float64, endPort *float64) Port {
	_init_.Initialize()

	if err := validatePort_UdpRangeParameters(startPort, endPort); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Port",
		"udpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

func Port_DNS_TCP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"DNS_TCP",
		&returns,
	)
	return returns
}

func Port_DNS_UDP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"DNS_UDP",
		&returns,
	)
	return returns
}

func Port_HTTP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"HTTP",
		&returns,
	)
	return returns
}

func Port_HTTPS() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"HTTPS",
		&returns,
	)
	return returns
}

func Port_IMAP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"IMAP",
		&returns,
	)
	return returns
}

func Port_IMAPS() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"IMAPS",
		&returns,
	)
	return returns
}

func Port_LDAP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"LDAP",
		&returns,
	)
	return returns
}

func Port_MSSQL() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"MSSQL",
		&returns,
	)
	return returns
}

func Port_MYSQL_AURORA() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"MYSQL_AURORA",
		&returns,
	)
	return returns
}

func Port_NFS() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"NFS",
		&returns,
	)
	return returns
}

func Port_POP3() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"POP3",
		&returns,
	)
	return returns
}

func Port_POP3S() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"POP3S",
		&returns,
	)
	return returns
}

func Port_POSTGRES() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"POSTGRES",
		&returns,
	)
	return returns
}

func Port_RDP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"RDP",
		&returns,
	)
	return returns
}

func Port_SMB() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"SMB",
		&returns,
	)
	return returns
}

func Port_SMTP() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"SMTP",
		&returns,
	)
	return returns
}

func Port_SSH() Port {
	_init_.Initialize()
	var returns Port
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.Port",
		"SSH",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Port) ToRuleJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toRuleJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Port) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

