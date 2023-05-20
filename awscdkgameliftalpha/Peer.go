package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Peer object factories.
//
// The static methods on this object can be used to create peer objects
// which represent a connection partner in inbound permission rules.
//
// Use this object if you need to represent connection partners using plain IP addresses.
//
// Example:
//   var build build
//
//
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &BuildFleetProps{
//   	FleetName: jsii.String("test-fleet"),
//   	Content: build,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   	RuntimeConfiguration: &RuntimeConfiguration{
//   		ServerProcesses: []serverProcess{
//   			&serverProcess{
//   				LaunchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   			},
//   		},
//   	},
//   	IngressRules: []ingressRule{
//   		&ingressRule{
//   			Source: gamelift.Peer_AnyIpv4(),
//   			Port: gamelift.Port_TcpRange(jsii.Number(100), jsii.Number(200)),
//   		},
//   	},
//   })
//   // Allowing a specific CIDR for port 1111 on UDP Protocol
//   fleet.AddIngressRule(gamelift.Peer_Ipv4(jsii.String("1.2.3.4/32")), gamelift.Port_Udp(jsii.Number(1111)))
//
// Experimental.
type Peer interface {
}

// The jsii proxy struct for Peer
type jsiiProxy_Peer struct {
	_ byte // padding
}

// Experimental.
func NewPeer() Peer {
	_init_.Initialize()

	j := jsiiProxy_Peer{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Peer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewPeer_Override(p Peer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Peer",
		nil, // no parameters
		p,
	)
}

// Any IPv4 address.
// Experimental.
func Peer_AnyIpv4() IPeer {
	_init_.Initialize()

	var returns IPeer

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Peer",
		"anyIpv4",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create an IPv4 peer from a CIDR.
// Experimental.
func Peer_Ipv4(cidrIp *string) IPeer {
	_init_.Initialize()

	if err := validatePeer_Ipv4Parameters(cidrIp); err != nil {
		panic(err)
	}
	var returns IPeer

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Peer",
		"ipv4",
		[]interface{}{cidrIp},
		&returns,
	)

	return returns
}

