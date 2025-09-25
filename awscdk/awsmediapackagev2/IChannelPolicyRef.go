package awsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPolicy.
// Experimental.
type IChannelPolicyRef interface {
	constructs.IConstruct
	// A reference to a ChannelPolicy resource.
	// Experimental.
	ChannelPolicyRef() *ChannelPolicyReference
}

// The jsii proxy for IChannelPolicyRef
type jsiiProxy_IChannelPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IChannelPolicyRef) ChannelPolicyRef() *ChannelPolicyReference {
	var returns *ChannelPolicyReference
	_jsii_.Get(
		j,
		"channelPolicyRef",
		&returns,
	)
	return returns
}

