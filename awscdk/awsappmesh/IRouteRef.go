package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Route.
// Experimental.
type IRouteRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRouteRef
type jsiiProxy_IRouteRef struct {
	internal.Type__constructsIConstruct
}

