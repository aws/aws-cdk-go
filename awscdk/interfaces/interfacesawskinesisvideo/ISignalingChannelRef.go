package interfacesawskinesisvideo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisvideo/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalingChannel.
// Experimental.
type ISignalingChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SignalingChannel resource.
	// Experimental.
	SignalingChannelRef() *SignalingChannelReference
}

// The jsii proxy for ISignalingChannelRef
type jsiiProxy_ISignalingChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISignalingChannelRef) SignalingChannelRef() *SignalingChannelReference {
	var returns *SignalingChannelReference
	_jsii_.Get(
		j,
		"signalingChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISignalingChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISignalingChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

