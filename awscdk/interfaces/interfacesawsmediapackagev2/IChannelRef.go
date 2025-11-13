package interfacesawsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Channel.
// Experimental.
type IChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Channel resource.
	// Experimental.
	ChannelRef() *ChannelReference
}

// The jsii proxy for IChannelRef
type jsiiProxy_IChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IChannelRef) ChannelRef() *ChannelReference {
	var returns *ChannelReference
	_jsii_.Get(
		j,
		"channelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

