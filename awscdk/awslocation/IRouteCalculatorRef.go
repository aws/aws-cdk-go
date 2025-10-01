package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteCalculator.
// Experimental.
type IRouteCalculatorRef interface {
	constructs.IConstruct
	// A reference to a RouteCalculator resource.
	// Experimental.
	RouteCalculatorRef() *RouteCalculatorReference
}

// The jsii proxy for IRouteCalculatorRef
type jsiiProxy_IRouteCalculatorRef struct {
	internal.Type__constructsIConstruct
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

