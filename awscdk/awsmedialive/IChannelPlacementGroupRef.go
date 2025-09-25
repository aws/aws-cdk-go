package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPlacementGroup.
// Experimental.
type IChannelPlacementGroupRef interface {
	constructs.IConstruct
	// A reference to a ChannelPlacementGroup resource.
	// Experimental.
	ChannelPlacementGroupRef() *ChannelPlacementGroupReference
}

// The jsii proxy for IChannelPlacementGroupRef
type jsiiProxy_IChannelPlacementGroupRef struct {
	internal.Type__constructsIConstruct
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

