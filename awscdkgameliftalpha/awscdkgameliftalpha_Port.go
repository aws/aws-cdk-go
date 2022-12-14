// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for classes that provide the connection-specification parts of a security group rule.
//
// Example:
//   var build build
//
//
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &buildFleetProps{
//   	fleetName: jsii.String("test-fleet"),
//   	content: build,
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_C4, ec2.instanceSize_LARGE),
//   	runtimeConfiguration: &runtimeConfiguration{
//   		serverProcesses: []serverProcess{
//   			&serverProcess{
//   				launchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   			},
//   		},
//   	},
//   	ingressRules: []ingressRule{
//   		&ingressRule{
//   			source: gamelift.peer.anyIpv4(),
//   			port: gamelift.port.tcpRange(jsii.Number(100), jsii.Number(200)),
//   		},
//   	},
//   })
//   // Allowing a specific CIDR for port 1111 on UDP Protocol
//   fleet.addIngressRule(gamelift.peer.ipv4(jsii.String("1.2.3.4/32")), gamelift.port.udp(jsii.Number(1111)))
//
// Experimental.
type Port interface {
	// Produce the ingress rule JSON for the given connection.
	// Experimental.
	ToJson() interface{}
}

// The jsii proxy struct for Port
type jsiiProxy_Port struct {
	_ byte // padding
}

// Experimental.
func NewPort(props *PortProps) Port {
	_init_.Initialize()

	if err := validateNewPortParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Port{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Port",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewPort_Override(p Port, props *PortProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Port",
		[]interface{}{props},
		p,
	)
}

// Any TCP traffic.
// Experimental.
func Port_AllTcp() Port {
	_init_.Initialize()

	var returns Port

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Port",
		"allTcp",
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
		"@aws-cdk/aws-gamelift-alpha.Port",
		"allUdp",
		nil, // no parameters
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
		"@aws-cdk/aws-gamelift-alpha.Port",
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
		"@aws-cdk/aws-gamelift-alpha.Port",
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
		"@aws-cdk/aws-gamelift-alpha.Port",
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
		"@aws-cdk/aws-gamelift-alpha.Port",
		"udpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Port) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

