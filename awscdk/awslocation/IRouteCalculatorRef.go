package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteCalculator.
// Experimental.
type IRouteCalculatorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRouteCalculatorRef
type jsiiProxy_IRouteCalculatorRef struct {
	internal.Type__constructsIConstruct
}

