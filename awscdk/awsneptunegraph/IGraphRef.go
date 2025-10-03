package awsneptunegraph

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Graph.
// Experimental.
type IGraphRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGraphRef
type jsiiProxy_IGraphRef struct {
	internal.Type__constructsIConstruct
}

