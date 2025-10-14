package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelAssociation.
// Experimental.
type IChannelAssociationRef interface {
	constructs.IConstruct
	// A reference to a ChannelAssociation resource.
	// Experimental.
	ChannelAssociationRef() *ChannelAssociationReference
}

// The jsii proxy for IChannelAssociationRef
type jsiiProxy_IChannelAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IChannelAssociationRef) ChannelAssociationRef() *ChannelAssociationReference {
	var returns *ChannelAssociationReference
	_jsii_.Get(
		j,
		"channelAssociationRef",
		&returns,
	)
	return returns
}

