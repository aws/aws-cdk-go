package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelAssociation.
// Experimental.
type IChannelAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ChannelAssociation resource.
	// Experimental.
	ChannelAssociationRef() *ChannelAssociationReference
}

// The jsii proxy for IChannelAssociationRef
type jsiiProxy_IChannelAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IChannelAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

