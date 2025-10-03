package awsoam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsoam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Sink.
// Experimental.
type ISinkRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISinkRef
type jsiiProxy_ISinkRef struct {
	internal.Type__constructsIConstruct
}

