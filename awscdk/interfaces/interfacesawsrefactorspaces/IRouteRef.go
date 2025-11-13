package interfacesawsrefactorspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrefactorspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Route.
// Experimental.
type IRouteRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Route resource.
	// Experimental.
	RouteRef() *RouteReference
}

// The jsii proxy for IRouteRef
type jsiiProxy_IRouteRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRouteRef) RouteRef() *RouteReference {
	var returns *RouteReference
	_jsii_.Get(
		j,
		"routeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

