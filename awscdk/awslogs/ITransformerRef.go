package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Transformer.
// Experimental.
type ITransformerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Transformer resource.
	// Experimental.
	TransformerRef() *TransformerReference
}

// The jsii proxy for ITransformerRef
type jsiiProxy_ITransformerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITransformerRef) TransformerRef() *TransformerReference {
	var returns *TransformerReference
	_jsii_.Get(
		j,
		"transformerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransformerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransformerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

