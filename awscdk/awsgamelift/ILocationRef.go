package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Location.
// Experimental.
type ILocationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationRef
type jsiiProxy_ILocationRef struct {
	internal.Type__constructsIConstruct
}

