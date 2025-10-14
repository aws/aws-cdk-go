package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipeline.
// Experimental.
type IPipelineRef interface {
	constructs.IConstruct
	// A reference to a Pipeline resource.
	// Experimental.
	PipelineRef() *PipelineReference
}

// The jsii proxy for IPipelineRef
type jsiiProxy_IPipelineRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPipelineRef) PipelineRef() *PipelineReference {
	var returns *PipelineReference
	_jsii_.Get(
		j,
		"pipelineRef",
		&returns,
	)
	return returns
}

