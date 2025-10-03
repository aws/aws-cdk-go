package awsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Plan.
// Experimental.
type IPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPlanRef
type jsiiProxy_IPlanRef struct {
	internal.Type__constructsIConstruct
}

