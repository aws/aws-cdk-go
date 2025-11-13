package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerAssociation.
// Experimental.
type IRouteServerAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RouteServerAssociation resource.
	// Experimental.
	RouteServerAssociationRef() *RouteServerAssociationReference
}

// The jsii proxy for IRouteServerAssociationRef
type jsiiProxy_IRouteServerAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRouteServerAssociationRef) RouteServerAssociationRef() *RouteServerAssociationReference {
	var returns *RouteServerAssociationReference
	_jsii_.Get(
		j,
		"routeServerAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteServerAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteServerAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

