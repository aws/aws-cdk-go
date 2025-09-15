package awsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelGroup.
// Experimental.
type IChannelGroupRef interface {
	constructs.IConstruct
	// A reference to a ChannelGroup resource.
	// Experimental.
	ChannelGroupRef() *ChannelGroupReference
}

// The jsii proxy for IChannelGroupRef
type jsiiProxy_IChannelGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IChannelGroupRef) ChannelGroupRef() *ChannelGroupReference {
	var returns *ChannelGroupReference
	_jsii_.Get(
		j,
		"channelGroupRef",
		&returns,
	)
	return returns
}

