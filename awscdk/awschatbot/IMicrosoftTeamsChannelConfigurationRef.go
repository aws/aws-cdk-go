package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MicrosoftTeamsChannelConfiguration.
// Experimental.
type IMicrosoftTeamsChannelConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MicrosoftTeamsChannelConfiguration resource.
	// Experimental.
	MicrosoftTeamsChannelConfigurationRef() *MicrosoftTeamsChannelConfigurationReference
}

// The jsii proxy for IMicrosoftTeamsChannelConfigurationRef
type jsiiProxy_IMicrosoftTeamsChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMicrosoftTeamsChannelConfigurationRef) MicrosoftTeamsChannelConfigurationRef() *MicrosoftTeamsChannelConfigurationReference {
	var returns *MicrosoftTeamsChannelConfigurationReference
	_jsii_.Get(
		j,
		"microsoftTeamsChannelConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrosoftTeamsChannelConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMicrosoftTeamsChannelConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

