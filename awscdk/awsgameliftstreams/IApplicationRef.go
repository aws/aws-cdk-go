package awsgameliftstreams

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgameliftstreams/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Experimental.
type IApplicationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationRef
type jsiiProxy_IApplicationRef struct {
	internal.Type__constructsIConstruct
}

