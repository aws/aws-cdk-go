package awssupportapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackChannelConfiguration.
// Experimental.
type ISlackChannelConfigurationRef interface {
	constructs.IConstruct
	// A reference to a SlackChannelConfiguration resource.
	// Experimental.
	SlackChannelConfigurationRef() *SlackChannelConfigurationReference
}

// The jsii proxy for ISlackChannelConfigurationRef
type jsiiProxy_ISlackChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
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

