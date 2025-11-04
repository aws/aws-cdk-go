package awsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingControl.
// Experimental.
type IRoutingControlRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RoutingControl resource.
	// Experimental.
	RoutingControlRef() *RoutingControlReference
}

// The jsii proxy for IRoutingControlRef
type jsiiProxy_IRoutingControlRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRoutingControlRef) RoutingControlRef() *RoutingControlReference {
	var returns *RoutingControlReference
	_jsii_.Get(
		j,
		"routingControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingControlRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingControlRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

