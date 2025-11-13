package interfacesawsgameliftstreams

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgameliftstreams/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamGroup.
// Experimental.
type IStreamGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StreamGroup resource.
	// Experimental.
	StreamGroupRef() *StreamGroupReference
}

// The jsii proxy for IStreamGroupRef
type jsiiProxy_IStreamGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStreamGroupRef) StreamGroupRef() *StreamGroupReference {
	var returns *StreamGroupReference
	_jsii_.Get(
		j,
		"streamGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

