package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for classes that provide the connection-specification parts of a security group rule.
//
// Example:
//   var loadBalancer applicationLoadBalancer
//
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVPC"))
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	vpc: vpc,
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   	}),
//   })
//
//   project.connections.allowTo(loadBalancer, ec2.port.tcp(jsii.Number(443)))
//
// Experimental.
type Port interface {
	// Whether the rule containing this port range can be inlined into a securitygroup or not.
	// Experimental.
	CanInlineRule() *bool
	// Produce the ingress/egress rule JSON for the given connection.
	// Experimental.
	ToRuleJson() interface{}
	// Experimental.
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


// Experimental.
func NewPort(props *PortProps) Port {
	_init_.Initialize()

	if err := validateNewPortParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Port{}

	_jsii_.Create(
		"monocdk.aws_ec2.Port",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewPort_Override(p Port, props *PortProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.Port",
		[]interface{}{props},
		p,
	)
}

// A single AH port.
// Experimental.
func Port_Ah() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"ah",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All ICMP traffic.
// Experimental.
func Port_AllIcmp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"allIcmp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any TCP traffic.
// Experimental.
func Port_AllTcp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"allTcp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All traffic.
// Experimental.
func Port_AllTraffic() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"allTraffic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any UDP traffic.
// Experimental.
func Port_AllUdp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"allUdp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A single ESP port.
// Experimental.
func Port_Esp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"esp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ICMP ping (echo) traffic.
// Experimental.
func Port_IcmpPing() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"icmpPing",
		nil, // no parameters
		&returns,
	)

	return returns
}

// All codes for a single ICMP type.
// Experimental.
func Port_IcmpType(type_ *float64) Port {
	_init_.Initialize()

	if err := validatePort_IcmpTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"icmpType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// A specific combination of ICMP type and code.
// See: https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml
//
// Experimental.
func Port_IcmpTypeAndCode(type_ *float64, code *float64) Port {
	_init_.Initialize()

	if err := validatePort_IcmpTypeAndCodeParameters(type_, code); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"icmpTypeAndCode",
		[]interface{}{type_, code},
		&returns,
	)

	return returns
}

// A single TCP port.
// Experimental.
func Port_Tcp(port *float64) Port {
	_init_.Initialize()

	if err := validatePort_TcpParameters(port); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A TCP port range.
// Experimental.
func Port_TcpRange(startPort *float64, endPort *float64) Port {
	_init_.Initialize()

	if err := validatePort_TcpRangeParameters(startPort, endPort); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"tcpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// A single UDP port.
// Experimental.
func Port_Udp(port *float64) Port {
	_init_.Initialize()

	if err := validatePort_UdpParameters(port); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"udp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A UDP port range.
// Experimental.
func Port_UdpRange(startPort *float64, endPort *float64) Port {
	_init_.Initialize()

	if err := validatePort_UdpRangeParameters(startPort, endPort); err != nil {
		panic(err)
	}
	var returns Port

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.Port",
		"udpRange",
		[]interface{}{startPort, endPort},
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

