package awschatbot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MicrosoftTeamsChannelConfiguration.
// Experimental.
type IMicrosoftTeamsChannelConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMicrosoftTeamsChannelConfigurationRef
type jsiiProxy_IMicrosoftTeamsChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
}

