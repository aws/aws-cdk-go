package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackChannelConfiguration.
// Experimental.
type ISlackChannelConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SlackChannelConfiguration resource.
	// Experimental.
	SlackChannelConfigurationRef() *SlackChannelConfigurationReference
}

// The jsii proxy for ISlackChannelConfigurationRef
type jsiiProxy_ISlackChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISlackChannelConfigurationRef) SlackChannelConfigurationRef() *SlackChannelConfigurationReference {
	var returns *SlackChannelConfigurationReference
	_jsii_.Get(
		j,
		"slackChannelConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackChannelConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

