package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipeline.
// Experimental.
type IPipelineRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Pipeline resource.
	// Experimental.
	PipelineRef() *PipelineReference
}

// The jsii proxy for IPipelineRef
type jsiiProxy_IPipelineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IPipelineRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipelineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

