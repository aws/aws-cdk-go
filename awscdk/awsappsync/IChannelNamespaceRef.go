package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelNamespace.
// Experimental.
type IChannelNamespaceRef interface {
	constructs.IConstruct
	// A reference to a ChannelNamespace resource.
	// Experimental.
	ChannelNamespaceRef() *ChannelNamespaceReference
}

// The jsii proxy for IChannelNamespaceRef
type jsiiProxy_IChannelNamespaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IChannelNamespaceRef) ChannelNamespaceRef() *ChannelNamespaceReference {
	var returns *ChannelNamespaceReference
	_jsii_.Get(
		j,
		"channelNamespaceRef",
		&returns,
	)
	return returns
}

