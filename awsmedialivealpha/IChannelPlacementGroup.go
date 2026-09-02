package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Channel Placement Group.
// Experimental.
type IChannelPlacementGroup interface {
	interfacesawsmedialive.IChannelPlacementGroupRef
	awscdk.IResource
	// The ARN of the channel placement group.
	// Experimental.
	ChannelPlacementGroupArn() *string
	// The ID of the channel placement group.
	// Experimental.
	ChannelPlacementGroupId() *string
}

// The jsii proxy for IChannelPlacementGroup
type jsiiProxy_IChannelPlacementGroup struct {
	internal.Type__interfacesawsmedialiveIChannelPlacementGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannelPlacementGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IChannelPlacementGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) ChannelPlacementGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelPlacementGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) ChannelPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) ChannelPlacementGroupRef() *interfacesawsmedialive.ChannelPlacementGroupReference {
	var returns *interfacesawsmedialive.ChannelPlacementGroupReference
	_jsii_.Get(
		j,
		"channelPlacementGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelPlacementGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

