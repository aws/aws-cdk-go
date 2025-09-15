package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Channel.
// Experimental.
type IChannelRef interface {
	constructs.IConstruct
	// A reference to a Channel resource.
	// Experimental.
	ChannelRef() *ChannelReference
}

// The jsii proxy for IChannelRef
type jsiiProxy_IChannelRef struct {
	internal.Type__constructsIConstruct
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

