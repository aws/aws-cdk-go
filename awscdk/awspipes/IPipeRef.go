package awspipes

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipe.
// Experimental.
type IPipeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPipeRef
type jsiiProxy_IPipeRef struct {
	internal.Type__constructsIConstruct
}

