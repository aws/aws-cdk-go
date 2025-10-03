package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rule.
// Experimental.
type IRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRuleRef
type jsiiProxy_IRuleRef struct {
	internal.Type__constructsIConstruct
}

