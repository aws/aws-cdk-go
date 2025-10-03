package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Service.
// Experimental.
type IServiceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceRef
type jsiiProxy_IServiceRef struct {
	internal.Type__constructsIConstruct
}

