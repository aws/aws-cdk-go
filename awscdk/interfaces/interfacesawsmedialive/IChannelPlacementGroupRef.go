package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPlacementGroup.
// Experimental.
type IChannelPlacementGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ChannelPlacementGroup resource.
	// Experimental.
	ChannelPlacementGroupRef() *ChannelPlacementGroupReference
}

// The jsii proxy for IChannelPlacementGroupRef
type jsiiProxy_IChannelPlacementGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IChannelPlacementGroupRef) ChannelPlacementGroupRef() *ChannelPlacementGroupReference {
	var returns *ChannelPlacementGroupReference
	_jsii_.Get(
		j,
		"channelPlacementGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

