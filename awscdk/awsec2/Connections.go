package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Manage the allowed network connections for constructs with Security Groups.
//
// Security Groups can be thought of as a firewall for network-connected
// devices. This class makes it easy to allow network connections to and
// from security groups, and between security groups individually. When
// establishing connectivity between security groups, it will automatically
// add rules in both security groups
//
// This object can manage one or more security groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var peer iPeer
//   var port port
//   var securityGroup securityGroup
//
//   connections := awscdk.Aws_ec2.NewConnections(&ConnectionsProps{
//   	DefaultPort: port,
//   	Peer: peer,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   })
//
type Connections interface {
	IConnectable
	// The network connections associated with this resource.
	Connections() Connections
	// The default port configured for this connection peer, if available.
	DefaultPort() Port
	SecurityGroups() *[]ISecurityGroup
	// Add a security group to the list of security groups managed by this object.
	AddSecurityGroup(securityGroups ...ISecurityGroup)
	// Allow connections from the peer on our default port.
	//
	// Even if the peer has a default port, we will always use our default port.
	AllowDefaultPortFrom(other IConnectable, description *string)
	// Allow default connections from all IPv4 ranges.
	AllowDefaultPortFromAnyIpv4(description *string)
	// Allow hosts inside the security group to connect to each other.
	AllowDefaultPortInternally(description *string)
	// Allow connections from the peer on our default port.
	//
	// Even if the peer has a default port, we will always use our default port.
	AllowDefaultPortTo(other IConnectable, description *string)
	// Allow connections from the peer on the given port.
	AllowFrom(other IConnectable, portRange Port, description *string)
	// Allow from any IPv4 ranges.
	AllowFromAnyIpv4(portRange Port, description *string)
	// Allow hosts inside the security group to connect to each other on the given port.
	AllowInternally(portRange Port, description *string)
	// Allow connections to the peer on the given port.
	AllowTo(other IConnectable, portRange Port, description *string)
	// Allow to all IPv4 ranges.
	AllowToAnyIpv4(portRange Port, description *string)
	// Allow connections to the security group on their default port.
	AllowToDefaultPort(other IConnectable, description *string)
}

// The jsii proxy struct for Connections
type jsiiProxy_Connections struct {
	jsiiProxy_IConnectable
}

func (j *jsiiProxy_Connections) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connections) DefaultPort() Port {
	var returns Port
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connections) SecurityGroups() *[]ISecurityGroup {
	var returns *[]ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}


func NewConnections(props *ConnectionsProps) Connections {
	_init_.Initialize()

	if err := validateNewConnectionsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connections{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Connections",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewConnections_Override(c Connections, props *ConnectionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Connections",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_Connections) AddSecurityGroup(securityGroups ...ISecurityGroup) {
	args := []interface{}{}
	for _, a := range securityGroups {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addSecurityGroup",
		args,
	)
}

func (c *jsiiProxy_Connections) AllowDefaultPortFrom(other IConnectable, description *string) {
	if err := c.validateAllowDefaultPortFromParameters(other); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowDefaultPortFrom",
		[]interface{}{other, description},
	)
}

func (c *jsiiProxy_Connections) AllowDefaultPortFromAnyIpv4(description *string) {
	_jsii_.InvokeVoid(
		c,
		"allowDefaultPortFromAnyIpv4",
		[]interface{}{description},
	)
}

func (c *jsiiProxy_Connections) AllowDefaultPortInternally(description *string) {
	_jsii_.InvokeVoid(
		c,
		"allowDefaultPortInternally",
		[]interface{}{description},
	)
}

func (c *jsiiProxy_Connections) AllowDefaultPortTo(other IConnectable, description *string) {
	if err := c.validateAllowDefaultPortToParameters(other); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowDefaultPortTo",
		[]interface{}{other, description},
	)
}

func (c *jsiiProxy_Connections) AllowFrom(other IConnectable, portRange Port, description *string) {
	if err := c.validateAllowFromParameters(other, portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowFrom",
		[]interface{}{other, portRange, description},
	)
}

func (c *jsiiProxy_Connections) AllowFromAnyIpv4(portRange Port, description *string) {
	if err := c.validateAllowFromAnyIpv4Parameters(portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowFromAnyIpv4",
		[]interface{}{portRange, description},
	)
}

func (c *jsiiProxy_Connections) AllowInternally(portRange Port, description *string) {
	if err := c.validateAllowInternallyParameters(portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowInternally",
		[]interface{}{portRange, description},
	)
}

func (c *jsiiProxy_Connections) AllowTo(other IConnectable, portRange Port, description *string) {
	if err := c.validateAllowToParameters(other, portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowTo",
		[]interface{}{other, portRange, description},
	)
}

func (c *jsiiProxy_Connections) AllowToAnyIpv4(portRange Port, description *string) {
	if err := c.validateAllowToAnyIpv4Parameters(portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowToAnyIpv4",
		[]interface{}{portRange, description},
	)
}

func (c *jsiiProxy_Connections) AllowToDefaultPort(other IConnectable, description *string) {
	if err := c.validateAllowToDefaultPortParameters(other); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"allowToDefaultPort",
		[]interface{}{other, description},
	)
}

