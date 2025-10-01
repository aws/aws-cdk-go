package awspipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipe.
// Experimental.
type IPipeRef interface {
	constructs.IConstruct
	// A reference to a Pipe resource.
	// Experimental.
	PipeRef() *PipeReference
}

// The jsii proxy for IPipeRef
type jsiiProxy_IPipeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPipeRef) PipeRef() *PipeReference {
	var returns *PipeReference
	_jsii_.Get(
		j,
		"pipeRef",
		&returns,
	)
	return returns
}

