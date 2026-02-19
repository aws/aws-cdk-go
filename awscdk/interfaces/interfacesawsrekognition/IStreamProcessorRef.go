package interfacesawsrekognition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamProcessor.
// Experimental.
type IStreamProcessorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StreamProcessor resource.
	// Experimental.
	StreamProcessorRef() *StreamProcessorReference
}

// The jsii proxy for IStreamProcessorRef
type jsiiProxy_IStreamProcessorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IStreamProcessorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStreamProcessorRef) StreamProcessorRef() *StreamProcessorReference {
	var returns *StreamProcessorReference
	_jsii_.Get(
		j,
		"streamProcessorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamProcessorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamProcessorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

