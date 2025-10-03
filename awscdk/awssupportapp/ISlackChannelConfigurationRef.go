package awssupportapp

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackChannelConfiguration.
// Experimental.
type ISlackChannelConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISlackChannelConfigurationRef
type jsiiProxy_ISlackChannelConfigurationRef struct {
	internal.Type__constructsIConstruct
}

