package awsneptunegraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Graph.
// Experimental.
type IGraphRef interface {
	constructs.IConstruct
	// A reference to a Graph resource.
	// Experimental.
	GraphRef() *GraphReference
}

// The jsii proxy for IGraphRef
type jsiiProxy_IGraphRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGraphRef) GraphRef() *GraphReference {
	var returns *GraphReference
	_jsii_.Get(
		j,
		"graphRef",
		&returns,
	)
	return returns
}

