package awsoam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsoam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Sink.
// Experimental.
type ISinkRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Sink resource.
	// Experimental.
	SinkRef() *SinkReference
}

// The jsii proxy for ISinkRef
type jsiiProxy_ISinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISinkRef) SinkRef() *SinkReference {
	var returns *SinkReference
	_jsii_.Get(
		j,
		"sinkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISinkRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

