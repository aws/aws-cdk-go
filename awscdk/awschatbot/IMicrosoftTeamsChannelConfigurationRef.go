package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MicrosoftTeamsChannelConfiguration.
// Experimental.
type IMicrosoftTeamsChannelConfigurationRef interface {
	constructs.IConstruct
	// A reference to a MicrosoftTeamsChannelConfiguration resource.
	// Experimental.
	MicrosoftTeamsChannelConfigurationRef() *MicrosoftTeamsChannelConfigurationReference
}

// The jsii proxy for IMicrosoftTeamsChannelConfigurationRef
type jsiiProxy_IMicrosoftTeamsChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
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

