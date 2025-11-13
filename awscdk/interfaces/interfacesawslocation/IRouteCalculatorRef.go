package interfacesawslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteCalculator.
// Experimental.
type IRouteCalculatorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RouteCalculator resource.
	// Experimental.
	RouteCalculatorRef() *RouteCalculatorReference
}

// The jsii proxy for IRouteCalculatorRef
type jsiiProxy_IRouteCalculatorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRouteCalculatorRef) RouteCalculatorRef() *RouteCalculatorReference {
	var returns *RouteCalculatorReference
	_jsii_.Get(
		j,
		"routeCalculatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteCalculatorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteCalculatorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

