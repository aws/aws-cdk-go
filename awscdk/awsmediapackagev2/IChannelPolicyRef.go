package awsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPolicy.
// Experimental.
type IChannelPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ChannelPolicy resource.
	// Experimental.
	ChannelPolicyRef() *ChannelPolicyReference
}

// The jsii proxy for IChannelPolicyRef
type jsiiProxy_IChannelPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IChannelPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

