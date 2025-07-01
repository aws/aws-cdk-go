package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing/internal"
)

// Reference to a listener's port just created.
//
// This implements IConnectable with a default port (the port that an ELB
// listener was just created on) for a given security group so that it can be
// conveniently used just like any Connectable. E.g:
//
//    const listener = elb.addListener(...);
//
//    listener.connections.allowDefaultPortFromAnyIPv4();
//    // or
//    instance.connections.allowToDefaultPort(listener);
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var port port
//   var securityGroup securityGroup
//
//   listenerPort := awscdk.Aws_elasticloadbalancing.NewListenerPort(securityGroup, port)
//
type ListenerPort interface {
	awsec2.IConnectable
	// The network connections associated with this resource.
	Connections() awsec2.Connections
}

// The jsii proxy struct for ListenerPort
type jsiiProxy_ListenerPort struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_ListenerPort) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


func NewListenerPort(securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) ListenerPort {
	_init_.Initialize()

	if err := validateNewListenerPortParameters(securityGroup, defaultPort); err != nil {
		panic(err)
	}
	j := jsiiProxy_ListenerPort{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		&j,
	)

	return &j
}

func NewListenerPort_Override(l ListenerPort, securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		l,
	)
}

