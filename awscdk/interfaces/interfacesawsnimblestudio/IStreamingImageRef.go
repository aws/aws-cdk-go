package interfacesawsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamingImage.
// Experimental.
type IStreamingImageRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StreamingImage resource.
	// Experimental.
	StreamingImageRef() *StreamingImageReference
}

// The jsii proxy for IStreamingImageRef
type jsiiProxy_IStreamingImageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IStreamingImageRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IStreamingImageRef) StreamingImageRef() *StreamingImageReference {
	var returns *StreamingImageReference
	_jsii_.Get(
		j,
		"streamingImageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamingImageRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStreamingImageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

