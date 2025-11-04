package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingProfile.
// Experimental.
type IRoutingProfileRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RoutingProfile resource.
	// Experimental.
	RoutingProfileRef() *RoutingProfileReference
}

// The jsii proxy for IRoutingProfileRef
type jsiiProxy_IRoutingProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRoutingProfileRef) RoutingProfileRef() *RoutingProfileReference {
	var returns *RoutingProfileReference
	_jsii_.Get(
		j,
		"routingProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingProfileRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

